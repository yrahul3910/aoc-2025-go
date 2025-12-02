package day02

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func SolvePuzzle1(input string) int {
	ranges := strings.Split(input, ",")
	sum := 0

	for _, id_range := range ranges {
		range_bounds := strings.Split(id_range, "-")

		range_start, err := strconv.ParseInt(strings.TrimSpace(range_bounds[0]), 10, 64)
		if err != nil {
			panic(err)
		}

		range_end, err := strconv.ParseInt(strings.TrimSpace(range_bounds[1]), 10, 64)
		if err != nil {
			panic(err)
		}

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

			if j%dividend == 0 {
				sum += int(j)
			}
		}
	}

	return sum
}

func SolvePuzzle2(input string) int {
	// TODO: solve puzzle 2
	return 0
}
