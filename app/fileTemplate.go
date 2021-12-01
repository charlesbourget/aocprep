package app

func FileTemplate() string {
	return `package main

import (
	"fmt"
	"github.com/charlesbourget/aoc{{.Year}}/lib"
)

func main() {
	input, err := lib.Read("day{{.Day}}/input")
	if err != nil {
		fmt.Println("Error while reading input. ", err)
		return
	}

	fmt.Printf("Part 1: %s\n", part1(input))
	fmt.Printf("Part 2: %s\n", part2(input))
}

func part1(input []string) string {

	return ""
}

func part2(input []string) string {

	return ""
}
		`
}
