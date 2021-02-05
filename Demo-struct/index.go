package main

import "fmt"

func main() {
	var p1 person
	p1.firstName = "john"
	p1.lastName = "deo"
	p1.age = 30
	p1.eyeColor = "Blue"

	fmt.Println(p1)

	var p2 person
	p2.firstName = "parin"
	p2.lastName = "manchada"
	p2.age = 30
	p2.eyeColor = "Red"

	shoW(p1)
	shoW(p2)
}

type person struct {
	firstName string
	lastName  string
	age       int
	eyeColor  string
}

func shoW(p person) {

	fmt.Println(p)

}
