package main

import "fmt"

func main() {
	fmt.Println("Structs in Golang.")
	// No Inheritance in Golang; No Super or Parent.

	shreyas := User{"Shreyas", "shreyas@go.dev", true, 23}
	fmt.Println(shreyas)
	fmt.Printf("shreyas details are %+v\n", shreyas)
	fmt.Printf("Name is %v and email is %v\n", shreyas.Name, shreyas.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
