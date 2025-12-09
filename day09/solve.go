package day09

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Coordinate struct {
	x int
	y int
}

func (c Coordinate) String() string {
	return fmt.Sprintf("(%d, %d)", c.x, c.y)
}

func ParseInput(input string) []Coordinate {
	arr := make([]Coordinate, 0)

	for line := range strings.SplitSeq(input, "\n") {
		if len(strings.TrimSpace(line)) == 0 {
			continue
		}

		parts := strings.Split(line, ",")

		x, _ := strconv.ParseInt(parts[0], 10, 64)
		y, _ := strconv.ParseInt(parts[1], 10, 64)

		arr = append(arr, Coordinate{int(x), int(y)})
	}

	return arr
}

func SolvePuzzle1(input string) int {
	coords := ParseInput(input)
	maxArea := 0

	for i := range len(coords) {
		for j := i + 1; j < len(coords); j++ {
			curX := float64(coords[i].x - coords[j].x)
			curY := float64(coords[i].y - coords[j].y)
			curArea := int(math.Abs(curX)+1) * int(math.Abs(curY)+1)

			if curArea > maxArea {
				maxArea = curArea
			}
		}
	}

	return maxArea
}

func SolvePuzzle2(input string) int {
	// TODO: solve puzzle 2
	return 0
}
