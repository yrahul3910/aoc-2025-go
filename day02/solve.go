package day02

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func get_range(input string) (int, int) {
	range_bounds := strings.Split(input, "-")

	range_start, err := strconv.ParseInt(strings.TrimSpace(range_bounds[0]), 10, 64)
	if err != nil {
		panic(err)
	}

	range_end, err := strconv.ParseInt(strings.TrimSpace(range_bounds[1]), 10, 64)
	if err != nil {
		panic(err)
	}

	return int(range_start), int(range_end)
}

func SolvePuzzle1(input string) int {
	ranges := strings.Split(input, ",")
	sum := 0

	for _, id_range := range ranges {
		range_start, range_end := get_range(id_range)

		for j := range_start; j <= range_end; j++ {
			// Simple case
			if j < 100 && j%11 == 0 {
				sum += int(j)
				continue
			}

			n_digits := int(math.Floor(math.Log10(float64(j))) + 1)
			if n_digits%2 == 1 {
				continue
			}

			n_dividend_zeros := n_digits/2 - 1

			dividend, err := strconv.ParseInt(fmt.Sprintf("1%s1", strings.Repeat("0", n_dividend_zeros)), 10, 64)
			if err != nil {
				panic(err)
			}

			if j%int(dividend) == 0 {
				sum += int(j)
			}
		}
	}

	return sum
}

func has_repeated_pattern(s string) bool {
	n := len(s)
	for patternLen := 1; patternLen <= n/2; patternLen++ {
		if n%patternLen == 0 {
			pattern := s[:patternLen]
			if strings.Repeat(pattern, n/patternLen) == s {
				return true
			}
		}
	}
	return false
}

func SolvePuzzle2(input string) int {
	ranges := strings.Split(input, ",")
	sum := 0

	for _, id_range := range ranges {
		range_start, range_end := get_range(id_range)

		for j := range_start; j <= range_end; j++ {
			cur_num := fmt.Sprintf("%d", j)
			if has_repeated_pattern(cur_num) {
				sum += j
			}
		}
	}

	return sum
}
