package day03

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func PrintArray(arr []int) {
	for _, i := range arr {
		fmt.Print(i)
	}
	fmt.Println()
}

func ParseLine(line string) []int {
	batteries := make([]int, 0, len(line))

	for _, r := range line {
		val, err := strconv.ParseInt(string(r), 10, 8)
		if err != nil {
			panic(err)
		}

		batteries = append(batteries, int(val))
	}

	return batteries
}

func HighestJoltage(arr []int) int {
	n := len(arr)

	tens := 0

	for i := range n - 1 {
		if arr[i] > arr[tens] {
			tens = i
		}
	}

	ones := arr[tens+1]

	for i := tens + 1; i < n; i++ {
		if arr[i] > ones {
			ones = arr[i]
		}
	}

	return arr[tens]*10 + ones
}

func SolvePuzzle1(input string) int {
	lines := strings.Split(input, "\n")
	sum := 0

	for _, line := range lines {
		if len(strings.TrimSpace(line)) == 0 {
			break
		}

		batteries := ParseLine(line)
		PrintArray(batteries)
		cur := HighestJoltage(batteries)
		fmt.Println(cur)

		sum += cur
	}

	return sum
}

func HighestJoltageTwelveDigit(arr []int) uint64 {
	const DIGITS = 12
	n := len(arr)

	indices := make([]int, DIGITS)
	indices[0] = 0

	for k := range DIGITS {
		remaining := DIGITS - 1 - k

		if k > 0 {
			indices[k] = indices[k-1] + 1
		}

		for i := indices[k]; i < n-remaining; i++ {
			if arr[i] > arr[indices[k]] {
				indices[k] = i
			}
		}
	}

	sum := uint64(0)
	for i := range uint64(DIGITS) {
		exp := uint64(DIGITS-1) - i
		sum += uint64(arr[indices[i]]) * uint64(math.Pow10(int(exp)))
	}

	return sum
}

func SolvePuzzle2(input string) uint64 {
	lines := strings.Split(input, "\n")
	sum := uint64(0)

	for _, line := range lines {
		if len(strings.TrimSpace(line)) == 0 {
			break
		}

		batteries := ParseLine(line)
		PrintArray(batteries)
		cur := HighestJoltageTwelveDigit(batteries)
		fmt.Println(cur)

		sum += cur
	}

	return sum
}
