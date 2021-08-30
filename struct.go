package main

import (
	"fmt"
)

type player struct {
	name   string
	number int
}

func removeByIndex(arr []player, index int) []player {
	return append(arr[:index], arr[index+1:]...)
}

func main() {
	teamA := []player{}

	teamA = append(teamA, player{name: "John", number: 10})
	teamA = append(teamA, player{name: "Alex", number: 9})
	teamA = append(teamA, player{name: "Smith", number: 3})

	// remove
	teamA = removeByIndex(teamA, 1)

	fmt.Println(teamA[:1])
}
