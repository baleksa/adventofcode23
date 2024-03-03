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
	fmt.Println(sumOfGearRatios(sch))
}

func sumOfGearRatios(sch [][]rune) int {
	sum := 0
	for i := range sch {
		for j := range sch[i] {
			c := sch[i][j]
			if c == '*' {
				ratio, isGear := checkStar(sch, i, j)
				if isGear {
					// fmt.Printf("Line:%d Column:%d Ratio=%d\n", i+1, j+1, ratio)
					sum += ratio
				}
			}
		}
	}
	return sum
}

func checkStar(sch [][]rune, i, j int) (ratio int, isGear bool) {
	allNums := []int{}

	if i > 0 {
		allNums = append(allNums, parseLine(sch[i-1], j)...)
	}
	if i < len(sch)-1 {
		allNums = append(allNums, parseLine(sch[i+1], j)...)
	}
	if j > 0 {
		allNums = append(allNums, parseLine(sch[i][0:j+1], j)...)
	}
	if j < len(sch[i])-1 {
		allNums = append(allNums, parseLine(sch[i][j:], 0)...)
	}

	if len(allNums) != 2 {
		return ratio, false
	}

	isGear = true
	ratio = allNums[0] * allNums[1]
	return
}

func parseLine(line []rune, i int) []int {
	nums := []int{}
	ystart := i - 1

	if ystart < 0 {
		ystart = 0
	}
	yfinish := i + 2
	if yfinish > len(line) {
		yfinish = len(line)
	}

	for y := ystart; y < yfinish; {
		if !unicode.IsDigit(line[y]) {
			y++
			continue
		}
		num, j := parseNum(line, y)
		nums = append(nums, num)
		y = j
	}

	return nums
}

func parseNum(line []rune, i int) (int, int) {
	j := i
	for i >= 0 && unicode.IsDigit(line[i]) {
		i--
	}
	for j < len(line) && unicode.IsDigit(line[j]) {
		j++
	}

	num, err := strconv.Atoi(string(line[i+1 : j]))
	if err != nil {
		log.Fatal(err)
	}

	return num, j
}
