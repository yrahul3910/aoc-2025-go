package day04

import (
	"strings"
)

type index struct {
	x int
	y int
}

func ValidateIndex(i int, j int, r int, c int) bool {
	return (i >= 0 && i < r && j >= 0 && j < c)
}

func Neighbors(i int, j int, r int, c int) []index {
	neighbors := make([]index, 0, 8) // Max 8 neighbors

	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			if dx == 0 && dy == 0 {
				continue
			}
			if ValidateIndex(i+dx, j+dy, r, c) {
				neighbors = append(neighbors, index{i + dx, j + dy})
			}
		}
	}

	return neighbors
}

// Is it removeable or removable? Hmmm....
func GetRemovableBoxes(arr []string) []index {
	r := len(arr)
	c := len(arr[0])

	removableBoxes := make([]index, 0, r*c)

	for i := range r {
		if len(strings.TrimSpace(arr[i])) == 0 {
			continue
		}

		for j := range c {
			if arr[i][j] != '@' {
				continue
			}

			neighbors := Neighbors(i, j, r, c)

			curCount := 0
			for _, n := range neighbors {
				if len(strings.TrimSpace(arr[n.x])) == 0 {
					continue
				}

				if arr[n.x][n.y] == '@' {
					curCount++
				}
			}

			if curCount < 4 { // Problem constraint
				removableBoxes = append(removableBoxes, index{i, j})
			}
		}
	}

	return removableBoxes
}

func SolvePuzzle1(input string) int {
	lines := strings.Split(input, "\n")
	return len(GetRemovableBoxes(lines))
}

func SolvePuzzle2(input string) int {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	totalRemoved := 0

	for true {
		removableIndices := GetRemovableBoxes(lines)
		removed := len(removableIndices)
		totalRemoved += removed

		for _, idx := range removableIndices {
			bytes := []byte(lines[idx.x])
			bytes[idx.y] = 'x'
			lines[idx.x] = string(bytes)
		}

		if removed == 0 {
			break
		}
	}

	return totalRemoved
}
