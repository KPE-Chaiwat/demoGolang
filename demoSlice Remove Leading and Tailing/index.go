package main

import "fmt"

func main() {
	var number = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}

	showSlice(number)

	//leading number

	number = number[1:len(number)] // ตัดตัวที่1 ถึงตัวสุดท้าย
	showSlice(number)
	number = number[1:len(number)] 
	showSlice(number)

	//tailing remove 
	var number1 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}

	number1 = number1[0:len(number1)-1] 
	showSlice(number1)

	number1 = number1[0:len(number1)-1] 
	showSlice(number1)
}
func showSlice(x []int) {

	fmt.Printf("len=%d cap=%d slice=%v\n ", len(x), cap(x), x)
}
