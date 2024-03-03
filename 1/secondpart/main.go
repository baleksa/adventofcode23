package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var digits = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Error opening input file.")
	}
	defer file.Close()
	in := bufio.NewScanner(file)
	sum := 0
	for in.Scan() {
		sum += getValue(in.Text())
	}
	fmt.Println("Sum:", sum)
}

func getValue(line string) int {
	value := getFirstValue(line)*10 + getLastValue(line)
	return value
}

func getFirstValue(line string) int {
	indw, wd, hasw := getFirstWrittenDigit(line)
	indd, dd, hasd := getFirstDigit(line)
	if !hasw {
		return dd
	}
	if !hasd {
		return wd
	}
	if indw < indd {
		return wd
	}
	return dd
}

func getFirstDigit(line string) (ind, dig int, has bool) {
	for i, c := range line {
		if unicode.IsDigit(c) {
			dig, _ = strconv.Atoi(string(c))
			ind = i
			has = true
			return
		}
	}
	return
}

func getFirstWrittenDigit(line string) (ind, dig int, has bool) {
	ind = -1
	for i, digitWord := range digits {
		cInd := strings.Index(line, digitWord)
		if cInd != -1 && (ind == -1 || cInd < ind) {
			has = true
			dig = i + 1
			ind = cInd
		}
	}
	return
}

func getLastValue(line string) int {
	indw, wd, hasw := getLastWrittenDigit(line)
	indd, dd, hasd := getLastDigit(line)
	if !hasw {
		return dd
	}
	if !hasd {
		return wd
	}
	if indw > indd {
		return wd
	}
	return dd
}

func getLastWrittenDigit(line string) (ind, dig int, has bool) {
	ind = -1
	for i, digitWord := range digits {
		cInd := strings.LastIndex(line, digitWord)
		if cInd != -1 && (ind == -1 || cInd > ind) {
			has = true
			dig = i + 1
			ind = cInd
		}
	}
	return
}

func getLastDigit(line string) (ind, dig int, has bool) {
	lineRunes := []rune(line)
	for i := len(lineRunes) - 1; i >= 0; i-- {
		c := lineRunes[i]
		if unicode.IsDigit(c) {
			dig, _ = strconv.Atoi(string(c))
			ind = i
			has = true
			return
		}
	}
	return
}
