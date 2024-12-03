package main

import (
	"fmt"

	"github.com/lukeberry99/aoc-2024/pkg/files"
	"github.com/lukeberry99/aoc-2024/pkg/ints"
	"github.com/lukeberry99/aoc-2024/pkg/slices"
)

func main() {
	fmt.Println("Part 1:", part1("input.txt"))
	fmt.Println("Part 2:", part2("input.txt"))
}

func part1(name string) int {
	columns := files.ReadColumns(name, 2)

	columnsA := slices.SortSlice(ints.FromStringSlice(columns[0]))
	columnsB := slices.SortSlice(ints.FromStringSlice(columns[1]))

	sum := 0
	for i := 0; i < len(columnsA); i++ {
		if columnsA[i] > columnsB[i] {
			sum += columnsA[i] - columnsB[i]
		} else {
			sum += columnsB[i] - columnsA[i]
		}
	}

	return sum
}

func part2(name string) int {
	columns := files.ReadColumns(name, 2)

	columnsA := slices.SortSlice(ints.FromStringSlice(columns[0]))
	columnsB := ints.FromStringSlice(columns[1])

	freqMap := make(map[int]int)
	for _, num := range columnsB {
		freqMap[num]++
	}

	similarity := 0

	for _, num := range columnsA {
		if count, exists := freqMap[num]; exists {
			similarity += num * count
		}
	}

	return similarity
}
