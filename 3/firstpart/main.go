package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

func main() {
	sch := [][]rune{}

	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	in := bufio.NewScanner(f)
	for in.Scan() {
		line := in.Text()
		sch = append(sch, []rune(line))
	}
	fmt.Println(sumOfPartNumber(sch))
}

func sumOfPartNumber(sch [][]rune) int {
	sum := 0

	for i := range sch {
		for j := 0; j < len(sch[i]); {
			c := sch[i][j]
			if !unicode.IsDigit(c) {
				j++
				continue
			}
			num, k, isPartnum := parseNum(sch, i, j)
			if isPartnum {
				// fmt.Printf("[%d/%d] %d\n", i, j, num)
				sum += num
			}
			j = k
		}
	}

	return sum
}

func parseNum(sch [][]rune, i, j int) (num, k int, isPartnum bool) {
	line := sch[i]
	for k = j; k < len(line) && unicode.IsDigit(line[k]); k++ {
	}

	num, err := strconv.Atoi(string(line[j:k]))
	if err != nil {
		log.Fatal(err)
	}

	isPartnum = checkIfPartNum(sch, i, j, k)

	return
}

func checkIfPartNum(sch [][]rune, i, j, k int) bool {
	if i > 0 && checkLine(sch[i-1], j-1, k+1) {
		return true
	}
	if i < len(sch)-1 && checkLine(sch[i+1], j-1, k+1) {
		return true
	}
	if j > 0 && checkColumn(sch, j-1, i-1, i+2) {
		return true
	}
	if k < len(sch[i]) && checkColumn(sch, k, i-1, i+2) {
		return true
	}

	return false
}

func checkColumn(sch [][]rune, j, xstart, xfinish int) bool {
	if xstart < 0 {
		xstart = 0
	}
	if xfinish > len(sch) {
		xfinish = len(sch)
	}

	y := j
	for x := xstart; x < xfinish; x++ {
		if isSymbol(sch[x][y]) {
			return true
		}
	}

	return false
}

func checkLine(line []rune, ystart, yfinish int) bool {
	if ystart < 0 {
		ystart = 0
	}
	if yfinish > len(line) {
		yfinish = len(line)
	}

	return withSymol(line[ystart:yfinish])
}

func withSymol(line []rune) bool {
	for _, c := range line {
		if isSymbol(c) {
			return true
		}
	}
	return false
}

func isSymbol(c rune) bool {
	return c != '.' && !unicode.IsDigit(c)
}
