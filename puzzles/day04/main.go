package main

import (
	"fmt"

	"github.com/lukeberry99/aoc-2024/pkg/files"
)

func main() {
	fmt.Println("Part 1:", part1("input.txt"))
	fmt.Println("Part 2:", part2("input.txt"))
}

func part1(name string) int {
	input := files.ReadLines(name)
	count := 0

	directions := [][2]int{
		{-1, -1}, // up-left
		{-1, 0},  // up
		{-1, 1},  // up-right
		{0, -1},  // left
		{0, 1},   // right
		{1, -1},  // down-left
		{1, 0},   // down
		{1, 1},   // down-right
	}

	target := "XMAS"
	rows := len(input)
	cols := len(input[0])

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if input[row][col] == 'X' {
				for _, dir := range directions {
					found := true
					for i := 0; i < len(target); i++ {
						newRow := row + i*dir[0]
						newCol := col + i*dir[1]

						if newRow < 0 || newRow >= rows || newCol < 0 || newCol >= cols {
							found = false
							break
						}

						if input[newRow][newCol] != target[i] {
							found = false
							break
						}
					}
					if found {
						count++
					}
				}
			}
		}
	}

	return count
}

func part2(name string) int {
	input := files.ReadLines(name)
	count := 0

	rows := len(input)
	cols := len(input[0])

	for row := 1; row < rows-1; row++ {
		for col := 1; col < cols-1; col++ {
			if input[row][col] != 'A' {
				continue
			}

			topLeft := input[row-1][col-1]
			topRight := input[row-1][col+1]
			bottomLeft := input[row+1][col-1]
			bottomRight := input[row+1][col+1]

			diag1MAS := (topLeft == 'M' && bottomRight == 'S')
			diag1SAM := (topLeft == 'S' && bottomRight == 'M')

			diag2MAS := (topRight == 'M' && bottomLeft == 'S')
			diag2SAM := (topRight == 'S' && bottomLeft == 'M')

			if (diag1MAS || diag1SAM) && (diag2MAS || diag2SAM) {
				count++
			}
		}
	}

	return count
}
