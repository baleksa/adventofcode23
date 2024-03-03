package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Race struct {
	time, distance int
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	io := bufio.NewScanner(f)
	io.Scan()
	timesStr := io.Text()
	io.Scan()
	distStr := io.Text()

	times := strings.Fields(timesStr)[1:]
	dists := strings.Fields(distStr)[1:]

	fmt.Println(times)
	fmt.Println(dists)
	races := []Race{}
	for i, timeStr := range times {
		var time, dist int
		fmt.Sscanf(timeStr, "%d", &time)
		fmt.Sscanf(dists[i], "%d", &dist)
		races = append(races, Race{time, dist})
	}

	fmt.Println(races)

	totalProd := 1
	for _, race := range races {
		totalProd *= numOfWays(race)
	}

	fmt.Println(totalProd)

}

func numOfWays(race Race) int {
	time := race.time
	dist := race.distance
	n := 0
	for i := 1; i < time; i++ {
		idist := (time - i) * i
		if idist > dist {
			n++
		}
	}
	return n
}
