package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to Slices Lecture.")

	var fruitList = []string{"Apple", "Orange", "Grape"}
	fmt.Printf("Type of fruitlist is %T\n", fruitList)

	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Println(fruitList)

	fruitList = append(fruitList[:3])
	fmt.Println(fruitList)

	highScores := make([]int, 4)

	highScores[0] = 234
	highScores[1] = 344
	highScores[2] = 154
	highScores[3] = 634

	highScores = append(highScores, 555, 666, 777)

	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))

	sort.Ints(highScores)
	fmt.Println(highScores)

	// how to remove a value from slices based on index

	var courses = []string{"reactjs", "js", "swift", "python", "ruby"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)

}
