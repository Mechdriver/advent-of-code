package playground

import (
	"fmt"
	"maps"
	"math"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

type Box struct {
	x int
	y int
	z int
}

func getDistance(pointA Box, pointB Box) float64 {
	return math.Sqrt(math.Pow(float64(pointA.x)-float64(pointB.x), 2) + math.Pow(float64(pointA.y)-float64(pointB.y), 2) + math.Pow(float64(pointA.z)-float64(pointB.z), 2))
}

func buildCircuits(boxes []Box, conLimit int) []map[Box]bool {
	connections := 0
	circuits := make([]map[Box]bool, 0)
	distanceMap := make(map[float64][2]Box)
	var distances []float64

	for i := range boxes {
		for j := i + 1; j < len(boxes); j++ {
			distance := getDistance(boxes[i], boxes[j])

			distances = append(distances, distance)
			distanceMap[distance] = [2]Box{boxes[i], boxes[j]}
		}
	}

	sort.Float64s(distances)

	for connections < conLimit {
		circuitMade := false
		connectionNdx := -1
		pruneNdxes := make(map[int]bool)
		boxes := distanceMap[distances[connections]]
		boxA := boxes[0]
		boxB := boxes[1]

		for i, circuit := range circuits {
			if !circuitMade {
				if circuit[boxA] {
					circuit[boxB] = true
					circuitMade = true
					connectionNdx = i
				} else if circuit[boxB] {
					circuit[boxA] = true
					circuitMade = true
					connectionNdx = i
				}
			} else if circuit[boxA] || circuit[boxB] {
				maps.Copy(circuits[connectionNdx], circuits[i])
				pruneNdxes[i] = true
			}
		}

		if len(circuits) == 0 || !circuitMade {
			circuits = append(circuits, map[Box]bool{boxA: true, boxB: true})
		}

		updatedCircuits := make([]map[Box]bool, 0)
		for i := range circuits {
			if !pruneNdxes[i] {
				updatedCircuits = append(updatedCircuits, circuits[i])
			}
		}

		circuits = updatedCircuits

		connections++

		// for _, circuit := range circuits {
		// 	fmt.Println(circuit)
		// }
		// fmt.Println("===============================================")
	}
	// fmt.Println(len(circuits))

	return circuits
}

func buildFullCircuit(boxes []Box) (Box, Box) {
	circuits := make([]map[Box]bool, 0)
	distanceMap := make(map[float64][2]Box)
	var distances []float64

	for i := range boxes {
		for j := i + 1; j < len(boxes); j++ {
			distance := getDistance(boxes[i], boxes[j])

			distances = append(distances, distance)
			distanceMap[distance] = [2]Box{boxes[i], boxes[j]}
		}
		circuits = append(circuits, map[Box]bool{boxes[i]: true})
	}

	sort.Float64s(distances)

	ndx := 0

	for len(circuits) > 1 {
		circuitMade := false
		connectionNdx := -1
		pruneNdxes := make(map[int]bool)
		boxes := distanceMap[distances[ndx]]
		boxA := boxes[0]
		boxB := boxes[1]

		for i, circuit := range circuits {
			if !circuitMade {
				if circuit[boxA] {
					circuit[boxB] = true
					circuitMade = true
					connectionNdx = i
				} else if circuit[boxB] {
					circuit[boxA] = true
					circuitMade = true
					connectionNdx = i
				}
			} else if circuit[boxA] || circuit[boxB] {
				maps.Copy(circuits[connectionNdx], circuits[i])
				pruneNdxes[i] = true
			}
		}

		if len(circuits) == 0 || !circuitMade {
			circuits = append(circuits, map[Box]bool{boxA: true, boxB: true})
		}

		updatedCircuits := make([]map[Box]bool, 0)
		for i := range circuits {
			if !pruneNdxes[i] {
				updatedCircuits = append(updatedCircuits, circuits[i])
			}
		}

		circuits = updatedCircuits

		ndx++
	}

	// for _, circuit := range circuits {
	// 	fmt.Println(circuit)
	// }
	// fmt.Println(len(circuits))

	lastBoxes := distanceMap[distances[ndx-1]]

	return lastBoxes[0], lastBoxes[1]
}

func Part2() int {
	buffer, err := os.ReadFile("./inputs/day8/test_input.txt")

	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(buffer), "\n")

	if lines[len(lines)-1] == "" {
		// Drop last empty line
		lines = lines[:len(lines)-1]
	}

	var boxes []Box

	for _, line := range lines {
		coords := strings.Split(line, ",")
		x, err := strconv.Atoi(coords[0])

		if err != nil {
			fmt.Println(coords)
			panic(err)
		}

		y, err := strconv.Atoi(coords[1])

		if err != nil {
			panic(err)
		}

		z, err := strconv.Atoi(coords[2])

		if err != nil {
			panic(err)
		}

		point := Box{x, y, z}

		boxes = append(boxes, point)
	}

	boxA, boxB := buildFullCircuit(boxes)

	// fmt.Println("A: ", boxA, " B: ", boxB)

	return boxA.x * boxB.x
}

func Part1() int {
	product := 1
	buffer, err := os.ReadFile("./inputs/day8/test_input.txt")

	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(buffer), "\n")

	if lines[len(lines)-1] == "" {
		// Drop last empty line
		lines = lines[:len(lines)-1]
	}

	var boxes []Box

	for _, line := range lines {
		coords := strings.Split(line, ",")
		x, err := strconv.Atoi(coords[0])

		if err != nil {
			fmt.Println(coords)
			panic(err)
		}

		y, err := strconv.Atoi(coords[1])

		if err != nil {
			panic(err)
		}

		z, err := strconv.Atoi(coords[2])

		if err != nil {
			panic(err)
		}

		point := Box{x, y, z}

		boxes = append(boxes, point)
	}

	circuits := buildCircuits(boxes, 1000)
	circuitLens := make([]int, 0)

	for _, circuit := range circuits {
		circuitLens = append(circuitLens, len(circuit))
	}

	sort.Ints(circuitLens)
	slices.Reverse(circuitLens)

	// fmt.Println(circuitLens)

	for i := 0; i < 3; i++ {
		product *= circuitLens[i]
	}

	return product
}
