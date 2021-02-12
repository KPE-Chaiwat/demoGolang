package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 1) //เป้นการจองที่(buffer)ให้ตัวแปล

	ch <- 1 //send to chanal ch
	fmt.Println("run1")
	fmt.Println(<-ch)

	ch <- 2
	fmt.Println("run2")
	fmt.Println(<-ch)

	time.Sleep(5 * time.Second)
}
