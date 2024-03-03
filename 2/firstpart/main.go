package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		return
	}
	defer f.Close()

	in := bufio.NewScanner(f)
	sumId := 0
	for in.Scan() {
		line := in.Text()
		id, possible := checkLine(line)
		if possible {
			sumId += id
		}
	}
	fmt.Println("Sum of IDs:", sumId)
}

var Cnt = [...]int{12, 13, 14}

const (
	RED = iota
	GREEN
	BLUE
)

func checkLine(line string) (int, bool) {
	id, maxr, maxg, maxb := parse(line)
	if maxr <= Cnt[RED] && maxg <= Cnt[GREEN] && maxb <= Cnt[BLUE] {
		return id, true
	}
	return id, false
}

func parse(line string) (int, int, int, int) {
	parts := strings.Split(line, ":")
	gameIdStr := parts[0]
	bagTakes := parts[1]
	id, _ := strconv.Atoi(strings.Split(gameIdStr, " ")[1])

	var maxr, maxg, maxb int
	for _, take := range strings.Split(bagTakes, ";") {
		for _, colorCnt := range strings.Split(take, ",") {
			parts := strings.Split(strings.TrimSpace(colorCnt), " ")
			cnt, _ := strconv.Atoi(parts[0])
			color := parts[1]
			switch color {
			case "red":
				if cnt > maxr {
					maxr = cnt
				}
			case "green":
				if cnt > maxg {
					maxg = cnt
				}
			case "blue":
				if cnt > maxb {
					maxb = cnt
				}
			}
		}
	}
	return id, maxr, maxg, maxb
}
