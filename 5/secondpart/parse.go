package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"sort"
	"strings"
)

func parseInput(r io.Reader) (seeds []Interval, functions []Function) {
	in := bufio.NewScanner(r)
	in.Scan()
	firstLine := in.Text()
	seeds = getSeeds(firstLine)
	in.Scan()

	funStr := ""
	// funName := ""
	for in.Scan() {
		line := in.Text()
		if line == "" {
			continue
		}
		if strings.Contains(line, "map") {
			if funStr != "" {
				// fmt.Println(funName)
				// fmt.Print(funStr)
				// fmt.Println("END", funName)
				functions = append(functions, getFunction(funStr))
			}
			// funName = line
			funStr = ""
			continue
		}
		funStr += line + "\n"
	}
	functions = append(functions, getFunction(funStr))
	return
}

func getSeeds(line string) []Interval {
	seeds := []Interval{}
	fields := strings.Fields(line)[1:]
	for i := 0; i < len(fields)-1; i += 2 {
		var start, width uint64
		s := strings.Join(fields[i:i+2], " ")
		if n, err := fmt.Sscanf(s, "%d %d", &start, &width); n != 2 {
			log.Fatal(err)
		}
		seeds = append(seeds, newInterval(start, width))
	}
	return seeds
}

func getFunction(funStr string) Function {
	// fmt.Println("****************")
	// fmt.Print(funStr)
	// fmt.Println("****************")
	mappingsStr := strings.Split(strings.TrimSpace(funStr), "\n")
	function := Function{}
	for _, mapStr := range mappingsStr {
		function = append(function, getMapping(mapStr))
	}
	sort.Slice(function, func(i, j int) bool {
		return function[i].x.start < function[j].x.start || (function[i].x.start == function[j].x.start && function[i].x.end < function[j].x.end)
	})

	return function
}

func getMapping(mapStr string) Mapping {
	mapping := Mapping{}
	if n, err := fmt.Sscanf(mapStr, "%d %d %d", &mapping.dst, &mapping.src, &mapping.width); n != 3 {
		log.Fatal(err)
	}
	mapping.x.start = mapping.src
	mapping.x.end = mapping.src + mapping.width
	mapping.y.start = mapping.dst
	mapping.y.end = mapping.dst + mapping.width

	return mapping
}
