package movie_theater

import (
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x int
	y int
}

type Boundary struct {
	min int
	max int
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

func addToBounds(key int, val int, bounds map[int]Boundary) {
	_, okX := bounds[key]

	if okX {
		boundary := bounds[key]
		newBoundary := Boundary{min(boundary.min, val), max(boundary.max, val)}
		bounds[key] = newBoundary
	} else {
		bounds[key] = Boundary{val, val}
	}
}

func drawBoundaryLine(xToYBounds map[int]Boundary, yToXBounds map[int]Boundary, pointA Point, pointB Point) {
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

func isPointWithinBounds(xToYBounds map[int]Boundary, yToXBounds map[int]Boundary, point Point) bool {
	xBounds := yToXBounds[point.y]
	yBounds := xToYBounds[point.x]

	return isBetween(point.x, xBounds.min, xBounds.max) && isBetween(point.y, yBounds.min, yBounds.max)
}

func isLineWithinBounds(xToYBounds map[int]Boundary, yToXBounds map[int]Boundary, pointA Point, pointB Point) bool {
	ndx := pointA.x
	for ndx < pointB.x {
		xBounds := yToXBounds[pointA.y]
		yBounds := xToYBounds[ndx]

		if !isBetween(ndx, xBounds.min, xBounds.max) {
			return false
		}

		if !isBetween(pointA.y, yBounds.min, yBounds.max) {
			return false
		}
		ndx++
	}

	ndx = pointA.x
	for ndx > pointB.x {
		xBounds := yToXBounds[pointA.y]
		yBounds := xToYBounds[ndx]

		if !isBetween(ndx, xBounds.min, xBounds.max) {
			return false
		}

		if !isBetween(pointA.y, yBounds.min, yBounds.max) {
			return false
		}
		ndx--
	}

	ndx = pointA.y
	for ndx < pointB.y {
		xBounds := yToXBounds[ndx]
		yBounds := xToYBounds[pointA.x]

		if !isBetween(pointA.x, xBounds.min, xBounds.max) {
			return false
		}

		if !isBetween(ndx, yBounds.min, yBounds.max) {
			return false
		}
		ndx++
	}

	ndx = pointA.y
	for ndx > pointB.y {
		xBounds := yToXBounds[ndx]
		yBounds := xToYBounds[pointA.x]

		if !isBetween(pointA.x, xBounds.min, xBounds.max) {
			return false
		}

		if !isBetween(ndx, yBounds.min, yBounds.max) {
			return false
		}
		ndx--
	}

	return true
}

func createShapeBoundary(points []Point) (map[int]Boundary, map[int]Boundary) {
	xToYBounds := make(map[int]Boundary)
	yToXBounds := make(map[int]Boundary)

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

			if isPointWithinBounds(xToYBounds, yToXBounds, c3) && isPointWithinBounds(xToYBounds, yToXBounds, c4) {
				rectangle := [5]Point{c1, c2, c3, c4, c1}
				isValid := true

				for i := 1; i < len(rectangle); i++ {
					pointA := rectangle[i-1]
					pointB := rectangle[i]

					if !isLineWithinBounds(xToYBounds, yToXBounds, pointA, pointB) {
						isValid = false
						break
					}
				}

				if isValid {
					area := calcArea(points[i], points[j])
					if area > maxArea {
						maxArea = area
					}
				}
			}
		}
	}

	return maxArea
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

	return findLargestGreenRedArea(points)
}
