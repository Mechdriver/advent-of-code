package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	buffer, err := os.ReadFile("test_input.txt")

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

	fmt.Printf("Password: %d\n", password)
}
