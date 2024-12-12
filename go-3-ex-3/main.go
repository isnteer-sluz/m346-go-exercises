package main

import "fmt"

const (
	Lower = 1
	Upper = 30
)

func main() {
	for Lower := 1; Lower < Upper; Lower++ {
		if Lower%3 == 0 && Lower%5 == 0 {
			fmt.Printf("FizzBuzz ")
		} else if Lower%3 == 0 {
			fmt.Printf("Fizz ")
		} else if Lower%5 == 0 {
			fmt.Printf("Buzz ")
		} else {
			fmt.Printf("%d ", Lower)
		}
		fmt.Println()
	}
}
