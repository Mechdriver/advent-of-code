package lobby

import (
	"os"
	"strconv"
	"strings"
)

func findLargeJoltage(bank string) int {
	var battery string
	max, ndx := 0, 0

	for i := 0; i < len(bank)-1; i++ {
		num, err := strconv.Atoi(string(bank[i]))

		if err != nil {
			panic(err)
		}

		if num > max {
			max = num
			ndx = i + 1
		}
	}

	battery += strconv.Itoa(max)
	max = 0

	for i := ndx; i < len(bank); i++ {
		num, err := strconv.Atoi(string(bank[i]))

		if err != nil {
			panic(err)
		}

		if num > max {
			max = num
		}
	}

	battery += strconv.Itoa(max)

	joltage, err := strconv.Atoi(battery)

	if err != nil {
		panic(err)
	}

	return joltage
}

func findNSizedJoltage(bank string, size int) int {
	battery := ""
	max, ndx, nextNdx := 0, 0, 0
	window := len(bank) - size + 1

	for len(battery) < size {
		for i := ndx; i < ndx+window; i++ {
			num, err := strconv.Atoi(string(bank[i]))

			if err != nil {
				panic(err)
			}

			if num > max {
				max = num
				nextNdx = i + 1
			}
		}
		battery += strconv.Itoa(max)
		ndx = nextNdx
		max = 0
		nextNdx = 0
		window = len(bank) - ndx - (size - len(battery)) + 1
	}

	joltage, err := strconv.Atoi(battery)

	if err != nil {
		panic(err)
	}

	return joltage
}

func Part1() int {
	totalJoltage := 0

	buffer, err := os.ReadFile("./inputs/day3/test_input.txt")

	if err != nil {
		panic(err)
	}

	inputs := strings.Split(string(buffer), "\n")

	for i := range inputs {
		totalJoltage += findLargeJoltage(inputs[i])
	}

	return totalJoltage
}

func Part2() int {
	totalJoltage := 0

	buffer, err := os.ReadFile("./inputs/day3/test_input.txt")

	if err != nil {
		panic(err)
	}

	inputs := strings.Split(string(buffer), "\n")

	for i := range inputs {
		totalJoltage += findNSizedJoltage(inputs[i], 12)
	}

	return totalJoltage
}
