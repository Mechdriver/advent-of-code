package cafeteria

import (
	"os"
	"strconv"
	"strings"
)

type Range = struct {
	start int
	end   int
}

func getFreshCount(freshRanges []Range, ids []int) int {
	count := 0

	for _, id := range ids {
		for _, freshRange := range freshRanges {
			if id >= freshRange.start && id <= freshRange.end {
				count++
				break
			}
		}
	}

	return count
}

func Part1() int {
	var freshRanges []Range
	var ids []int

	buffer, err := os.ReadFile("./inputs/day5/test_input.txt")

	if err != nil {
		panic(err)
	}

	inputs := strings.Split(string(buffer), "\n")

	for i := range inputs {
		input := inputs[i]

		if input == "" {
			continue
		} else if strings.Contains(input, "-") {
			stringRange := strings.Split(input, "-")

			start, err := strconv.Atoi(stringRange[0])

			if err != nil {
				panic(err)
			}

			end, err := strconv.Atoi(stringRange[1])

			if err != nil {
				panic(err)
			}

			freshRanges = append(freshRanges, Range{start, end})
		} else {
			id, err := strconv.Atoi(input)

			if err != nil {
				panic(err)
			}

			ids = append(ids, id)
		}
	}

	return getFreshCount(freshRanges, ids)
}
