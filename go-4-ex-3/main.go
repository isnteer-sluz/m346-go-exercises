package main

import "math"

// TODO: implement the function computeQuadraticFormula
func computeQuadraticFormula(a float64, b float64, c float64) []float64 {

	var diskriminante float64 = math.Pow(b, 2) - 4*a*c

	if diskriminante > 0 {

		var x1 float64 = (-b + math.Sqrt(diskriminante)) / (2 * a)
		var x2 float64 = (-b - math.Sqrt(diskriminante)) / (2 * a)

		return []float64{x1, x2}

	} else if diskriminante == 0 {

		var x1 float64 = -b / (2 * a)

		return []float64{x1}

	} else {
		return nil
	}

}

func main() {
	// TODO: call the function computeQuadraticFormula
	fmt.Println(computeQuadraticFormula(3, 4, 1)) // zwei Lösungen
	fmt.Println(computeQuadraticFormula(2, 4, 2)) // eine Lösung
	fmt.Println(computeQuadraticFormula(3, 4, 2)) // keine Lösung
}
