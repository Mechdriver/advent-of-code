package gift_shop

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func findBadIds(start int, end int) int {
	summedIds := 0
	// fmt.Printf("\n%d - %d\n", start, end)

	for i := start; i <= end; i++ {
		id := strconv.Itoa(i)
		// fmt.Printf("i: %d\nID: %s\n", i, id)

		left := id[:(len(id) / 2)]
		right := id[(len(id) / 2):]

		// fmt.Printf("%s & %s\n", left, right)

		if left == right {
			fmt.Printf("Bad ID: %s\n", id)
			summedIds += i
			// fmt.Printf("count: %d\n", count)
		}
	}

	return summedIds
}

func findBadIdSequences(start int, end int) int {
	summedIds := 0

	for i := start; i <= end; i++ {
		id := strconv.Itoa(i)

		badId := false

		for j := 1; j <= len(id)/2; j++ {
			subStr := id[:j]
			subCount := strings.Count(id, subStr)

			if subCount*len(subStr) == len(id) {
				badId = true
			}
		}

		if badId {
			summedIds += i
		}
	}

	return summedIds
}

func Part1() int {
	summedIds := 0
	buffer, err := os.ReadFile("./inputs/day2/example_input.txt")

	if err != nil {
		panic(err)
	}

	inputs := strings.Split(string(buffer), ",")

	for i := range inputs {
		idRange := strings.Split(inputs[i], "-")
		start, err := strconv.Atoi(idRange[0])

		if err != nil {
			panic(err)
		}

		end, err := strconv.Atoi(strings.TrimSpace(idRange[1]))

		if err != nil {
			panic(err)
		}

		summedIds += findBadIds(start, end)
	}
	return summedIds
}

func Part2() int {
	summedIds := 0
	buffer, err := os.ReadFile("./inputs/day2/test_input.txt")

	if err != nil {
		panic(err)
	}

	inputs := strings.Split(string(buffer), ",")

	for i := range inputs {
		idRange := strings.Split(inputs[i], "-")
		start, err := strconv.Atoi(idRange[0])

		if err != nil {
			panic(err)
		}

		end, err := strconv.Atoi(strings.TrimSpace(idRange[1]))

		if err != nil {
			panic(err)
		}

		summedIds += findBadIdSequences(start, end)
	}
	return summedIds
}
