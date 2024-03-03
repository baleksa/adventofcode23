package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

type Map []string
type Point struct {
	x, y int
}

func main() {
	input_file_path := os.Args[1]
	input_file, _ := os.Open(input_file_path)
	in := bufio.NewScanner(input_file)
	diag := Map{}
	for in.Scan() {
		diag = append(diag, in.Text())
	}
	// for _, line := range diag {
	// 	fmt.Println(line)
	// }
	start := findStart(diag)
	fmt.Println("Start =>", start)
	nextNodes := allConnectedTo(start, diag)
	fmt.Println("Nodes connected to start node =>", nextNodes)
	l, r := nextNodes[0], nextNodes[1]
	fmt.Println("Max distance =>", maxDistance(start, l, r, diag))
}

func maxDistance(start, left, right Point, diag Map) int {
	dist := 1
	prevLeft := start
	prevRight := start
	for left != right {
		prevLeft, left = left, next(prevLeft, left, diag)
		prevRight, right = right, next(prevRight, right, diag)
		dist++
	}
	if left == right {
		return dist
	}
	return dist - 1
}

func next(prevRight, right Point, diag Map) Point {
	for _, node := range allConnections(right, diag) {
		if node != prevRight {
			return node
		}
	}
	panic("unreachable")
}

func findStart(diag Map) Point {
	for i, line := range diag {
		for j, point := range line {
			if point == 'S' {
				return Point{i, j}
			}
		}
	}
	panic("unreachable")
}

func allConnectedTo(p Point, diag Map) []Point {
	js := []int{-1, 0, 1}
	is := []int{-1, 0, 1}

	result := []Point{}
	for _, i := range is {
		for _, j := range js {
			if i == 0 && j == 0 {
				continue
			}
			x := p.x + i
			y := p.y + j
			if x >= 0 && x < len(diag) && y >= 0 && y < len(diag[x]) {
				point := Point{x, y}
				if slices.Contains(allConnections(point, diag), p) {
					result = append(result, point)
				}
			}
		}
	}
	return result
}

func allConnections(p Point, diag Map) []Point {
	connections := []Point{}
	poss_conn := []Point{}
	tile := diag[p.x][p.y]
	switch tile {
	case '|':
		poss_conn = append(poss_conn, Point{p.x - 1, p.y}, Point{p.x + 1, p.y})
	case '-':
		poss_conn = append(poss_conn, Point{p.x, p.y - 1}, Point{p.x, p.y + 1})
	case 'L':
		poss_conn = append(poss_conn, Point{p.x - 1, p.y}, Point{p.x, p.y + 1})
	case 'J':
		poss_conn = append(poss_conn, Point{p.x - 1, p.y}, Point{p.x, p.y - 1})
	case '7':
		poss_conn = append(poss_conn, Point{p.x + 1, p.y}, Point{p.x, p.y - 1})
	case 'F':
		poss_conn = append(poss_conn, Point{p.x + 1, p.y}, Point{p.x, p.y + 1})
	}
	for _, point := range poss_conn {
		if point.x >= 0 && point.x < len(diag) && point.y >= 0 && point.y < len(diag[point.x]) {
			connections = append(connections, Point{point.x, point.y})
		}
	}
	return connections
}
