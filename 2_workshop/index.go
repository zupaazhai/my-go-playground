package main

import "fmt"

const TOKEN_KEY = "thisistokenkey_itwillsecret"

func main() {
	fmt.Println("Workshop 2")

	var number int = 0
	var message string = "WS 2"
	var isOk bool

	fmt.Println(number, message)
	fmt.Println("isOk => ", isOk)

	var members [2]string
	members[0] = "John"
	members[1] = "Sandy"

	players := [2]string{}
	players[0] = "Adam"

	fmt.Println(members)
	fmt.Println(players)

	//constant
	fmt.Println("TOKEN_KEY", TOKEN_KEY)

	var name string

	fmt.Printf("name is empty '%s' \n", name)
}
