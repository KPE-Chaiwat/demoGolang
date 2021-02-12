package main

import (
	"fmt"
	"time"
)

func main() {
	go run1()
	go run2()

	time.Sleep(5 * time.Second)
}

func run1() {
	for i := 0; i < 100; i++ {
		fmt.Println("Kuy Parin running 1")
	}
}

func run2() {
	for i := 0; i < 100; i++ {
		fmt.Println("Kuy parin running 2")
	}
}
