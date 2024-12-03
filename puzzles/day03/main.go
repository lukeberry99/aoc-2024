package main

import (
	"fmt"
	"regexp"

	"github.com/lukeberry99/aoc-2024/pkg/files"
	"github.com/lukeberry99/aoc-2024/pkg/ints"
)

func main() {
	fmt.Println("Part 1:", part1("input.txt"))
	fmt.Println("Part 2:", part2("input.txt"))
}

func part1(name string) int {
	input := files.Read(name)

	pattern := `mul\((\d+),(\d+)\)`

	re, err := regexp.Compile(pattern)
	if err != nil {
		panic("error compiling regex")
	}

	sanitised := re.FindAllStringSubmatch(input, -1)

	sum := 0

	for _, match := range sanitised {
		if len(match) == 3 {
			num1 := ints.FromString(match[1])
			num2 := ints.FromString(match[2])

			product := num1 * num2

			sum = sum + product
		}
	}

	return sum
}

func part2(name string) int {
	input := files.Read(name)

	pattern := `mul\((\d+),(\d+)\)|don't\(\)|do\(\)`

	re, err := regexp.Compile(pattern)
	if err != nil {
		panic("error compiling regex")
	}

	sanitised := re.FindAllStringSubmatch(input, -1)

	sum := 0

	canMultiply := true
	for _, match := range sanitised {
		if len(match) > 0 {
			fullMatch := match[0]

			if fullMatch == "don't()" {
				canMultiply = false
				continue
			}

			if fullMatch == "do()" {
				canMultiply = true
				continue
			}

			if canMultiply && len(match) == 3 {
				num1 := ints.FromString(match[1])
				num2 := ints.FromString(match[2])

				product := num1 * num2

				sum = sum + product
			}
		}
	}

	return sum
}
