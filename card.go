//go:generate stringer -type=Suit,Rank

package deck

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// Suit represents each playing card suit.
type Suit uint8

const (
	// Spade suits are sorted first in most new decks.
	Spade Suit = iota
	// Diamond suits are sorted second.
	Diamond
	// Club suits are sorted third.
	Club
	// Heart suits are sorted last.
	Heart
	// Joker cards are a special case.
	Joker
)

var suits = [...]Suit{Spade, Diamond, Club, Heart}

// Rank represents each playing card rank.
type Rank uint8

const (
	_ Rank = iota
	// Ace represents rank 1.
	Ace
	// Two represents rank 2.
	Two
	// Three represents rank 3.
	Three
	// Four represents rank 4.
	Four
	// Five represents rank 5.
	Five
	// Six represents rank 6.
	Six
	// Seven represents rank 7.
	Seven
	// Eight represents rank 8.
	Eight
	// Nine represents rank 9.
	Nine
	// Ten represents rank 10.
	Ten
	// Jack represents rank 11.
	Jack
	// Queen represents rank 12.
	Queen
	// King represents rank 13.
	King
)

const (
	minRank = Ace
	maxRank = King
)

// Card data includes a Suit and a Rank.
type Card struct {
	Suit
	Rank
}

func (c Card) String() string {
	if c.Suit == Joker {
		return c.Suit.String()
	}

	return fmt.Sprintf("%s of %ss", c.Rank.String(), c.Suit.String())
}

// New creates a slice or deck of type Card.
func New(opts ...func([]Card) []Card) []Card {
	var cards []Card
	for _, suit := range suits {
		for rank := minRank; rank <= maxRank; rank++ {
			cards = append(cards, Card{Suit: suit, Rank: rank})
		}
	}

	for _, opt := range opts {
		cards = opt(cards)
	}

	return cards
}

// DefaultSort returns a sorted deck of cards.
func DefaultSort(cards []Card) []Card {
	sort.Slice(cards, Less(cards))
	return cards
}

// Sort ...
func Sort(less func(cards []Card) func(i, j int) bool) func([]Card) []Card {
	return func(cards []Card) []Card {
		sort.Slice(cards, less(cards))
		return cards
	}
}

// Less ...
func Less(cards []Card) func(i, j int) bool {
	return func(i, j int) bool {
		return absRank(cards[i]) < absRank(cards[j])
	}
}

func absRank(c Card) int {
	return int(c.Suit)*int(maxRank) + int(c.Rank)
}

// Shuffle sorts the deck of cards in a random order.
func Shuffle(cards []Card) []Card {
	ret := make([]Card, len(cards))
	r := rand.New(rand.NewSource(time.Now().Unix()))
	perm := r.Perm(len(cards))

	for i, j := range perm {
		ret[i] = cards[j]
	}

	return ret
}

// Jokers special case...
func Jokers(n int) func([]Card) []Card {
	return func(cards []Card) []Card {
		for i := 0; i < n; i++ {
			cards = append(cards, Card{
				Rank: Rank(i),
				Suit: Joker,
			})
		}

		return cards
	}
}
