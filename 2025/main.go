package main

import (
	// "fmt"
	"fmt"
	"mechdriverAdventOfCode/gift_shop"
	// "mechdriverAdventOfCode/secret_entrance"
)

func main() {
	// password := secret_entrance.Part2()

	// fmt.Printf("Password: %d\n", password)

	summedIds := gift_shop.Part2()

	fmt.Printf("Bad ID Sum: %d\n", summedIds)
}
