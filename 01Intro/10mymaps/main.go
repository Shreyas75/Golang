package main

import "fmt"

func main() {
	fmt.Println("Maps in Golang.")

	languages := make(map[string]string)

	languages["JS"] = "JavaScript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("List of all languages: ", languages)
	fmt.Println("JS is short for: ", languages["JS"])
	fmt.Println("PY is short for: ", languages["PY"])

	delete(languages, "RB")
	fmt.Println("List of languages after deletion: ", languages)

	// Loops are pretty interesting in golang

	for _, value := range languages {
		fmt.Printf("For key v, values is %v\n", value)
	}

}
