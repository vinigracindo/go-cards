package card_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vinigracindo/go-cards/internal/domain/card"
	"github.com/vinigracindo/go-cards/pkg/ctypes"
)

func TestCard_NewCard(t *testing.T) {
	newCard := card.NewCard(ctypes.Ace, ctypes.Spade)
	assert.IsType(t, &card.Card{}, newCard)
}

func TestCard_String(t *testing.T) {
	tests := []struct {
		name string
		card *card.Card
		want string
	}{
		{
			name: "Ace of Spade",
			card: card.NewCard(ctypes.Ace, ctypes.Spade),
			want: "A[♠]",
		},
		{
			name: "Jack of Diamond",
			card: card.NewCard(ctypes.Jack, ctypes.Diamond),
			want: "J[♦]",
		},
		{
			name: "Ten of Club",
			card: card.NewCard(ctypes.Ten, ctypes.Club),
			want: "10[♣]",
		},
		{
			name: "Two of Heart",
			card: card.NewCard(ctypes.Two, ctypes.Heart),
			want: "2[♥]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.card.String())
		})
	}
}
