package trash_compactor

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func doHomework(homeWork [][]int, symbols []string) int {
	homeWorkAnswer := 0

	for j := range homeWork[0] {
		symbol := symbols[j]
		val := 0

		if symbol == "*" {
			val = 1
		}

		for i := range homeWork {
			switch symbol {
			case "*":
				val *= homeWork[i][j]
			case "+":
				val += homeWork[i][j]
			}

			if i > 0 {
				fmt.Print(" ", symbol, " ")
			}

			fmt.Print(homeWork[i][j])
		}
		fmt.Print(" = ", val)
		fmt.Println("\n=========")
		homeWorkAnswer += val
	}

	return homeWorkAnswer
}

func doCephalopodHomework(homework [][]int, symbols []string) int {
	homeWorkAnswer := 0

	for i, problem := range homework {
		symbol := symbols[i]
		val := 0

		if symbol == "*" {
			val = 1
		}

		for _, num := range problem {
			switch symbol {
			case "*":
				val *= num
			case "+":
				val += num
			}
		}
		homeWorkAnswer += val
	}

	return homeWorkAnswer
}

func Part1() int {
	buffer, err := os.ReadFile("./inputs/day6/test_input.txt")

	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(buffer), "\n")
	// Drop last empty line
	lines = lines[:len(lines)-1]

	var homeWork [][]int
	var symbols []string

	allWhiteSpace := "\\s+"
	regex := regexp.MustCompile(allWhiteSpace)

	for i := 0; i < len(lines)-1; i++ {
		parsedLine := regex.Split(strings.TrimSpace(lines[i]), -1)
		fmt.Println("Parsed: ", parsedLine)
		var homeWorkLine []int

		for _, strVal := range parsedLine {
			println("Val: ", strVal)
			num, err := strconv.Atoi(strVal)

			if err != nil {
				panic(err)
			}

			homeWorkLine = append(homeWorkLine, num)
		}
		homeWork = append(homeWork, homeWorkLine)
	}

	symbols = regex.Split(strings.TrimSpace(lines[len(lines)-1]), -1)

	return doHomework(homeWork, symbols)
}

func Part2() int {
	buffer, err := os.ReadFile("./inputs/day6/test_input.txt")

	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(buffer), "\n")
	// Drop last empty line
	lines = lines[:len(lines)-1]

	var homeWorkGrid [][]string
	var homeWork [][]int

	allWhiteSpace := "\\s+"
	regex := regexp.MustCompile(allWhiteSpace)
	symbols := regex.Split(strings.TrimSpace(lines[len(lines)-1]), -1)

	for i := 0; i < len(lines)-1; i++ {
		homeWorkGrid = append(homeWorkGrid, strings.Split(lines[i], ""))
	}

	var numberLine = make([]int, 0)

	for j := range homeWorkGrid[0] {
		valString := ""

		for i := range homeWorkGrid {
			if homeWorkGrid[i][j] != " " {
				valString += homeWorkGrid[i][j]
			}
		}

		if valString != "" {
			val, err := strconv.Atoi(valString)

			if err != nil {
				panic(err)
			}

			numberLine = append(numberLine, val)
		} else {
			homeWork = append(homeWork, numberLine)
			numberLine = make([]int, 0)
		}
	}

	homeWork = append(homeWork, numberLine)

	return doCephalopodHomework(homeWork, symbols)
}
