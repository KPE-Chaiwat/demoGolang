package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {

	r := rectangle{width: 10, height: 50}
	c := circle{radius: 200}

	fmt.Printf("Circle: %f\n", getArea(c))
	fmt.Printf("Rectangle: %f\n", getArea(r))

	showInfo(c)
	showInfo(r)

}

type circle struct {
	radius float64
}
type shape interface {
	area() float64
}

type rectangle struct {
	width, height float64
}

//เป็นการเอา struct มายัดใน ฟังก์ชัน เพื่อรที่จะสามารถใช้ตัวแปลในstruct ตัวที่เอายัดได้
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r rectangle) area() float64 {
	return r.width * r.height
}

func getArea(s shape) float64 {
	return s.area()
}


func showInfo(s shape){
	t := reflect.TypeOf(s).Name()
	switch t {
	case "rectangle":
	  r := s.(rectangle); // cast from shape to rectangle
	  fmt.Printf("Rect  width: %f, height: %f\n", r.width, r.height)
	  break
	case "circle":
	  c := s.(circle); // cast from shape to circle
	  fmt.Printf("Circle radius: %f\n", c.radius)
	}	
}