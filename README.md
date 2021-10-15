# cards

Playing card library for use in card games like Blackjack.  


## Sample Card Output
[4♠][9♥][3♥][K♦]


## Getting Started
Import cards library into your go package
```bash
import (
  "github.com/mbarley333/cards"
)
```

## Features
* create shuffled deck(s) of any size
```bash
  deck := cards.NewDeck(
		cards.WithNumberOfDecks(6),
	)
```
* render the playing card emoji for use in command line

```bash
  card := cards.Card{cards.Rank: Ace, cards.Suit: Heart}
  card.Render()
```



