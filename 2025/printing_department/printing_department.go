package printing_department

import (
	"fmt"
	"os"
	"strings"
)

type Point struct {
	X int
	Y int
}

func isAccessible(grid [][]string, loc Point) bool {
	rollCount := 0
	lenY := len(grid) - 1
	lenX := len(grid[0]) - 1

	// Up
	if loc.Y > 0 {
		if grid[loc.Y-1][loc.X] == "@" {
			rollCount++
		}
	}
	// Up-Right
	if loc.Y > 0 && loc.X < lenX {
		if grid[loc.Y-1][loc.X+1] == "@" {
			rollCount++
		}
	}
	// Right
	if loc.X < lenX {
		if grid[loc.Y][loc.X+1] == "@" {
			rollCount++
		}
	}
	// Down-Right
	if loc.Y < lenY && loc.X < lenX {
		if grid[loc.Y+1][loc.X+1] == "@" {
			rollCount++
		}
	}
	// Down
	if loc.Y < lenY {
		if grid[loc.Y+1][loc.X] == "@" {
			rollCount++
		}
	}
	// Down-Left
	if loc.Y < lenY && loc.X > 0 {
		if grid[loc.Y+1][loc.X-1] == "@" {
			rollCount++
		}
	}
	// Left
	if loc.X > 0 {
		if grid[loc.Y][loc.X-1] == "@" {
			rollCount++
		}
	}
	// Up-Left
	if loc.Y > 0 && loc.X > 0 {
		if grid[loc.Y-1][loc.X-1] == "@" {
			rollCount++
		}
	}

	return rollCount < 4
}

func findAccessibleRolls(grid [][]string) int {
	count := 0

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == "@" {
				if isAccessible(grid, Point{j, i}) {
					count++
					fmt.Print("x")
				} else {
					fmt.Print("@")
				}
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}

	return count
}

func Part1() int {
	var floorGrid [][]string
	buffer, err := os.ReadFile("./inputs/day4/test_input.txt")

	if err != nil {
		panic(err)
	}

	inputs := strings.Split(string(buffer), "\n")

	for i := range inputs {
		if len(inputs[i]) > 0 {
			floorGrid = append(floorGrid, strings.Split(inputs[i], ""))
		}
	}

	return findAccessibleRolls(floorGrid)
}
