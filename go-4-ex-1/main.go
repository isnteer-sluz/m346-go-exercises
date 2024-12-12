package main

import "fmt"

// TODO: implement the function computeGrade
func computeGrade(gotPoints int, maxPoints int) float64 {

	var result float64 = ((float64(gotPoints) / float64(maxPoints)) * 5) + 1

	if result < 1 {
		return 1
	} else if result > 6 {
		return 6
	} else {
		return result
	}
}

func main() {
	fmt.Println(computeGrade(0, 100))
}
