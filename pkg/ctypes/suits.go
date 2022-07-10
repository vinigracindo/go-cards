package ctypes

//go:generate stringer -type=Suit
type Suit string

const (
	Spade   Suit = "♠"
	Heart   Suit = "♥"
	Diamond Suit = "♦"
	Club    Suit = "♣"
)

func AllSuits() []Suit {
	return []Suit{Spade, Heart, Diamond, Club}
}
