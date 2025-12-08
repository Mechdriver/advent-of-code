package laboratories

import (
	"fmt"
	"os"
	"strings"
)

type Point struct {
	display string
	paths   int
}

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

func printManifoldPoints(manifoldPoints [][]Point) {
	for _, points := range manifoldPoints {
		var displayList []string
		var pathsList []int
		for _, point := range points {
			displayList = append(displayList, point.display)
			pathsList = append(pathsList, point.paths)
		}
		fmt.Println(displayList, pathsList)
	}
	fmt.Println()
}

func buildSplits(manifoldPoints [][]Point, verbose bool) [][]Point {
	for i := 1; i < len(manifoldPoints); i++ {
		for j := 0; j < len(manifoldPoints[i]); j++ {

			if manifoldPoints[i][j].display == "." {
				if manifoldPoints[i-1][j].display == "S" || manifoldPoints[i-1][j].display == "|" {
					manifoldPoints[i][j].display = "|"
				}
			}

			if manifoldPoints[i][j].display == "^" && manifoldPoints[i-1][j].display == "|" {
				if j > 0 {
					manifoldPoints[i][j-1].display = "|"
				}

				if j < len(manifoldPoints[i]) {
					manifoldPoints[i][j+1].display = "|"
				}
			}
		}
		if verbose {
			printManifoldPoints(manifoldPoints)
		}
	}

	return manifoldPoints
}

func buildPaths(manifoldPoints [][]Point, verbose bool) [][]Point {
	for i := 1; i < len(manifoldPoints); i++ {
		for j := 0; j < len(manifoldPoints[i]); j++ {

			if manifoldPoints[i][j].display == "|" {
				switch manifoldPoints[i-1][j].display {
				case "|":
					manifoldPoints[i][j].paths = manifoldPoints[i-1][j].paths
				case "S":
					manifoldPoints[i][j].paths = 1
				}

				if j < len(manifoldPoints[i])-1 && manifoldPoints[i][j+1].display == "^" {
					manifoldPoints[i][j].paths += manifoldPoints[i-1][j+1].paths
				}

				if j > 0 && manifoldPoints[i][j-1].display == "^" {
					manifoldPoints[i][j].paths += manifoldPoints[i-1][j-1].paths
				}
			}
		}
		if verbose {
			printManifoldPoints(manifoldPoints)
		}
	}
	return manifoldPoints
}

func getTimelinesSmart(manifoldPoints [][]Point) int {
	timelines := 0

	builtManifold := buildPaths(buildSplits(manifoldPoints, false), false)

	for _, point := range builtManifold[len(builtManifold)-1] {
		timelines += point.paths
	}

	return timelines
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
	buffer, err := os.ReadFile("./inputs/day7/test_input.txt")

	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(buffer), "\n")

	if lines[len(lines)-1] == "\n" {
		// Drop last empty line
		lines = lines[:len(lines)-1]
	}

	var manifoldPoints [][]Point

	for i := 0; i < len(lines)-1; i++ {
		var points []Point
		for _, item := range strings.Split(lines[i], "") {
			points = append(points, Point{item, 0})
		}
		manifoldPoints = append(manifoldPoints, points)
	}

	return getTimelinesSmart(manifoldPoints)
}
