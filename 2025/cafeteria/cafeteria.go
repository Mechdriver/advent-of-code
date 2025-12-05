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

func getAllFreshCount(freshRanges []Range) int {
	// Have list of ranges to check
	// Have list of checked ranges

	// while list is not empty
	// if range's minimum is within checked ranges
	// bump the min to the checked max + 1
	// else if it is less than min
	// add new range to check {min: currMin, max: checkedMin - 1}
	// currMin = checkedMax + 1
	// else if it is greater than max
	// do nothing

	// if range's max is within checked ranges
	// bump the max to the checked min - 1
	// else if it is greater than max
	// add new range to check {min: checkedMax + 1, max: currentMax}
	// currMax = checkedMin - 1
	// else if it is less than min
	// do nothing

	// When the range has been adjusted, do max - min + 1 and add to count if sum > 0

	count := 0
	skip := false
	unchecked := freshRanges
	checkedRanges := make([]Range, 0)

	for len(unchecked) > 0 {
		currRange := unchecked[0]
		unchecked = unchecked[1:]

		for _, checkedRange := range checkedRanges {
			// Check min
			if currRange.start >= checkedRange.start && currRange.start <= checkedRange.end {
				currRange.start = checkedRange.end + 1
			} else if currRange.start < checkedRange.start && currRange.end > checkedRange.start {
				unchecked = append(unchecked, Range{currRange.start, checkedRange.start - 1})
				currRange.start = checkedRange.start
				skip = true
			}

			// Check max
			if currRange.end >= checkedRange.start && currRange.end <= checkedRange.end {
				currRange.end = checkedRange.start - 1
			} else if currRange.end > checkedRange.end && currRange.start < checkedRange.end {
				unchecked = append(unchecked, Range{checkedRange.end + 1, currRange.end})
				currRange.end = checkedRange.end
				skip = true
			}
		}

		idAmount := currRange.end - currRange.start + 1
		if idAmount > 0 && !skip {
			count += idAmount
			checkedRanges = append(checkedRanges, currRange)
		}
		skip = false
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

func Part2() int {
	var freshRanges []Range

	buffer, err := os.ReadFile("./inputs/day5/test_input.txt")

	if err != nil {
		panic(err)
	}

	inputs := strings.Split(string(buffer), "\n")

	for i := range inputs {
		input := inputs[i]

		if input == "" {
			break
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
		}
	}

	return getAllFreshCount(freshRanges)
}
