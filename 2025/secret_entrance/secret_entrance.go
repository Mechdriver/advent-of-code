package secret_entrance

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Part1() int {
	buffer, err := os.ReadFile("./inputs/day1/test_input.txt")

	if err != nil {
		panic(err)
	}

	inputs := strings.Split(string(buffer), "\n")

	dial := 50
	password := 0

	for i := range inputs {
		line := inputs[i]
		direction := line[:1]
		amount, err := strconv.Atoi(line[1:])

		if err != nil {
			panic(err)
		}

		if direction == "L" {
			amount *= -1
		}

		amount %= 100

		dial += amount

		if dial < 0 {
			dial += 100
		} else if dial > 99 {
			dial -= 100
		}

		if dial == 0 {
			password++
		}
	}

	return password
}

func Part2() int {
	buffer, err := os.ReadFile("./inputs/day1/test_input.txt")

	if err != nil {
		panic(err)
	}

	inputs := strings.Split(string(buffer), "\n")

	dial := 50
	password := 0

	for i := range inputs {
		line := inputs[i]
		direction := line[:1]
		amount, err := strconv.Atoi(line[1:])

		if err != nil {
			panic(err)
		}

		password += amount / 100
		amount %= 100

		if direction == "L" {
			amount *= -1
		}

		wasZero := dial == 0

		dial += amount

		if dial < 0 {
			dial += 100

			if !wasZero {
				password++
			}
		} else if dial > 99 {
			dial -= 100

			if !wasZero {
				password++
			}
		} else if dial == 0 {
			password++
		}

		// fmt.Printf("Input: %s\nDial: %d\nPassword: %d\n\n", line, dial, password)

	}

	return password
}
