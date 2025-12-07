package day05

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
	"yrahul3910/aoc-2025-go/utils"
)

type _range struct {
	start int64
	end   int64
}

func (r _range) String() string {
	return fmt.Sprintf("[%d, %d]", r.start, r.end)
}

func isFresh(ranges []_range, val int) bool {
	for _, curRange := range ranges {
		if int(curRange.start) <= val && val <= int(curRange.end) {
			return true
		}
	}

	return false
}

/** Get sorted, merged ranges. Also returns the index of the line in `inputs` that contains the first ingredient. */
func GetSortedRanges(input string) ([]_range, int) {
	lines := strings.Split(input, "\n")
	ranges := make([]_range, 0)
	i := 0

	for len(lines[i]) > 0 {
		parts := strings.Split(lines[i], "-")
		start, err := strconv.ParseInt(parts[0], 10, 64)
		if err != nil {
			panic(err)
		}

		end, err := strconv.ParseInt(parts[1], 10, 64)
		if err != nil {
			panic(err)
		}

		ranges = append(ranges, _range{start, end})
		i++
	}
	i++ // skip the blank line

	// sort ranges
	slices.SortFunc(ranges, func(a, b _range) int {
		return int(a.start) - int(b.start)
	})

	sortedRanges := make([]_range, 0)
	sortedRanges = append(sortedRanges, ranges[0])
	ct := 1 // current count of sortedRanges

	// merge ranges
	for j := 1; j < len(ranges); j++ {
		curRange := ranges[j]

		if curRange.start > sortedRanges[ct-1].end+1 {
			sortedRanges = append(sortedRanges, curRange)
			ct++

			continue
		}

		sortedRanges[ct-1].end = max(curRange.end, sortedRanges[ct-1].end)
	}
	fmt.Print("Final sorted ranges are: ")

	return sortedRanges, i
}

func SolvePuzzle1(input string) int {
	lines := strings.Split(input, "\n")

	sortedRanges, i := GetSortedRanges(input)
	fresh := 0

	for len(lines[i]) > 0 {
		val, err := strconv.ParseInt(lines[i], 10, 64)
		if err != nil {
			panic(err)
		}

		if isFresh(sortedRanges, int(val)) {
			fresh++
		}

		i++
	}

	return fresh
}

func SolvePuzzle2(input string) int {
	sortedRanges, _ := GetSortedRanges(input)
	utils.PrintArray(sortedRanges)
	ct := 0

	for _, curRange := range sortedRanges {
		ct += int(curRange.end) - int(curRange.start) + 1
	}

	return ct
}
