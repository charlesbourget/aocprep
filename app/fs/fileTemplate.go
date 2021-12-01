package fs

func SourceFileTemplate() string {
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

func Part1(input []string) string {

	return ""
}

func Part2(input []string) string {

	return ""
}

`
}

func TestFileTemplate() string {
	return `package main

import (
	"fmt"
	"github.com/charlesbourget/aoc{{.Year}}/lib"
	"testing"
)

func TestPart1(t *testing.T) {
	input, err := lib.Read("input.test")
	if err != nil {
		fmt.Println("Error while reading input. ", err)
		return
	}

	expected := ""
	result := Part1(input)
	if result != expected {
		t.Fatalf('Part1() = %s, want %s, error', result, expected)
	}
}

func TestPart2(t *testing.T) {
	input, err := lib.Read("input.test")
	if err != nil {
		fmt.Println("Error while reading input. ", err)
		return
	}

	expected := ""
	result := Part2(input)
	if result != expected {
		t.Fatalf('Part2() = %s, want %s, error', result, expected)
	}
}

`

}
