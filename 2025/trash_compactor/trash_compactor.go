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
		val := 0
		if symbols[j] == "*" {
			val = 1
		}
		for i := range homeWork {
			switch symbols[j] {
			case "*":
				val *= homeWork[i][j]
			case "+":
				val += homeWork[i][j]
			}

			if i > 0 {
				fmt.Print(" ", symbols[j], " ")
			}

			fmt.Print(homeWork[i][j])
		}
		fmt.Print(" = ", val)
		fmt.Println("\n=========")
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
