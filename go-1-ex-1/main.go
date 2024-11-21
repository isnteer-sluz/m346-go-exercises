package main

import "fmt"

// Deklaration
var firstName string = "Nils"
var lastName string = "Steiner"
var dayOfBirth int = 3
var monthOfBirth int = 12
var yearOfBirth int = 2007
var numberOfSiblings int = 0
var heightInMeters float64 = 1.90
var zodiacSign rune = 'S'

func main() {
	// TODO: Declare and initialize the variables being used in the output!

	// Ausgabe
	fmt.Printf("Vor- und Nachname: %s %s\n", firstName, lastName)
	fmt.Printf("Geburtsdatum: %d.%d.%d\n", dayOfBirth, monthOfBirth, yearOfBirth)
	fmt.Printf("Anzahl Geschwister: %d\n", numberOfSiblings)
	fmt.Printf("Grösse (in Metern): %.2f\n", heightInMeters)
	fmt.Printf("Sternzeichen: %c\n", zodiacSign)
}
