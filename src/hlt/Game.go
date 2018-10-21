package hlt

import (
	"fmt"
	"hlt/gameconfig"
	"hlt/input"
	"hlt/log"
)

// Game - Structure holding all metadata for the game
type Game struct {
	numPlayers int
	Me         *Player
	players    []*Player
	Map        *GameMap
	TurnNumber int
}

func (g *Game) String() string {
	return fmt.Sprintf("Game{NumPlayers=%d,Me=%s,players=%d,Map=%s}", g.numPlayers, g.Me.String(), len(g.players), g.Map.String())
}

// Ready - When run, notifies the server that the bot is ready to start
func (g *Game) Ready(name string) {
	fmt.Println(name)
}

// NewGame - Creates a new game for bot named by the parameter
func NewGame() *Game {
	var input = input.GetInstance()
	if !input.Scanner.Scan() {
		return nil
	}

	var constantsString = input.Scanner.Text()
	gameconfig.Init(constantsString)
	var numPlayers, _ = input.GetInt()
	var myID, _ = input.GetInt()
	var players = make([]*Player, numPlayers)
	for i := 0; i < len(players); i++ {
		players[i] = NewPlayer()
	}
	var gameMap = GenerateGameMap()
	var me = players[myID]
	return &Game{numPlayers, me, players, gameMap, 0}
}

// UpdateFrame - Runs a single turn in the game
func (g *Game) UpdateFrame() {
	var logger = log.GetInstance()
	var input = input.GetInstance()
	g.TurnNumber, _ = input.GetInt()
	logger.Printf("=============== TURN %d ================\n", g.TurnNumber)
	for i := 0; i < len(g.players); i++ {
		// Player ID variable being read here. It isn't used, so we can't assign it in Go without error, but we still need to consume it from the input
		input.GetInt()
		var numShips, _ = input.GetInt()
		var numDropoffs, _ = input.GetInt()
		var halite, _ = input.GetInt()
		g.players[i].Update(numShips, numDropoffs, halite)
	}
	g.Map.Update()
	for i := 0; i < len(g.players); i++ {
		var player = g.players[i]
		for j := 0; j < len(player.Ships); i++ {
			var ship = player.Ships[j]
			g.Map.AtEntity(ship.E).MarkUnsafe(ship)
		}
		for k := 0; k < len(player.Dropoffs); k++ {
			var dropoff = player.Dropoffs[k]
			g.Map.AtEntity(dropoff.E).structure = dropoff.E
		}
	}
}

// EndTurn -
func (g *Game) EndTurn(commands []Command) {
	var command Command
	for i := 0; i < len(commands); i++ {
		command = commands[i]
		fmt.Print(command.CommandString())
		fmt.Print(" ")
	}
	fmt.Println()
}
