package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	in := bufio.NewScanner(f)
	sumPoints := 0
	for in.Scan() {
		card := in.Text()
		points := cardValue(card)
		fmt.Println("Points:", points)
		sumPoints += points
	}
	fmt.Println("Points:", sumPoints)
}

func cardValue(card string) int {
	splitChar := "|"
	cardIDLen := 9
	splitCharInd := strings.Index(card, splitChar)
	winningNumsStr := strings.TrimSpace(card[cardIDLen:splitCharInd])
	numsStr := strings.TrimSpace(card[splitCharInd+1:])
	winningNums := getInts(winningNumsStr)
	nums := getInts(numsStr)

	fmt.Println(winningNums, nums)
	numsMatched := lenOfIntersection(winningNums, nums)
	fmt.Println("Total:", numsMatched)
	return int(math.Pow(2, float64(numsMatched-1)))
}

func lenOfIntersection(a1, a2 []int) int {
	A := map[int]bool{}
	B := map[int]bool{}

	for _, x := range a1 {
		A[x] = true
	}
	for _, x := range a2 {
		B[x] = true
	}
	lenOfInter := 0

	for key := range A {
		if B[key] {
			fmt.Println("Got", key)
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
