package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Race struct {
	time, distance int
}

func main() {
	f, err := os.Open("test_input.txt")
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

	time, _ := strconv.Atoi(strings.Join(times, ""))
	dist, _ := strconv.Atoi(strings.Join(dists, ""))

	fmt.Println(time)
	fmt.Println(dist)

	fmt.Println(numOfWays(Race{time, dist}))

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
