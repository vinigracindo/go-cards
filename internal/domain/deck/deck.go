package deck

import (
	"math/rand"
	"time"

	"github.com/vinigracindo/go-cards/internal/domain/card"
	"github.com/vinigracindo/go-cards/pkg/ctypes"
)

type Deck struct {
	Cards []*card.Card
}

func (deck *Deck) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	for i := range deck.Cards {
		j := rand.Intn(i + 1)
		deck.Cards[i], deck.Cards[j] = deck.Cards[j], deck.Cards[i]
	}
}

func (deck *Deck) PopCard() (*card.Card, error) {
	last := len(deck.Cards) - 1
	if last < 0 {
		return nil, ErrNoCards
	}
	card := deck.Cards[last]
	deck.Cards = deck.Cards[:last]
	return card, nil
}

func NewDeck() *Deck {
	deck := &Deck{}
	deck.Cards = make([]*card.Card, 0)
	for _, suit := range ctypes.AllSuits() {
		for _, value := range ctypes.AllValues() {
			deck.Cards = append(deck.Cards, card.NewCard(value, suit))
		}
	}
	return deck
}
