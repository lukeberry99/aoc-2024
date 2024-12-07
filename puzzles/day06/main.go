package main

import (
	"fmt"
	"strings"

	"github.com/lukeberry99/aoc-2024/pkg/files"
	"github.com/lukeberry99/aoc-2024/pkg/ints"
)

func main() {
	fmt.Println("Part 1:", part1("input.txt"))
	fmt.Println("Part 2:", part2("input.txt"))
}

func part1(name string) int {
	input := files.ReadLines(name)

	return calcInput(input, false)
}

func calcInput(input []string, includeConcat bool) int {
	res := 0

	for _, line := range input {
		parts := strings.Split(line, ": ")
		testVal := ints.FromString(parts[0])
		numStrings := strings.Fields(parts[1])
		numbers := make([]int, len(numStrings))
		for i, numStr := range numStrings {
			numbers[i] = ints.FromString(numStr)
		}

		numOps := len(numbers) - 1
		opCombos := operatorCombos(numOps, includeConcat)

		for _, ops := range opCombos {
			if evaluate(numbers, ops) == testVal {
				res += testVal
				break
			}
		}
	}

	return res
}

func operatorCombos(n int, includeConcat bool) [][]string {
	if n == 0 {
		return [][]string{}
	}

	baseOps := []string{"+", "*"}
	if includeConcat {
		baseOps = append(baseOps, "||")
	}

	if n == 1 {
		var singleCombos [][]string
		for _, op := range baseOps {
			singleCombos = append(singleCombos, []string{op})
		}
		return singleCombos
	}

	subCombo := operatorCombos(n-1, includeConcat)
	var combos [][]string

	for _, sub := range subCombo {
		for _, op := range baseOps {
			combos = append(combos, append([]string{op}, sub...))
		}
	}

	return combos
}

func evaluate(numbers []int, operators []string) int {
	res := numbers[0]
	for i, op := range operators {
		switch op {
		case "+":
			res += numbers[i+1]
		case "*":
			res *= numbers[i+1]
		case "||":
			res = ints.FromString(fmt.Sprintf("%d%d", res, numbers[i+1]))
		}

	}

	return res
}

func part2(name string) int {
	input := files.ReadLines(name)

	return calcInput(input, true)
}
