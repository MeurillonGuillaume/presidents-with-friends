package carddeck

import (
	"math/rand"
	"time"

	"github.com/presidents-with-friends/presidents-with-friends/server/game-asset/card"
)

const (
	cardsInDeck = 54 // 52 cards + 2 jokers
)

type CardDeck struct {
	cards     [cardsInDeck]card.Card
	usedCards [cardsInDeck]card.Card
}

// DrawRandomHands - Draw cards for each player in random order
func (cd *CardDeck) DrawRandomHands(playerCount int) (result [][]card.Card) {
	cardsPerPlayer := cardsInDeck / playerCount

	for player := 0; player < playerCount; player++ {
		for cardNo := 0; cardNo < cardsPerPlayer; cardNo++ {
			// Set random seed
			rand.Seed(time.Now().UnixNano())
			// Assign random cards to player
		}
	}

	return
}

// Shuffle - Assign values to deck
func (cd *CardDeck) Shuffle() {
	rand.Seed(time.Now().UnixNano())

	cd.reset()

	var tmp [cardsInDeck]card.Card
	for color := 0; color < card.DifferentColors; color++ {
		for i := 0; i < card.DifferentCardTypes; i++ {
			tmp[i+(1*color)] = card.NewCard(i, card.Colors[color])
		}
	}

	rand.Shuffle(cardsInDeck, func(i, j int) {
		tmp[i], tmp[j] = tmp[j], tmp[i]
	})

	cd.cards = tmp
}

// reset - empty current values
func (cd *CardDeck) reset() {
	cd.cards = [cardsInDeck]card.Card{}
	cd.usedCards = [cardsInDeck]card.Card{}
}
