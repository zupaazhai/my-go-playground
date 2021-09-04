package main

import "fmt"

func foo() {
	fmt.Println("Message from foo()")
}

func bar() string {
	return "Return message as string"
}

func baz(name string) string {
	return "My name is " + name
}

func main() {
	fmt.Println("Workshop 3 function")
	foo()

	fmt.Println("bar()", bar())
	fmt.Println("func has return", baz("Supachai"))
}
