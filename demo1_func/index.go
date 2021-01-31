package main

import "fmt"

func main() {
	const a, b = 1500, 300
	var x,y = swap(a,b)
	showmessage()
	sum(200, 500)
	message("kuyparin")
	showlevel("mirana", 3)
	fmt.Printf("%d+%d=%d\n", a, b, sum1(a, b))
	fmt.Printf("%d,%d=>%d,%d\n", a, b, x,y)
}

func showmessage() {

	fmt.Println("parin")

}

func sum(a int, b int) {
	c := a + b

	fmt.Println(c)

}
func sum1(a int, b int) int {

	return a + b

}

func message(msg string) {

	fmt.Println(msg)

}
func showlevel(y string, x int) {
	fmt.Printf(y)
	fmt.Println(x)

}
func swap(x int, y int) (int, int) {

	return y,x

}
