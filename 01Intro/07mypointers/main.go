package main

import "fmt"

func main() {
	fmt.Println("Welcome to a class on Pointers.")

	// var ptr *int
	// fmt.Println("Value of pointer is ", ptr)

	myNumber := 23

	var ptr *int = &myNumber
	fmt.Println("The Address of myNumber is ", &myNumber)
	fmt.Println("The Value inside of pointer is ", ptr)
	fmt.Println("The Value of pointer is ", *ptr)
	fmt.Println("The Address of pointer is ", &ptr)

	*ptr = *ptr + 2
	fmt.Println("New value inside myNumber is: ", myNumber)

}
