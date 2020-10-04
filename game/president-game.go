package game

import (
	"github.com/google/uuid"
)

type PresidentGame struct {
	id             uuid.UUID
	name           string
	maxPlayers     int
	currentPlayers int
	inProgress     bool
}

// NewPresidentGame - Create a new game session for Presidents
func NewPresidentGame(name string, maxPlayers int) Game {
	return &PresidentGame{
		name:       name,
		maxPlayers: maxPlayers,
	}
}

// Start - Launch the game
func (g *PresidentGame) Start() {
	g.inProgress = true
}
