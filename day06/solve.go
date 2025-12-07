package day06

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
	"yrahul3910/aoc-2025-go/utils"
)

// make an "enum"
type Operator int

// what even is this
const (
	Add Operator = iota
	Multiply
)

var OpString = map[Operator]string{
	Add:      "+",
	Multiply: "*",
}

var StringOp = map[string]Operator{
	"+": Add,
	"*": Multiply,
}

func (o Operator) String() string {
	return OpString[o]
}

type Problem struct {
	operands []int64
	op       Operator
}

func (problem Problem) String() string {
	s := "Problem {\n"
	for i, op := range problem.operands {
		s += fmt.Sprintf("%d ", op)

		if i < len(problem.operands)-1 {
			s += fmt.Sprintf("%s ", problem.op)
		}
	}
	s += "\n}"

	return s
}

func PrintProblems(problems []Problem) {
	utils.PrintArray(problems)
	fmt.Printf("#problems: %d\n", len(problems))
}

func ParseInputPart1(input string) []Problem {
	lines := strings.Split(input, "\n")
	problemStrings := make([][]string, 0)

	for _, line := range lines {
		if len(strings.TrimSpace(line)) == 0 {
			continue
		}
		curOperands := strings.Fields(line)
		problemStrings = append(problemStrings, curOperands)
	}

	problems := make([]Problem, 0)
	numProblems := len(problemStrings[0])
	numOperands := len(problemStrings)

	for j := range numProblems {
		curOperands := make([]int64, 0)

		for i := range numOperands - 1 {
			curOperand, err := strconv.ParseInt(problemStrings[i][j], 10, 64)
			if err != nil {
				panic(err)
			}

			curOperands = append(curOperands, curOperand)
		}
		curOp := StringOp[problemStrings[numOperands-1][j]]
		problems = append(problems, Problem{curOperands, curOp})
	}

	return problems
}

func EvaluateProblem(problem Problem) int64 {
	result := int64(0)
	if problem.op == Multiply {
		result++
	}

	for _, op := range problem.operands {
		switch problem.op {
		case Add:
			result += op
		case Multiply:
			result *= op
		}
	}

	return result
}

func SolvePuzzle1(input string) int {
	problems := ParseInputPart1(input)
	PrintProblems(problems)
	finalOperands := make([]int64, 0)

	for _, problem := range problems {
		finalOperands = append(finalOperands, EvaluateProblem(problem))
	}

	return int(EvaluateProblem(Problem{finalOperands, Add}))
}

func ParseInputPart2(input string) []Problem {
	lines := strings.Split(input, "\n")
	problemStrings := make([][]rune, 0)

	for _, line := range lines {
		if len(strings.TrimSpace(line)) == 0 {
			continue
		}
		curOperands := []rune(line)
		problemStrings = append(problemStrings, curOperands)
	}

	// pre-compute indices of problem operators
	opIndices := make([]int, 0)
	sepIndices := make([]int, 0) // "separator" columns
	for i, r := range problemStrings[len(problemStrings)-1] {
		if r != ' ' {
			opIndices = append(opIndices, i)
			if i > 0 {
				sepIndices = append(sepIndices, i-1)
			}
		}
	}

	problems := make([]Problem, 0)
	curOperands := make([]int64, 0)

	for j := len(problemStrings[0]) - 1; j >= 0; j-- {
		if slices.Contains(sepIndices, j) {
			continue
		}

		curOperand := int64(0)

		for i := range len(problemStrings) - 1 {
			curDigit := 0
			if problemStrings[i][j] != ' ' {
				curDigit = int(problemStrings[i][j] - '0')
				curOperand = curOperand*10 + int64(curDigit)
			}
		}
		curOperands = append(curOperands, curOperand)

		if slices.Contains(opIndices, j) {
			curOp := StringOp[string(problemStrings[len(problemStrings)-1][j])]
			problems = append(problems, Problem{slices.Clone(curOperands), curOp})

			curOperands = slices.Delete(curOperands, 0, len(curOperands))
		}
	}

	return problems
}

func SolvePuzzle2(input string) int {
	problems := ParseInputPart2(input)
	PrintProblems(problems)
	finalOperands := make([]int64, 0)

	for _, problem := range problems {
		finalOperands = append(finalOperands, EvaluateProblem(problem))
	}

	return int(EvaluateProblem(Problem{finalOperands, Add}))
}
