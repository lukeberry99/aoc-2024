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
	data := files.ReadLines(name)

	safeCount := 0
	for i := 0; i < len(data); i++ {
		if isSafe(data[i]) {
			safeCount++
		}
	}

	return safeCount
}

func part2(name string) int {
	data := files.ReadLines(name)

	safeCount := 0
	for i := 0; i < len(data); i++ {
		if isSafeProblemDampener(data[i]) {
			safeCount++
		}
	}

	return safeCount
}

func isSafe(inp string) bool {
	spl := ints.FromStringSlice(strings.Split(inp, " "))

	if checkSequence(spl) {
		return true
	}

	return false
}

func isSafeProblemDampener(inp string) bool {
	spl := ints.FromStringSlice(strings.Split(inp, " "))

	if checkSequence(spl) {
		return true
	}

	for i := 0; i < len(spl); i++ {
		modified := make([]int, 0, len(spl)-1)
		modified = append(modified, spl[:i]...)
		modified = append(modified, spl[i+1:]...)

		if checkSequence(modified) {
			return true
		}
	}

	return false
}

func checkSequence(nums []int) bool {
	if len(nums) < 2 {
		return true
	}

	prev := nums[0]
	increasing := nums[1] > prev

	for i := 1; i < len(nums); i++ {
		curr := nums[i]
		diff := curr - prev

		if increasing && diff <= 0 {
			return false
		}
		if !increasing && diff >= 0 {
			return false
		}

		absDiff := diff
		if absDiff < 0 {
			absDiff = -absDiff
		}
		if absDiff < 1 || absDiff > 3 {
			return false
		}

		prev = curr
	}

	return true
}
