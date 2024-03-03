package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	input_fpath := os.Args[1]
	input_f, _ := os.Open(input_fpath)
	in := bufio.NewScanner(input_f)
	sum := 0
	for in.Scan() {
		line := in.Text()
		sum += CountArrangementsSlow(line)
	}
	fmt.Println("Sum =>", sum)
}

func CountArrangementsSlow(line string) int {
	parts := strings.Fields(line)
	record := []rune(parts[0])
	cont_groups := []int{}
	for _, s := range strings.Split(parts[1], ",") {
		num, _ := strconv.Atoi(s)
		cont_groups = append(cont_groups, num)
	}

	regexpStr := fmt.Sprintf(`^\.*#{%d}`, cont_groups[0])
	regexpStr += fmt.Sprint()
	for _, g := range cont_groups[1:] {
		regexpStr += fmt.Sprintf(`\.+#{%d}`, g)
	}
	regexpStr += `\.*$`

	// fmt.Println("Damaged record =>", string(record))
	// fmt.Println("Continuous groups =>", cont_groups)
	// fmt.Println("Regexp =>", regexpStr)

	validArrangement := regexp.MustCompile(regexpStr)
	cnt := countArrangementsSlow(record, 0, validArrangement)

	// fmt.Println("Valid arrangements =>", cnt)
	// fmt.Println("=====================")

	return cnt
}

func countArrangementsSlow(record []rune, i int, r *regexp.Regexp) int {
	if i == len(record) {
		// fmt.Println("Record =>", string(record))
		// fmt.Println("Regexp =>", r.String())
		if r.Match([]byte(string(record))) {
			return 1
		} else {
			return 0
		}
	}
	if record[i] != '?' {
		return countArrangementsSlow(record, i+1, r)
	}

	record[i] = '#'
	res := countArrangementsSlow(record, i+1, r)
	record[i] = '.'
	res += countArrangementsSlow(record, i+1, r)
	record[i] = '?'
	return res
}

func CountArrangements(line string) int {
	parts := strings.Fields(line)
	record := []rune(parts[0])
	cont_groups := []int{}
	for _, s := range strings.Split(parts[1], ",") {
		num, _ := strconv.Atoi(s)
		cont_groups = append(cont_groups, num)
	}
	return countArrangements(record, cont_groups)
}

func countArrangements(record []rune, cg []int) int {
	if len(record) == 0 {
		if len(cg) == 0 {
			return 1
		} else {
			return 0
		}
	}
	if len(cg) == 0 {
		return 1
	}
	if record[0] == '.' {
		return countArrangements(record[1:], cg)
	}
	if record[0] == '#' {
		cg[0]--
		if cg[0] == 0 {
			cg = cg[1:]
		}
		return countArrangements(record[1:], cg)
	}
	// record[0] == ?
	return 0
}
