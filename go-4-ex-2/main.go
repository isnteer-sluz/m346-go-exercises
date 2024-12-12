package main

// TODO: implement the function computeHypotenuse using math.Pow and math.Sqrt
func computeHypotenuse(a float64, b float64) float64 {
	return math.Sqrt(math.Pow(a, 2) + math.Pow(b, 2))
}

func main() {
	// TODO: call the function computeHypotenuse
	// Kannst du mir ein testing für die Funktion schreiben mit kommentaren was das resultat
	fmt.Println(computeHypotenuse(3, 4)) // 5
	fmt.Println(computeHypotenuse(5, 12)) // 13
	fmt.Println(computeHypotenuse(8, 15)) // 17
	fmt.Println(computeHypotenuse(7, 24)) // 25
}