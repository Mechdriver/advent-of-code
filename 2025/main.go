package main

import (
	// "fmt"
	"fmt"
	"mechdriverAdventOfCode/lobby"
	// "mechdriverAdventOfCode/gift_shop"
	// "mechdriverAdventOfCode/secret_entrance"
)

func main() {
	// password := secret_entrance.Part2()

	// fmt.Printf("Password: %d\n", password)

	// summedIds := gift_shop.Part2()

	totalJoltage := lobby.Part1()

	fmt.Printf("Total Joltage: %d\n", totalJoltage)
}
