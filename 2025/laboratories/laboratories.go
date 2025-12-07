package laboratories

import (
	"fmt"
	"os"
	"strings"
)

func getSplits(manifoldGrid [][]string) int {
	splitCount := 0

	for i := 1; i < len(manifoldGrid); i++ {
		for j := 0; j < len(manifoldGrid[i]); j++ {
			if manifoldGrid[i][j] == "." {
				if manifoldGrid[i-1][j] == "S" || manifoldGrid[i-1][j] == "|" {
					manifoldGrid[i][j] = "|"
				}
			}

			if manifoldGrid[i][j] == "^" && manifoldGrid[i-1][j] == "|" {
				if j > 0 {
					manifoldGrid[i][j-1] = "|"
				}

				if j < len(manifoldGrid[i]) {
					manifoldGrid[i][j+1] = "|"
				}
				splitCount++
			}
		}

		for _, line := range manifoldGrid {
			fmt.Println(line)
		}
	}

	return splitCount
}

func copyGrid(manifoldGrid [][]string) [][]string {
	manifoldCopy := make([][]string, len(manifoldGrid))

	for i := range manifoldGrid {
		manifoldCopy[i] = make([]string, len(manifoldGrid[i]))
		copy(manifoldCopy[i], manifoldGrid[i])
	}

	return manifoldCopy
}

func getTimelines(manifoldGrid [][]string, i int) int {
	timeLines := 0

	if i == len(manifoldGrid) {
		return 1
	}

	for j := 0; j < len(manifoldGrid[i]); j++ {
		if manifoldGrid[i][j] == "." {
			if manifoldGrid[i-1][j] == "S" || manifoldGrid[i-1][j] == "|" {
				manifoldCopy := copyGrid(manifoldGrid)
				manifoldCopy[i][j] = "|"
				timeLines += getTimelines(manifoldCopy, i+1)
			}
		}

		if manifoldGrid[i][j] == "^" && manifoldGrid[i-1][j] == "|" {
			if j > 0 {
				manifoldCopy := copyGrid(manifoldGrid)
				manifoldCopy[i][j-1] = "|"
				timeLines += getTimelines(manifoldCopy, i+1)
			}

			if j < len(manifoldGrid[i]) {
				manifoldCopy := copyGrid(manifoldGrid)
				manifoldCopy[i][j+1] = "|"
				timeLines += getTimelines(manifoldCopy, i+1)
			}
		}
	}

	return timeLines
}

func Part1() int {
	buffer, err := os.ReadFile("./inputs/day7/example_input.txt")

	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(buffer), "\n")

	if lines[len(lines)-1] == "\n" {
		// Drop last empty line
		lines = lines[:len(lines)-1]
	}

	var manifoldGrid [][]string

	for i := 0; i < len(lines)-1; i++ {
		manifoldGrid = append(manifoldGrid, strings.Split(lines[i], ""))
	}

	return getSplits(manifoldGrid)
}

func Part2() int {
	buffer, err := os.ReadFile("./inputs/day7/example_input.txt")

	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(buffer), "\n")

	if lines[len(lines)-1] == "\n" {
		// Drop last empty line
		lines = lines[:len(lines)-1]
	}

	var manifoldGrid [][]string

	for i := 0; i < len(lines)-1; i++ {
		manifoldGrid = append(manifoldGrid, strings.Split(lines[i], ""))
	}

	return getTimelines(manifoldGrid, 1)
}
