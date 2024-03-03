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
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Error opening input file.")
	}
	defer file.Close()
	in := bufio.NewScanner(file)
	sum := 0
	var f, l int
	for in.Scan() {
		line := []rune(in.Text())
		for _, c := range line {
			if unicode.IsDigit(c) {
				f, err = strconv.Atoi(string(c))
				break
			}
		}
		for i := len(line) - 1; i >= 0; i-- {
			if unicode.IsDigit(line[i]) {
				l, err = strconv.Atoi(string(line[i]))
				break
			}
		}
		sum += 10*f + l
	}
	fmt.Println("Sum:", sum)
}
