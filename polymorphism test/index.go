package main
import (
	"fmt"
	"math"
)

func main() {

	r := rectangle{width: 10, height: 50}
	c := circle{radius: 200}	

	fmt.Printf("Circle: %f\n", getArea(c))
	fmt.Printf("Rectangle: %f\n", getArea(r))
	

	
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

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r rectangle) area() float64 {
	return r.width * r.height
}

func getArea(s shape) float64{
	return s.area()
}

func showInfo(s shape){
   
}
