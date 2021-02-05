package main

import "fmt"

func main() {
	var p1 person
	p1.firstName = "john"
	p1.lastName = "deo"
	p1.age = 21
	p1.eyeColor = "Blue"

	fmt.Println(p1)

	var p2 person
	p2.firstName = "parin"
	p2.lastName = "manchada"
	p2.age = 30
	p2.eyeColor = "Red"

	shoW(p1)
	updatAge(&p1)
	shoW1(&p2)
	updatAge(&p2)
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
func shoW1(p *person) {

	fmt.Println(*p)

}

func updatAge(p *person) {

	p.age = p.age + 1

	fmt.Printf("update age=%d\n", p.age)
}
