package day03_test

import (
	"testing"
	"yrahul3910/aoc-2025-go/day03"
	"yrahul3910/aoc-2025-go/utils"
)

func TestSolvePuzzle1(t *testing.T) {
	input := utils.ReadInput(t)
	result := day03.SolvePuzzle1(input)
	utils.LogResult(t, result)
}

func TestSolvePuzzle2(t *testing.T) {
	input := utils.ReadInput(t)
	result := day03.SolvePuzzle2(input)
	utils.LogResultBigInt(t, result)
}
