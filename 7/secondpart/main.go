package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

type Hand struct {
	cards string
	bid   int
}

type Card rune

type HandsByStrength []Hand

var cards = [...]Card{'A', 'K', 'Q', 'T', '9', '8', '7', '6', '5', '4', '3', '2', 'J'}

func (h HandsByStrength) Len() int {
	return len(h)
}

func (h HandsByStrength) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func main() {
	input_file := "input.txt"
	if len(os.Args) == 2 {
		input_file = os.Args[1]
	}
	f, err := os.Open(input_file)
	if err != nil {
		log.Fatal(err)
	}
	in := bufio.NewScanner(f)
	hands := []Hand{}
	for in.Scan() {
		line := in.Text()
		hands = append(hands, parseHand(line))
	}

	fmt.Println(totalWinnings(hands))
}

func parseHand(line string) Hand {
	var bid int
	var cards string
	if n, err := fmt.Sscanf(line, "%s %d", &cards, &bid); n != 2 {
		log.Fatal(err)
	}
	return Hand{cards, bid}
}

func totalWinnings(hands []Hand) uint64 {

	fmt.Println("Before sorting:")
	for i, hand := range hands {
		fmt.Printf("%d => %s %s\n", i+1, hand.cards, typeStr[rank(hand)])
	}

	sortHands(hands)

	fmt.Println("After sorting:")
	for i, hand := range hands {
		fmt.Printf("%d => %s %s\n", i+1, hand.cards, typeStr[rank(hand)])
	}

	var total uint64 = 0
	for i, hand := range hands {
		total += uint64(i+1) * uint64(hand.bid)
	}
	return total
}

type Type int

var typeStr = []string{"FIOK", "FOK", "FH", "TOK", "TP", "OP", "HC"}

const (
	FIOK Type = iota
	FOK
	FH
	TOK
	TP
	OP
	HC
)

func hasOnePair(hand Hand) bool {
	for _, card := range cards {
		if strings.Count(hand.cards, string(card)) == 2 {
			return true
		}
	}
	return false
}

func hasTwoPair(hand Hand) bool {
	var c1 Card
	found := false
	for _, card := range cards {
		if strings.Count(hand.cards, string(card)) == 2 {
			c1 = card
			found = true
		}
	}
	if !found {
		return false
	}
	for _, card := range cards {
		if strings.Count(hand.cards, string(card)) == 2 && card != c1 {
			return true
		}
	}

	return false
}

func hasThree(hand Hand) bool {
	for _, card := range cards {
		if strings.Count(hand.cards, string(card)) == 3 {
			return true
		}
	}

	return false

}

func hasFullHouse(hand Hand) bool {
	var c1 Card
	found := false
	for _, card := range cards {
		if strings.Count(hand.cards, string(card)) == 3 {
			c1 = card
			found = true
		}
	}
	if !found {
		return false
	}
	for _, card := range cards {
		if strings.Count(hand.cards, string(card)) == 2 && card != c1 {
			return true
		}
	}

	return false

}

func hasFour(hand Hand) bool {
	for _, card := range cards {
		if strings.Count(hand.cards, string(card)) == 4 {
			return true
		}
	}

	return false
}

func hasFive(hand Hand) bool {
	for _, card := range cards {
		if strings.Count(hand.cards, string(card)) == 5 {
			return true
		}
	}
	return false
}

func rank(hand Hand) Type {
	nJokers := strings.Count(hand.cards, "J")
	if hasFive(hand) {
		return FIOK
	} else if hasFour(hand) {
		switch nJokers {
		case 4, 1:
			return FIOK
		default:
			return FOK
		}
	} else if hasFullHouse(hand) {
		switch nJokers {
		case 2, 3:
			return FIOK
		default:
			return FH
		}

	} else if hasThree(hand) {
		switch nJokers {
		case 3, 1:
			return FOK
		default:
			return TOK
		}
	} else if hasTwoPair(hand) {
		switch nJokers {
		case 1:
			return FH
		case 2:
			return FOK
		default:
			return TP
		}
	} else if hasOnePair(hand) {
		switch nJokers {
		case 1, 2:
			return TOK
		default:
			return OP
		}
	}
	if nJokers == 1 {
		return OP
	}
	return HC

}

func sortHands(hands []Hand) {
	sort.Sort(HandsByStrength(hands))
}

func index(card Card) int {
	for i := range cards {
		if card == cards[i] {
			return i
		}
	}
	panic("unreachable")
}

func (h HandsByStrength) Less(i, j int) bool {
	r1 := rank(h[i])
	r2 := rank(h[j])

	if r1 == r2 {
		c1 := []Card(h[i].cards)
		c2 := []Card(h[j].cards)
		for k := range c1 {
			ic1 := index(c1[k])
			ic2 := index(c2[k])
			if ic1 != ic2 {
				return ic1 > ic2
			}
		}
	}
	return r1 > r2
}
