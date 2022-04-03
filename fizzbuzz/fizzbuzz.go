package main

import (
	"fmt"
)

func main() {
	FizzBuzz(18)
}

func FizzBuzz(n int32) {
	var i int32
	for i = 1; i <= n; i++ {
		if i%15 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {

			fmt.Println(i)
		}
	}
}
