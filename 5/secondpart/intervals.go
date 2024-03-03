package main

import (
	"fmt"
	"sort"
)

type Interval struct {
	// Interval [start,end)
	start, end uint64
}

type Relation int

const (
	ACONTAINSB Relation = iota + 1
	AINSIDEOFB
	AOVERLEFTB
	AOVERRIGHTB
	NOOVERLAP
	SAME
)

func relation(a, b Interval) Relation {
	if (a.end <= b.start) || (a.start >= b.end) {
		return NOOVERLAP
	}
	// if a.start == b.start && a.end == b.end {
	// 	return SAME
	// }
	if a.start <= b.start && b.end <= a.end {
		return ACONTAINSB
	}
	if b.start <= a.start && a.end <= b.end {
		return AINSIDEOFB
	}
	if a.start < b.start && b.start < a.end && a.end <= b.end {
		return AOVERLEFTB
	}
	if b.start < a.start && a.start < b.end && b.end < a.end {
		return AOVERRIGHTB
	}

	panic("This should be unreachable!")
}

func newInterval(start, width uint64) Interval {
	return Interval{start, start + width}
}

func printIntervals(a []Interval) {
	if len(a) == 0 {
		fmt.Println("{}")
		return
	}
	fmt.Print("{\n")
	printInterval(a[0])
	fmt.Print("\n")
	for _, i := range a[1:] {
		printInterval(i)
		fmt.Print("\n")
	}
	fmt.Println("}")
}

func printInterval(i Interval) {
	fmt.Printf("[%11d-%11d)", i.start, i.end)
}

func isEmpty(i Interval) bool {
	return i.start == i.end
}

func removeEmptyIntervals(s []Interval) []Interval {
	if s == nil || len(s) == 0 {
		return s
	}
	emptyInds := []int{}
	for i, x := range s {
		if isEmpty(x) {
			emptyInds = append(emptyInds, i)
		}
	}

	reverse(emptyInds)

	for _, i := range emptyInds {
		s = remove(s, i)
	}

	return s
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func remove(s []Interval, i int) []Interval {
	if i < 0 || i > len(s)-1 {
		return s
	}
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func minValue(inters []Interval) uint64 {
	if len(inters) == 0 {
		return 0
	}
	minv := inters[0].start
	for _, i := range inters[1:] {
		minv = min(minv, i.start)
	}
	return minv
}

func unionOfIntervals(intervals []Interval) []Interval {
	if len(intervals) < 2 {
		return intervals
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].start < intervals[j].start ||
			(intervals[i].start == intervals[j].start && intervals[i].end < intervals[j].end)
	})
	// fmt.Println("Before union:")
	// printIntervals(intervals)

	j := 0
	for i := 0; i < len(intervals); {
		x := intervals[i]
		var k int
		for k = i + 1; k < len(intervals)-1 && intervals[k].start <= x.end; k++ {
			x.end = intervals[k].end
		}
		intervals[j] = x
		j = j + 1
		i = k
	}
	// if j != len(intervals) {
	// 	fmt.Println("After union:")
	// 	printIntervals(intervals[:j])
	//
	// }
	return intervals[:j]
}
