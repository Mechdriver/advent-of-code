package main

import (
	"mechdriverAdventOfCode/secret_entrance"
	"fmt"
)

func main() {
	password := secret_entrance.Part2()

	fmt.Printf("Password: %d\n", password)
}
