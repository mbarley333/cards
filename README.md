# cards

Emoji playing card library for use in card games like Blackjack.  

# Sample output
[4♠][9♥][3♥][K♦]



# Getting started
* Run this command inside your project
```bash
go get github.com/mbarley333/cards
```

* Import cards library into your go package
```bash
import "github.com/mbarley333/cards"

```



# Features
* create shuffled deck(s) of any size
```bash
  deck := cards.NewDeck(
		cards.WithNumberOfDecks(6),
	)
```
* render the playing card emoji for use in command line

```bash
  card := cards.Card{cards.Rank: Ace, cards.Suit: Heart}
  fmt.Println(card.Render())
```




