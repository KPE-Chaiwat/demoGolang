package main

import "fmt"

func intSeq() func() int {
    i := 0
    return func() int {
        i++
        return i
    }
}

 func main () {

	nextInt := intSeq()
fmt.Println(nextInt())
fmt.Println(nextInt())
fmt.Println(nextInt())
fmt.Println(nextInt())

nextInts := intSeq()
fmt.Println(nextInts())
fmt.Println(nextInts())

nextString:=myRerturn()

fmt.Println(nextString())

fmt.Println(nextString())
fmt.Println(nextString())
fmt.Println(nextString())



fmt.Println(myRerturn()())

 }


 func myRerturn () func () string{

y:=0
return func() string {
	y++
	return fmt.Sprintf("y = %d",y)
}

 }