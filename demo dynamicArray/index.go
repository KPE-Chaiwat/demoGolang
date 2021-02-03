package main

import "fmt"

func main() {
	var varlue = make([]int, 0, 10)
	varlue = append(varlue, 2)
	varlue = append(varlue, 4)
	varlue = append(varlue, 6)
	showSlice(varlue)
	var varlue1 []int
	varlue1 = append(varlue1, 2)
	varlue1 = append(varlue1, 4)
	varlue1 = append(varlue1, 8)
	showSlice(varlue1)
}

func showSlice(x []int) {

	fmt.Printf("len=%d cap=%d slice=%v\n ", len(x), cap(x), x)

}
