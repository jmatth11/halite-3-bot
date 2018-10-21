package logic

import (
	"hlt"
	"hlt/gameconfig"
)

// ShipDecision - Representation of what logic the ship should perform
type ShipDecision int

const (
	// Collect - Collect more halite
	Collect = ShipDecision(iota)
	// Return - Return to nearest drop off
	Return
	// Convert - Convert ship into new drop off point
	Convert
)

// GameAI - Object to store/handle overall game logic
type GameAI struct {
	game                 *hlt.Game
	config               *gameconfig.Constants
	shipsMarkedForReturn map[int]bool
}

// NewGameAI - Generate a new GameAI object
func NewGameAI(g *hlt.Game, c *gameconfig.Constants) *GameAI {
	return &GameAI{
		game:                 g,
		config:               c,
		shipsMarkedForReturn: make(map[int]bool),
	}
}

// ShipLogic - Figure out what decision the ship should make next
func (gm *GameAI) ShipLogic(ship *hlt.Ship) ShipDecision {
	currentCell := gm.game.Map.AtEntity(ship.E)
	maxHalite, _ := gm.config.GetInt(gameconfig.MaxHalite)
	if t, ok := gm.shipsMarkedForReturn[ship.E.ID()]; ok && t {
		return Return
	}
	if ship.IsFull() || (float64(ship.Halite)/float64(maxHalite)) > 0.8 {
		gm.shipsMarkedForReturn[ship.E.ID()] = true
		return Return
	}
	if currentCell.Halite < (maxHalite / 10) {
		return Collect
	}
	return Collect
}
