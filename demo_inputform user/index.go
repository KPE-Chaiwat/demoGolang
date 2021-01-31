package main

import ("fmt"
)

func main() {

	var a int
	var b int
	fmt.Println("number a: ")
	fmt.Scanln(&a)
	fmt.Println("number b: ")
	fmt.Scanln(&b)
	
	fmt.Print(a, "sum", b, "=")
	fmt.Println(sumVarlue(a,b))
}

func sumVarlue( a int,  b int)  int{
	return a + b
}
