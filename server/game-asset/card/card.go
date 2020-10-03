package card

const (
	Ace = iota + 1
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
	Joker
)

const (
	DifferentCardTypes = 14
	DifferentColors    = 2
)

var (
	Colors = []string{
		"Red",
		"Black",
	}
)

type Card struct {
	value int
	color string
}

func NewCard(value int, color string) Card {
	return Card{
		value: value,
		color: color,
	}
}
