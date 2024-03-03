package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input_fpath := os.Args[1]
	foldCoef, _ := strconv.Atoi(os.Args[2])
	input_f, _ := os.Open(input_fpath)
	in := bufio.NewScanner(input_f)
	sum := 0
	for in.Scan() {
		line := in.Text()
		sum += CountArrangements(line, foldCoef)
	}
	fmt.Println("Sum =>", sum)
}

func CountArrangements(line string, foldCoef int) int {
	parts := strings.Fields(line)
	record := []rune(parts[0])
	cont_groups := []int{}
	for _, s := range strings.Split(parts[1], ",") {
		num, _ := strconv.Atoi(s)
		cont_groups = append(cont_groups, num)
	}
	if foldCoef > 1 {
		tmp := make([]int, 0, foldCoef*len(cont_groups))
		for i := 0; i < foldCoef; i++ {
			tmp = append(tmp, cont_groups...)
		}
		cont_groups = tmp

		tmp2 := make([]rune, 0, foldCoef*(len(record)+1))
		tmp2 = append(tmp2, record...)
		for i := 1; i < foldCoef; i++ {
			tmp2 = append(tmp2, '?')
			tmp2 = append(tmp2, record...)
		}
		record = tmp2
	}

	// fmt.Println("Damaged record =>", string(record))
	// fmt.Println("Continuous groups =>", cont_groups)

	return CountArrangementsFast(record, cont_groups)
}

type Memo map[string]int

var memo Memo

func stringCg(cg []int) string {
	s := ""
	for _, x := range cg {
		s += fmt.Sprint(x)
	}
	return s
}
func stringInGroup(ig bool) string {
	if ig {
		return "1"
	}
	return "0"
}

func Key(record []rune, cg []int, inGroup bool) string {
	return string(record) + stringCg(cg) + stringInGroup(inGroup)
}

func CountArrangementsFast(record []rune, cg []int) int {
	memo = Memo{}
	res := countArrangementsFast(record, cg, false)
	// fmt.Println("Valid arrangements =>", res)
	return res
}

func countArrangementsFast(record []rune, cg []int, inGroup bool) int {
	key := Key(record, cg, inGroup)
	res, ok := memo[key]
	if ok {
		return res
	}
	if len(record) == 0 {
		if len(cg) == 0 || (len(cg) == 1 && cg[0] == 0) {
			return 1
		} else {
			return 0
		}
	}

	s := record[0]
	if s == '.' {
		if len(cg) == 0 {
			res = countArrangementsFast(record[1:], cg, false)
			memo[key] = res
			return res
		}
		if cg[0] == 0 {
			res := countArrangementsFast(record[1:], cg[1:], false)
			memo[key] = res
			return res
		}
		if inGroup {
			memo[key] = 0
			return 0
		} else {
			res := countArrangementsFast(record[1:], cg, false)
			memo[key] = res
			return res
		}
	}
	if s == '#' {
		if len(cg) == 0 || cg[0] == 0 {
			memo[key] = 0
			return 0
		}
		cg[0]--
		res := countArrangementsFast(record[1:], cg, true)
		cg[0]++
		memo[key] = res
		return res

	}
	// s == '?'
	if len(cg) == 0 {
		// s must be "."
		res := countArrangementsFast(record[1:], cg, false)
		memo[key] = res
		return res
	}
	if cg[0] == 0 {
		// s must be "."
		res := countArrangementsFast(record[1:], cg[1:], false)
		memo[key] = res
		return res
	}

	if inGroup {
		// s must be '#'
		cg[0]--
		res := countArrangementsFast(record[1:], cg, true)
		cg[0]++
		memo[key] = res
		return res
	}
	// inGroup == false
	// s can be '.'
	res = countArrangementsFast(record[1:], cg, false)
	// or '#'
	cg[0]--
	res += countArrangementsFast(record[1:], cg, true)
	cg[0]++

	memo[key] = res
	return res
}
