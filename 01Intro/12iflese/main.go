package main

import "fmt"

func main() {
	fmt.Println("If else in Golang.")

	loginCount := 53
	var result string

	if loginCount < 10 {
		result = "Regular User"
	} else if loginCount < 50 {
		result = "Boomer Mama! Express User"
	} else {
		result = "Bwahahahaha more than 50 logins you fucker"
	}

	fmt.Println(result)

	if 9%2 == 0 {
		fmt.Println("Number is EVEN")
	} else {
		fmt.Println("Number is ODD")
	}

	if num := 13; num < 10 {
		fmt.Println("Number is LESS THAN 10")

	} else {
		fmt.Println("Number is NOT LESS THAN 10")

	}

}
