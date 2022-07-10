package ctypes

//go:generate stringer -type=Number
type Value string

const (
	Two   Value = "2"
	Three Value = "3"
	Four  Value = "4"
	Five  Value = "5"
	Six   Value = "6"
	Seven Value = "7"
	Eight Value = "8"
	Nine  Value = "9"
	Ten   Value = "10"
	Jack  Value = "J"
	Queen Value = "Q"
	King  Value = "K"
	Ace   Value = "A"
)

func AllValues() []Value {
	return []Value{Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten, Jack, Queen, King, Ace}
}
