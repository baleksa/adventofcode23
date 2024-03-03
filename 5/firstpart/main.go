package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

type Mapping struct {
	dst, src, width uint64
}

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	in := bufio.NewScanner(f)

	seeds := getSeeds(in)
	inter := seeds
	// fmt.Println(inter)

	applied := map[int]bool{}
	var mapping Mapping
	for in.Scan() {
		line := in.Text()
		var src, dst, width uint64
		if n, _ := fmt.Sscanf(line, "%d %d %d", &dst, &src, &width); n != 3 {
			applied = make(map[int]bool)
			continue
		}
		mapping = Mapping{dst, src, width}
		// fmt.Println(mapping)
		for i := range inter {
			if applied[i] {
				continue
			}
			y := applyMapping(inter[i], mapping)
			if y != inter[i] {
				inter[i] = y
				applied[i] = true
			}
		}
	}

	fmt.Println(slices.Min(inter))

}

func applyMapping(x uint64, mapping Mapping) uint64 {
	y := x
	if x >= mapping.src && x < mapping.src+mapping.width {
		y = mapping.dst + (x - mapping.src)
	}

	return y
}

func getSeeds(in *bufio.Scanner) []uint64 {
	seeds := []uint64{}
	in.Scan()
	line := in.Text()
	for _, s := range strings.Fields(line)[1:] {
		var num uint64
		if n, err := fmt.Sscanf(s, "%d", &num); n != 1 {
			log.Fatal(err)
		}
		seeds = append(seeds, num)
	}
	return seeds
}
