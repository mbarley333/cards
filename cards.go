package cards

import (
	"fmt"
	"math/rand"
	"strings"

	"time"
)

type Suit rune

const (
	Club    Suit = '\u2663'
	Diamond Suit = '\u2666'
	Heart   Suit = '\u2665'
	Spade   Suit = '\u2660'
)

var suitMap = map[Suit]string{
	Spade:   "Spade",
	Diamond: "Diamond",
	Club:    "Club",
	Heart:   "Heart",
}

var suits = []Suit{Club, Diamond, Heart, Spade}

type Rank int

const (
	_ Rank = iota
	Ace
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
)

var ranks = []Rank{
	Ace, Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten, Jack, Queen, King,
}

var rankMap = map[Rank]string{
	Ace:   "A",
	Two:   "2",
	Three: "3",
	Four:  "4",
	Five:  "5",
	Six:   "6",
	Seven: "7",
	Eight: "8",
	Nine:  "9",
	Ten:   "10",
	Jack:  "J",
	Queen: "Q",
	King:  "K",
}

func (c Card) String() string {
	return fmt.Sprintf("%s of %ss", rankMap[c.Rank], suitMap[c.Suit])
}

type Color string

const (
	ColorRed   Color = "\033[31m"
	ColorReset Color = "\033[0m"
)

type Card struct {
	Rank Rank
	Suit Suit
}

type Deck struct {
	Cards  []Card
	Count  int
	Random *rand.Rand
}

type Option func(*Deck) error

func WithNumberOfDecks(number int) Option {
	return func(d *Deck) error {
		d.Count = number
		return nil
	}
}

func WithRandom(random *rand.Rand) Option {
	return func(d *Deck) error {
		d.Random = random
		return nil
	}
}

func NewDeck(opts ...Option) Deck {

	deck := &Deck{
		Count:  1,
		Random: rand.New(rand.NewSource(time.Now().UnixNano())),
	}

	for _, o := range opts {
		o(deck)
	}

	for i := 0; i < deck.Count; i++ {
		for _, suit := range suits {
			for _, rank := range ranks {
				deck.Cards = append(deck.Cards, Card{Suit: suit, Rank: rank})
			}

		}
	}
	shuffled_cards := make([]Card, len(deck.Cards))
	perm := deck.Random.Perm(len(deck.Cards))

	for i, j := range perm {
		shuffled_cards[i] = deck.Cards[j]
	}
	deck.Cards = shuffled_cards
	return Deck{
		Cards: shuffled_cards,
	}
}

func (c Card) Notation() string {

	card := fmt.Sprintf("%s%s", rankMap[c.Rank], string(c.Suit))

	return card
}

func (c Card) Render() string {
	builder := strings.Builder{}
	builder.WriteString("[")
	if c.Suit == Diamond || c.Suit == Heart {
		builder.WriteString(string(ColorRed))
	} else {
		builder.WriteString(string(ColorReset))
	}
	builder.WriteString(c.Notation())
	builder.WriteString("]")
	builder.WriteString(string(ColorReset))

	return builder.String()

}

func RenderHand(hand []Card) string {
	builder := strings.Builder{}
	for _, card := range hand {

		builder.WriteString(card.Render())
	}
	return builder.String()

}
