package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Go workshop 2")

	x := 100
	msg := "Lorem string"
	sleepSec := 3 * time.Second

	fmt.Printf("%T \n", x)
	fmt.Println(time.Now())
	time.Sleep(sleepSec)
	fmt.Println(time.Now())
	fmt.Println("This line awake after ", sleepSec)
	fmt.Printf("%T", sleepSec)
	fmt.Printf("%T, %s \b", msg, msg)
}
