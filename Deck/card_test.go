package deck

import (
	"fmt"
	"testing"
)

func ExampleCard() {
	fmt.Println(Card{Rank: Ace, Suit: Heart})
	fmt.Println(Card{Suit: Joker})
	fmt.Println(Card{Rank: Two, Suit: Spade})
	fmt.Println(Card{Rank: King, Suit: Diamond})
	fmt.Println(Card{Rank: Six, Suit: Club})

	// Output:
	// Ace of Hearts
	// Joker
	// Two of Spades
	// King of Diamonds
	// Six of Clubs
}

func TestNew(t *testing.T) {
	cards := New()
	if len(cards) != 13*4 {
		t.Error("Wrong number of cards")
	}
}

func TestDefaultSort(t *testing.T) {
	cards := New(DefaultSort)
	exp := Card{Suit: Spade, Rank: Ace}
	if exp != cards[0] {
		t.Error("Expected Ace of Spades. Got:", cards[0])
	}
}

func TestSort(t *testing.T) {
	cards := New(Sort(Less))
	exp := Card{Suit: Spade, Rank: Ace}
	if exp != cards[0] {
		t.Error("Expected Ace of Spades. Got:", cards[0])
	}
}

func TestJokers(t *testing.T) {
	cards := New(Jokers(3))
	count := 0
	for _, c := range cards {
		if c.Suit == Joker {
			count++
		}
	}
	if count != 3 {
		t.Error("Expected 3 Jokers. Got:", count)
	}
}

func TestFilter(t *testing.T) {
	filter := func(card Card) bool {
		return card.Rank == Two || card.Rank == Three
	}
	cards := New(Filter(filter))
	for _, c := range cards {
		if c.Rank == Two || c.Rank == Three {
			t.Error("Expected all Twos and Threes to be filtered out!")
		}
	}
}

func TestDeck(t *testing.T) {
	cards := New(Deck(3))
	total := 13 * 4 * 3
	if len(cards) != total {
		t.Errorf("Expected %d cards. Got %d", total, len(cards))
	}
}
