package main

import "fmt"

func main() {
	var msg string = "kuy Parin"
	var msgPointer *string = &msg

	fmt.Println(msg)
	fmt.Println(msgPointer)
	fmt.Println(*msgPointer)
	changMessage(msgPointer)
	fmt.Println(msg)
}

func changMessage(aPointer *string) {
	*aPointer = "Noob Parin"
}
