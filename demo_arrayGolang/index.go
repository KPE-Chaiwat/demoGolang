package main

import "fmt"

func main() {
	var myhero [4]string
	myhero[0] = "PA"
	myhero[1] = "jug"
	myhero[2] = "rubic"
	myhero[3] = "viper"

	fmt.Println(myhero[3])
	for _, i := range myhero {
		fmt.Println("My hero use to playing", i)
	}
	
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
	for _, i := range primes {

		fmt.Println(i)
	}
}
