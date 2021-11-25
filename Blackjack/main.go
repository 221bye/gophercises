package main

import (
	"fmt"
	"strings"

	"gophercises/deck"
)

type Hand []deck.Card

func (h Hand) String() string {
	strs := make([]string, len(h))
	for i := range h {
		strs[i] = h[i].String()
	}
	return strings.Join(strs, ", ") + "."
}

func (h Hand) DealerString() string {
	return h[0].String() + ", HIDDEN"
}

func (h Hand) Score() int {
	minScore := h.MinScore()
	if minScore > 11 {
		return minScore
	}
	for _, crd := range h {
		if crd.Rank == deck.Ace {
			return minScore + 10
		}
	}
	return minScore
}

func (h Hand) MinScore() int {
	score := 0
	for _, crd := range h {
		if crd.Rank > 10 {
			score += 10
		} else {
			score += int(crd.Rank)
		}
	}
}

func main() {
	cards := deck.New(deck.Deck(3), deck.Shuffle)
	var card deck.Card
	var player, dealer Hand

	for _, hand := range []*Hand{&player, &dealer} {
		for i := 0; i < 2; i++ {
			card, cards = cards[0], cards[1:]
			*hand = append(*hand, card)
		}
	}

	var input string
	playerScore := player.Score()

	for input != "s" || playerScore < 21 {
		fmt.Println("Player:", player, playerScore)
		fmt.Println("Dealer:", dealer.DealerString())
		fmt.Println("(h)it, (s)tand?")
		fmt.Scanf("%s\n", &input)
		switch input {
		case "h":
			card, cards = cards[0], cards[1:]
			player = append(player, card)
			playerScore = player.Score()
		}
	}

	for dealer.Score() <= 16 || (dealer.Score() == 17 && dealer.MinScore() != 17) {
		card, cards = cards[0], cards[1:]
		dealer = append(dealer, card)
	}

	playerScore = player.Score()
	dealerScore := dealer.Score()

	fmt.Println("Player:", player, playerScore)
	fmt.Println("Dealer:", dealer, dealerScore)

	switch {
	case playerScore > 21:
		fmt.Println("You lost")
	case dealerScore > 21:
		fmt.Println("You won")
	case playerScore > dealerScore:
		fmt.Println("You won")
	case playerScore < dealerScore:
		fmt.Println("You lost")
	case playerScore == dealerScore:
		fmt.Println("Draw")
	}
}
