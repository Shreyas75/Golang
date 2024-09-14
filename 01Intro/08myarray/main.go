package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in Golang.")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Mango"
	fruitList[3] = "Pineapple"

	fmt.Println("Fruit list is: ", fruitList)
	fmt.Println("Fruit list is: ", len(fruitList))

	var vegList = [3]string{"potato", "beans", "carrot"}
	fmt.Println("Vegy list is: ", len(vegList))

}
