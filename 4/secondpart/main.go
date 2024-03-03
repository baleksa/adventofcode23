package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Card struct {
	num     int
	cardStr string
	matches int
	nCopy   int
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	in := bufio.NewScanner(f)

	cards := []*Card{}
	i := 1
	for in.Scan() {
		card := Card{i, in.Text(), cardValue(in.Text()), 1}
		cards = append(cards, &card)
	}

	fmt.Println(totalCards(cards))
}

func totalCards(cards []*Card) int {
	countCopies(cards)
	total := 0
	for _, card := range cards {
		total += card.nCopy
	}

	return total
}

func countCopies(cards []*Card) {
	for i := 0; i < len(cards)-1; i++ {
		card := cards[i]
		for _, c := range cards[i+1 : i+1+card.matches] {
			c.nCopy += card.nCopy
		}
	}
}

func cardValue(card string) int {
	splitChar := "|"
	cardIDLen := 9
	splitCharInd := strings.Index(card, splitChar)
	winningNumsStr := strings.TrimSpace(card[cardIDLen:splitCharInd])
	numsStr := strings.TrimSpace(card[splitCharInd+1:])
	winningNums := getInts(winningNumsStr)
	nums := getInts(numsStr)

	numsMatched := lenOfIntersection(winningNums, nums)
	return numsMatched
}

func lenOfIntersection(a1, a2 []int) int {
	A := map[int]bool{}

	for _, x := range a1 {
		A[x] = true
	}
	lenOfInter := 0
	for _, x := range a2 {
		if A[x] {
			lenOfInter++
		}
	}

	return lenOfInter
}

func getInts(s string) []int {
	a := []int{}

	for _, numStr := range strings.Fields(s) {
		var num int
		_, err := fmt.Sscanf(numStr, "%d", &num)
		if err != nil {
			log.Fatal(err)
		}
		a = append(a, num)
	}

	return a
}
