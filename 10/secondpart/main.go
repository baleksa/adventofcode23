package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

type Map []string
type LoopTileMark [][]bool
type Point struct {
	x, y int
}

var S_NODE = map[string]byte{
	"input/test_input1": 'F',
	"input/test_input2": 'F',
	"input/test_input3": 'F',
	"input/test_input4": 'F',
	"input/test_input5": 'F',
	"input/test_input6": '7',
	"input/test_input7": '7',
	"input/input.txt":   'J'}

func main() {
	input_file_path := os.Args[1]
	if _, ok := S_NODE[input_file_path]; !ok {
		log.Fatal("No info about the start node.\nFind what tile is the start node. I was too lazy so i hardcoded values in a map.")
	}
	input_file, _ := os.Open(input_file_path)
	in := bufio.NewScanner(input_file)
	diag := Map{}
	for in.Scan() {
		diag = append(diag, in.Text())
	}
	// fmt.Print(" ")
	// for i := range diag[0] {
	// 	fmt.Print(i)
	// }
	// fmt.Println()
	// for i, line := range diag {
	// 	fmt.Print(i)
	// 	fmt.Println(line)
	// }

	fmt.Println("Tiles enclosed by the loop =>", countEnclosedTiles(diag))
}

func isVerticlaBoundary(r rune) bool {
	return r == '|' || r == 'L' || r == 'J' || r == '7' || r == 'F' || r == 'S'
}

func countEnclosedTiles(diag Map) int {
	start := findStart(diag)
	nextTiles := allConnectedTo(start, diag)
	left, right := nextTiles[0], nextTiles[1]
	isLoopTile := markLoopNodes(start, left, right, diag)
	diag[start.x] = strings.Replace(diag[start.x], "S", string(S_NODE[os.Args[1]]), 1)
	// for i := range isLoopTile {
	// 	for j := range isLoopTile[i] {
	// 		if isLoopTile[i][j] {
	// 			fmt.Print(1)
	// 		} else {
	// 			fmt.Print(0)
	// 		}
	// 	}
	// 	fmt.Println()
	// }
	nInLoopTiles := 0
	for i := range diag {
		insideOfLoop := false
		for j := 0; j < len(diag[i]); {
			if isLoopTile[i][j] {
				if diag[i][j] == '|' {
					insideOfLoop = !insideOfLoop
					j++
					continue
				}
				boundary_start := diag[i][j]
				j++
				for j < len(diag[i]) && diag[i][j] == '-' {
					j++
				}
				if j == len(diag[i]) || j == len(diag[i])-1 {
					continue
				}
				boundary_end := diag[i][j]

				j++

				if boundary_start == 'F' && boundary_end == '7' {
					continue
				}
				if boundary_start == 'L' && boundary_end == 'J' {
					continue
				}
				if boundary_start == 'F' && boundary_end == 'J' {
					insideOfLoop = !insideOfLoop
				}
				if boundary_start == 'L' && boundary_end == '7' {
					insideOfLoop = !insideOfLoop
				}
				continue
			}

			if insideOfLoop {
				// fmt.Println("In loop =>", Point{i, j}, string([]rune(diag[i])[j]))
				nInLoopTiles++
			}
			j++
		}
	}

	return nInLoopTiles
}

func markLoopNodes(start, left, right Point, diag Map) LoopTileMark {
	isLoopTile := make(LoopTileMark, len(diag))
	for i := range isLoopTile {
		isLoopTile[i] = make([]bool, len(diag[i]))
	}
	isLoopTile[start.x][start.y] = true
	prevLeft := start
	prevRight := start
	for left != right {
		isLoopTile[left.x][left.y] = true
		isLoopTile[right.x][right.y] = true
		prevLeft, left = left, next(prevLeft, left, diag)
		prevRight, right = right, next(prevRight, right, diag)
	}
	if left == right {
		isLoopTile[left.x][left.y] = true
	}

	return isLoopTile
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
