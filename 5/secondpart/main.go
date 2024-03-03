package main

import (
	"fmt"
	"log"
	"math/big"
	"os"
)

type Function []Mapping

type Mapping struct {
	x, y            Interval
	dst, src, width uint64
}

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	seeds, functions := parseInput(f)
	intervals := unionOfIntervals(seeds)
	// fmt.Println("Starting seeds intevals:")
	// printIntervals(intervals)

	for _, function := range functions {
		intervals = applyFunction(intervals, function)
		// fmt.Println("After function:")
		// printFunction(function)
		// fmt.Println("Intervals:")
		// printIntervals(intervals)
	}
	// for _, interval := range intervals {
	// 	if isEmpty(interval) {
	// 		panic("EMPTY INTERVAL!")
	// 	}
	// }
	fmt.Println(minValue(intervals))

}

func applyFunction(xIntervals []Interval, function Function) []Interval {
	// applied := map[int]bool{}
	yIntervals := []Interval{}
	unMapped := xIntervals
	for _, mapping := range function {
		newUnMapped := []Interval{}
		for _, xInterval := range unMapped {
			oldIntervals, newIntervals := applyMapping(xInterval, mapping)
			newUnMapped = unionOfIntervals(append(newUnMapped, oldIntervals...))
			yIntervals = unionOfIntervals(append(yIntervals, newIntervals...))
		}
		unMapped = newUnMapped
	}
	yIntervals = unionOfIntervals(append(yIntervals, unMapped...))
	return yIntervals
}

func applyMapping(xInterval Interval, mapping Mapping) (xIntervals, yIntervals []Interval) {
	rel := relation(xInterval, mapping.x)
	switch rel {
	// case SAME:
	// 	yIntervals = applyMappingInside(xInterval, mapping)
	case NOOVERLAP:
		xIntervals = []Interval{xInterval}
	case AINSIDEOFB:
		yIntervals = applyMappingInside(xInterval, mapping)
	case ACONTAINSB:
		xIntervals, yIntervals = applyMappingContains(xInterval, mapping)
	case AOVERLEFTB:
		xIntervals, yIntervals = applyMappingLeft(xInterval, mapping)
	case AOVERRIGHTB:
		xIntervals, yIntervals = applyMappingRight(xInterval, mapping)
	}

	return
}

func applyMappingInside(xInterval Interval, mapping Mapping) []Interval {
	yInterval := Interval{applyMappingValue(xInterval.start, mapping), applyMappingValue(xInterval.end, mapping)}
	return []Interval{yInterval}
}

func applyMappingContains(xInterval Interval, mapping Mapping) (xIntervals, yIntervals []Interval) {
	xIntervals = append(xIntervals, Interval{xInterval.start, mapping.x.start})
	xIntervals = append(xIntervals, Interval{mapping.x.end, xInterval.end})

	yIntervals = append(yIntervals, Interval{mapping.y.start, mapping.y.end})

	xIntervals, yIntervals = removeEmptyIntervals(xIntervals), removeEmptyIntervals(yIntervals)
	xIntervals, yIntervals = unionOfIntervals(xIntervals), unionOfIntervals(yIntervals)

	return
}

func applyMappingLeft(xInterval Interval, mapping Mapping) (xIntervals, yIntervals []Interval) {
	xIntervals = append(xIntervals, Interval{xInterval.start, mapping.x.start})
	yIntervals = append(yIntervals, Interval{mapping.y.start, applyMappingValue(xInterval.end, mapping)})

	xIntervals, yIntervals = removeEmptyIntervals(xIntervals), removeEmptyIntervals(yIntervals)

	xIntervals, yIntervals = unionOfIntervals(xIntervals), unionOfIntervals(yIntervals)

	return
}

func applyMappingRight(xInterval Interval, mapping Mapping) (xIntervals, yIntervals []Interval) {
	yIntervals = append(yIntervals, Interval{applyMappingValue(xInterval.start, mapping), mapping.y.end})
	xIntervals = append(xIntervals, Interval{mapping.x.end, xInterval.end})

	xIntervals, yIntervals = removeEmptyIntervals(xIntervals), removeEmptyIntervals(yIntervals)

	xIntervals, yIntervals = unionOfIntervals(xIntervals), unionOfIntervals(yIntervals)

	return
}

func applyMappingValue(x uint64, mapping Mapping) uint64 {
	var res, dst, src big.Int
	res.SetUint64(x)
	dst.SetUint64(mapping.dst)
	src.SetUint64(mapping.src)
	res.Add(&res, &dst)
	res.Sub(&res, &src)
	return res.Uint64()
}
