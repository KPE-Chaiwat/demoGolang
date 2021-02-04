package main

import "fmt"

func main() {

	var myHero = map[string]int{"Pa": 20,
		"voide":  50,
		"drow":   40,
		"Mirana": 40}

	fmt.Println(myHero)
	fmt.Println(myHero["drow"])

	var colors = make(map[string]string)
	colors["red"] = "#ff0"
	colors["green"] = "#00f"
	colors["blue"] = "#008"
	colors["black"] = "#000"

	fmt.Println("", colors)
	fmt.Println("", colors["black"])

	var courses = make(map[string]map[string]int)

	courses["thai"] = make(map[string]int)
	courses["thai"]["price"] = 1000
	courses["thai"]["discount"] = 200
	courses["thai"]["period"] = 30
	
	courses["math"] = make(map[string]int)
	courses["math"]["price"] = 3000
	courses["math"]["discount"] = 200
	courses["math"]["period"] = 20
	
	fmt.Println(courses)
	fmt.Println(courses["thai"]["price"])
	fmt.Println(courses["math"]["discount"])

}
