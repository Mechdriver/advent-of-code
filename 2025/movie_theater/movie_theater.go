package movie_theater

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x int
	y int
}

func calcArea(pointA Point, pointB Point) int {
	xLen := pointA.x - pointB.x + 1
	yLen := pointA.y - pointB.y + 1

	area := xLen * yLen

	if area < 0 {
		return area * -1
	}

	return area
}

func findLargestArea(points []Point) int {
	maxArea := 0

	for i := range points {
		for j := i; j < len(points); j++ {
			area := calcArea(points[i], points[j])
			if area > maxArea {
				maxArea = area
			}
		}
	}

	return maxArea
}

func drawLine(floorGrid [][]string, pointA Point, pointB Point) {
	ndx := pointA.x
	for ndx < pointB.x {
		floorGrid[pointA.y][ndx] = "X"
		ndx++
	}

	ndx = pointA.x
	for ndx > pointB.x {
		floorGrid[pointA.y][ndx] = "X"
		ndx--
	}

	ndx = pointA.y
	for ndx < pointB.y {
		floorGrid[ndx][pointA.x] = "X"
		ndx++
	}

	ndx = pointA.y
	for ndx > pointB.y {
		floorGrid[ndx][pointA.x] = "X"
		ndx--
	}
}

func isWithinShape(floorGrid [][]string, point Point) bool {
	up, down, left, right := false, false, false, false

	// Up
	for i := point.y; i > 0; i-- {
		if floorGrid[i][point.x] == "X" || floorGrid[i][point.x] == "#" {
			up = true
			break
		}
	}

	// Down
	for i := point.y; i < len(floorGrid); i++ {
		if floorGrid[i][point.x] == "X" || floorGrid[i][point.x] == "#" {
			down = true
			break
		}
	}

	// Left
	for j := point.x; j > 0; j-- {
		if floorGrid[point.y][j] == "X" || floorGrid[point.y][j] == "#" {
			left = true
			break
		}
	}

	// Right
	for j := point.x; j < len(floorGrid[0]); j++ {
		if floorGrid[point.y][j] == "X" || floorGrid[point.y][j] == "#" {
			right = true
		}
	}

	return up && down && left && right
}

func findLargestGreenRedArea(floorGrid [][]string, points []Point) int {
	maxArea := 0

	fmt.Println("Finding area...")
	for i := range points {
		for j := i; j < len(points); j++ {
			c1 := points[i]
			c2 := points[j]
			c3 := Point{c2.x, c1.y}
			c4 := Point{c1.x, c2.y}

			rectangle := [4]Point{c1, c2, c3, c4}
			isValid := true

			for _, point := range rectangle {
				if !isWithinShape(floorGrid, point) {
					isValid = false
				}
			}

			if isValid {
				area := calcArea(points[i], points[j])
				if area > maxArea {
					maxArea = area
				}
			}
		}
		if i%10 == 0 {
			fmt.Println("Finished ", i, " out of 496")
		}
	}

	return maxArea
}

func getMaxXandY(points []Point) (int, int) {
	maxX := 0
	maxY := 0

	for _, point := range points {
		if point.x > maxX {
			maxX = point.x
		}

		if point.y > maxY {
			maxY = point.y
		}
	}

	return maxX, maxY
}

func buildTiles(points []Point) [][]string {
	floorGrid := make([][]string, 0)
	maxX, maxY := getMaxXandY(points)

	for range maxY + 1 {
		var floorRow []string

		for range maxX + 1 {
			floorRow = append(floorRow, ".")
		}
		floorGrid = append(floorGrid, floorRow)
	}

	for i := 1; i < len(points); i++ {
		pointA := points[i-1]
		pointB := points[i]

		drawLine(floorGrid, pointA, pointB)

		floorGrid[pointA.y][pointA.x] = "#"
	}

	pointA := points[len(points)-1]
	pointB := points[0]

	drawLine(floorGrid, pointA, pointB)
	floorGrid[pointA.y][pointA.x] = "#"

	// for _, row := range floorGrid {
	// 	fmt.Println(row)
	// }

	return floorGrid
}

func parseInput() []Point {
	buffer, err := os.ReadFile("./inputs/day9/test_input.txt")

	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(buffer), "\n")

	if lines[len(lines)-1] == "" {
		// Drop last empty line
		lines = lines[:len(lines)-1]
	}

	var points []Point

	for _, line := range lines {
		coords := strings.Split(line, ",")
		x, err := strconv.Atoi(coords[0])

		if err != nil {
			panic(err)
		}

		y, err := strconv.Atoi(coords[1])

		if err != nil {
			panic(err)
		}

		point := Point{x, y}

		points = append(points, point)
	}

	return points
}

func Part1() int {
	points := parseInput()

	return findLargestArea(points)
}

func Part2() int {
	points := parseInput()

	fmt.Println("Building tiles...")
	floorGrid := buildTiles(points)
	fmt.Println("Tiles built!")

	// maxX, maxY := getMaxXandY(points)

	// fmt.Println("Max X: ", maxX)
	// fmt.Println("Max y: ", maxY)

	return findLargestGreenRedArea(floorGrid, points)
}
