Go Cards
========

```go
package main

import (
	"fmt"

	"github.com/vinigracindo/go-cards/pkg"
)

func main() {
    numberOfDecks = 1
    // Create a Dealer with numberOfDecks=1 Deck (52 cards)
    dealer := dealer.NewDealer(numberOfDecks)

    // Shuffle Cards
    // Seed: time.Now().UnixNano()
    dealer.ShuffleCards()

    // Burn a card from the top of the deck
    dealer.BurnCard()

    // Deal a card to the player (top of the deck)
    card, _ := dealer.Deal()
}
```

```go
dealer := dealer.NewDealer(numberOfDecks)
fmt.Println(dealer.Deck.Cards)
```

```console
[
 2[♠] 3[♠] 4[♠] 5[♠] 6[♠] 7[♠] 8[♠] 9[♠] 10[♠] J[♠] Q[♠] K[♠] A[♠] 
 2[♥] 3[♥] 4[♥] 5[♥] 6[♥] 7[♥] 8[♥] 9[♥] 10[♥] J[♥] Q[♥] K[♥] A[♥] 
 2[♦] 3[♦] 4[♦] 5[♦] 6[♦] 7[♦] 8[♦] 9[♦] 10[♦] J[♦] Q[♦] K[♦] A[♦] 
 2[♣] 3[♣] 4[♣] 5[♣] 6[♣] 7[♣] 8[♣] 9[♣] 10[♣] J[♣] Q[♣] K[♣] A[♣]
]
```

```go
dealer.ShuffleCards()
fmt.Println(dealer.Deck.Cards)
```

```console
# Random Result
[
 2[♦] 9[♣] 2[♥] 7[♥] K[♠] 7[♦] 2[♠] 8[♠] 9[♥] 7[♣] K[♣] 6[♣] Q[♠] 
 2[♣] 6[♦] K[♦] J[♣] A[♣] 5[♣] A[♥] 6[♥] Q[♣] K[♥] 3[♥] A[♠] 5[♥] 
 8[♦] 3[♦] 4[♣] 9[♦] 3[♠] 6[♠] 4[♥] A[♦] 4[♠] 7[♠] 10[♣] Q[♥] 8[♣] 
 10[♠] Q[♦] 5[♠] J[♥] 10[♥] 8[♥] 10[♦] J[♠] 5[♦] 4[♦] 3[♣] J[♦] 9[♠]
]
```

```go
dealer.BurnCard()
fmt.Println(dealer.Deck.Cards)
```

```console
# Remove Last
[
    2[♦] 9[♣] 2[♥] 7[♥] K[♠] 7[♦] 2[♠] 8[♠] 9[♥] 7[♣] K[♣] 6[♣] Q[♠] 
    2[♣] 6[♦] K[♦] J[♣] A[♣] 5[♣] A[♥] 6[♥] Q[♣] K[♥] 3[♥] A[♠] 5[♥] 
    8[♦] 3[♦] 4[♣] 9[♦] 3[♠] 6[♠] 4[♥] A[♦] 4[♠] 7[♠] 10[♣] Q[♥] 8[♣] 
    10[♠] Q[♦] 5[♠] J[♥] 10[♥] 8[♥] 10[♦] J[♠] 5[♦] 4[♦] 3[♣] J[♦]
]
```

```go
card, err := dealer.Deal()
fmt.Println("Card: ", card)
fmt.Println("Deck: ", card)
fmt.Println(dealer.Deck.Cards)
```

```console
Card: J[♦]
Deck:
[
 2[♦] 9[♣] 2[♥] 7[♥] K[♠] 7[♦] 2[♠] 8[♠] 9[♥] 7[♣] K[♣] 6[♣] Q[♠] 
 2[♣] 6[♦] K[♦] J[♣] A[♣] 5[♣] A[♥] 6[♥] Q[♣] K[♥] 3[♥] A[♠] 5[♥] 
 8[♦] 3[♦] 4[♣] 9[♦] 3[♠] 6[♠] 4[♥] A[♦] 4[♠] 7[♠] 10[♣] Q[♥] 8[♣] 
 10[♠] Q[♦] 5[♠] J[♥] 10[♥] 8[♥] 10[♦] J[♠] 5[♦] 4[♦] 3[♣]
]
```
