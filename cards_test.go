package cards_test

import (
	"math/rand"
	"testing"

	"github.com/mbarley333/cards"
)

func TestShuffle(t *testing.T) {
	t.Parallel()
	random := rand.New(rand.NewSource(1))

	deck := cards.NewDeck(
		cards.WithNumberOfDecks(1),
		cards.WithRandom(random),
	)

	//shuffle := deck.Shuffle(random)

	got := deck.Cards[0].String()

	want := "Nine of Diamonds"

	if want != got {
		t.Fatalf("want: %s, got: %s", want, got)
	}

}
