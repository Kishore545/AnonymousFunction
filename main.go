package main

import "fmt"

func Foo() {
	fmt.Println("Foo Ran")
}

func main() {
	Foo()
	func() {
		fmt.Println("Anonymous Function Ran")
	}()

	func(s string) {
		fmt.Println("This is an Anonymous Func showing My Name", s)
	}("Chiranjeevi")
}
