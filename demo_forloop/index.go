package main

import (
	"fmt"
)

func main() {
	fnFor()
	fnWhile()
}

func fnFor() {
	for index := 0; index < 10; index++ {
		fmt.Printf("For_index %d\n", index)
	}
}

func fnWhile() {
	count := 0
	for count < 5 {

		count++
		fmt.Printf("While_count %d\n", count)
	}
}
