package main

import (
	"fmt"
	"time"
)


func routine1 (c chan int,playload int){

c<-playload	
fmt.Println("kuy1")
c<-playload	
fmt.Println("kuy2")
c<-playload	
fmt.Println("kuy3")
c<-playload	
fmt.Println("kuy4")
c<-playload	
fmt.Println("kuy5")
c<-playload	
fmt.Println("kuy6")


}
func main() {
	ch := make(chan int,7) //เป้นการจองที่(buffer)ให้ตัวแปล

	go routine1(ch,1) // มันจะแยกไปrunงานย่อย ก่อนแล้วค่อยเข้ามา
	go routine1(ch,2)
	go routine1(ch,3)
	
	fmt.Println("run1")
	fmt.Println(<-ch)

	time.Sleep(2*time.Second)
}