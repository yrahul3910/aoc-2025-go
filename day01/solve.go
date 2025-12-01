package day01

import (
	"fmt"
	"strconv"
	"strings"
)

// Perform negative number-aware n % m
func mod(n int, m int) int {
	if n < 0 {
		return mod(n+m, m)
	}
	return n % m
}

func SolvePuzzle1(input string) int {
	// Dial starts at 50
	value := 50
	lines := string(input)
	password := 0

	for line := range strings.Lines(lines) {
		line := strings.Trim(line, "\r\n")
		dir := line[0]
		val := strings.TrimPrefix(line, string(dir))

		amount, err := strconv.Atoi(val)
		if err != nil {
			panic(err)
		}

		switch dir {
		case 'L':
			value = mod(value-amount, 100)
		case 'R':
			value = mod(value+amount, 100)
		default:
			panic(fmt.Sprintf("Invalid direction: %s\n", string(dir)))
		}

		if value == 0 {
			password += 1
		}
	}

	return password
}

func SolvePuzzle2(input string) int {
	// Dial starts at 50
	value := 50
	lines := string(input)
	password := 0

	for line := range strings.Lines(lines) {
		line := strings.Trim(line, "\r\n")
		dir := line[0]
		val := strings.TrimPrefix(line, string(dir))

		amount, err := strconv.Atoi(val)
		if err != nil {
			panic(err)
		}

		password += amount / 100
		amount = amount % 100

		switch dir {
		case 'L':
			if value > 0 && value-amount < 0 {
				// We cross 0
				password += 1
			}
			value = mod(value-amount, 100)
		case 'R':
			if value+amount > 100 {
				password += 1
			}
			value = mod(value+amount, 100)
		default:
			panic(fmt.Sprintf("Invalid direction: %s\n", string(dir)))
		}

		if value == 0 {
			password += 1
		}
	}

	return password
}
