package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/lukeberry99/aoc-2024/pkg/files"
)

func main() {
	fmt.Println("Part 1:", part1("input.txt"))
	fmt.Println("Part 2:", part2("input.txt"))
}

func part1(name string) int {
	input := files.ReadParagraphs(name)

	rulesInput := input[0]
	updatesInput := input[1]

	rules := parseRules(rulesInput)
	updates := parseUpdates(updatesInput)

	sumMiddlePages := 0
	for _, update := range updates {
		if isCorrectlyOrdered(update, rules) {
			sumMiddlePages += findMiddlePage(update)
		}
	}

	return sumMiddlePages
}

func parseRules(rulesInput []string) map[int][]int {
	rules := make(map[int][]int)
	for _, rule := range rulesInput {
		parts := strings.Split(rule, "|")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		rules[x] = append(rules[x], y)
	}
	return rules
}

func parseUpdates(updatesInput []string) [][]int {
	var updates [][]int
	for _, update := range updatesInput {
		pageStrings := strings.Split(update, ",")
		var pages []int
		for _, page := range pageStrings {
			p, _ := strconv.Atoi(page)
			pages = append(pages, p)
		}
		updates = append(updates, pages)
	}
	return updates
}

func isCorrectlyOrdered(update []int, rules map[int][]int) bool {
	pageToIndex := make(map[int]int)
	for index, pageNumber := range update {
		pageToIndex[pageNumber] = index
	}

	for sourcePage, targetPages := range rules {
		if sourceIndex, sourceExists := pageToIndex[sourcePage]; sourceExists {
			for _, targetPage := range targetPages {
				if targetIndex, targetExists := pageToIndex[targetPage]; targetExists {
					if sourceIndex >= targetIndex {
						return false
					}
				}
			}
		}
	}

	return true
}

func findMiddlePage(update []int) int {
	midIdx := len(update) / 2
	return update[midIdx]
}

func part2(name string) int {
	return 0
}
