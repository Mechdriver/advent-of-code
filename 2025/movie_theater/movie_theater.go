package movie_theater

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Point struct {
	x int
	y int
}

func calcArea(pointA Point, pointB Point) int {
	xLen := pointA.x - pointB.x
	if xLen < 0 {
		xLen *= -1
	}
	xLen++

	yLen := pointA.y - pointB.y
	if yLen < 0 {
		yLen *= -1
	}
	yLen++

	area := xLen * yLen

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

func addToBounds(key int, val int, bounds map[int][]int) {
	_, okX := bounds[key]

	if okX {
		bounds[key] = append(bounds[key], val)
		sort.Ints(bounds[key])
	} else {
		bounds[key] = []int{val}
	}
}

func drawBoundaryLine(xToYBounds map[int][]int, yToXBounds map[int][]int, pointA Point, pointB Point) {
	ndx := pointA.x
	for ndx <= pointB.x {
		addToBounds(ndx, pointA.y, xToYBounds)
		addToBounds(pointA.y, ndx, yToXBounds)
		ndx++
	}

	ndx = pointA.x
	for ndx >= pointB.x {
		addToBounds(ndx, pointA.y, xToYBounds)
		addToBounds(pointA.y, ndx, yToXBounds)
		ndx--
	}

	ndx = pointA.y
	for ndx <= pointB.y {
		addToBounds(ndx, pointA.x, yToXBounds)
		addToBounds(pointA.x, ndx, xToYBounds)
		ndx++
	}

	ndx = pointA.y
	for ndx >= pointB.y {
		addToBounds(ndx, pointA.x, yToXBounds)
		addToBounds(pointA.x, ndx, xToYBounds)
		ndx--
	}
}

func isBetween(num int, min int, max int) bool {
	return num >= min && num <= max
}

func isLineWithinBounds(xToYBounds map[int][]int, yToXBounds map[int][]int, pointA Point, pointB Point) bool {
	ndx := pointA.x
	for ndx < pointB.x {
		xBounds := yToXBounds[pointA.y]
		yBounds := xToYBounds[ndx]

		if !isBetween(ndx, xBounds[0], xBounds[len(xBounds)-1]) {
			return false
		}

		if !isBetween(pointA.y, yBounds[0], yBounds[len(yBounds)-1]) {
			return false
		}
		ndx++
	}

	ndx = pointA.x
	for ndx > pointB.x {
		xBounds := yToXBounds[pointA.y]
		yBounds := xToYBounds[ndx]

		if !isBetween(ndx, xBounds[0], xBounds[len(xBounds)-1]) {
			return false
		}

		if !isBetween(pointA.y, yBounds[0], yBounds[len(yBounds)-1]) {
			return false
		}
		ndx--
	}

	ndx = pointA.y
	for ndx < pointB.y {
		xBounds := yToXBounds[ndx]
		yBounds := xToYBounds[pointA.x]

		if !isBetween(pointA.x, xBounds[0], xBounds[len(xBounds)-1]) {
			return false
		}

		if !isBetween(ndx, yBounds[0], yBounds[len(yBounds)-1]) {
			return false
		}
		ndx++
	}

	ndx = pointA.y
	for ndx > pointB.y {
		xBounds := yToXBounds[ndx]
		yBounds := xToYBounds[pointA.x]

		if !isBetween(pointA.x, xBounds[0], xBounds[len(xBounds)-1]) {
			return false
		}

		if !isBetween(ndx, yBounds[0], yBounds[len(yBounds)-1]) {
			return false
		}
		ndx--
	}

	return true
}

func getBoundaryMaps(points []Point) (map[int][]int, map[int][]int) {
	xToYBounds := make(map[int][]int)
	yToXBounds := make(map[int][]int)

	for _, point := range points {
		addToBounds(point.x, point.y, xToYBounds)
		addToBounds(point.y, point.x, yToXBounds)
	}

	return xToYBounds, yToXBounds
}

func createShapeBoundary(points []Point) (map[int][]int, map[int][]int) {
	xToYBounds, yToXBounds := getBoundaryMaps(points)

	for i := 1; i < len(points); i++ {
		pointA := points[i-1]
		pointB := points[i]

		drawBoundaryLine(xToYBounds, yToXBounds, pointA, pointB)
	}

	pointA := points[len(points)-1]
	pointB := points[0]

	drawBoundaryLine(xToYBounds, yToXBounds, pointA, pointB)

	return xToYBounds, yToXBounds
}

func findLargestGreenRedArea(points []Point) int {
	maxArea := 0
	xToYBounds, yToXBounds := createShapeBoundary(points)

	for i := range points {
		for j := i; j < len(points); j++ {
			c1 := points[i]
			c2 := points[j]
			c3 := Point{c2.x, c1.y}
			c4 := Point{c1.x, c2.y}

			rectangle := [4]Point{c1, c2, c3, c4}
			isValid := true

			for i := 1; i < len(rectangle); i++ {
				pointA := rectangle[i-1]
				pointB := rectangle[i]

				if !isLineWithinBounds(xToYBounds, yToXBounds, pointA, pointB) {
					isValid = false
					break
				}
			}

			pointA := rectangle[len(rectangle)-1]
			pointB := rectangle[0]

			if !isLineWithinBounds(xToYBounds, yToXBounds, pointA, pointB) {
				isValid = false
			}

			if isValid {
				area := calcArea(points[i], points[j])
				if area > maxArea {
					maxArea = area
					// fmt.Println("Max Valid: ", points[i], points[j])
				}
			} else {
			}
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

	for _, row := range floorGrid {
		fmt.Println(row)
	}

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

	// buildTiles(points)

	return findLargestGreenRedArea(points)
}
