package main

import "fmt"

type FullName struct {
	// TODO: add fields
	FirstName string
	LastName string
}

// TODO: declare a structure for birth date

type BirthDate struct {
	Day int
	Month int
	Year int
}

type Profile struct {
	// TODO: embed full name and birth date information
	FullName
	BirthDate
	NumberOfSiblings byte
	ZodiacSign       rune
}

func main() {
	var me = Profile{
		// TODO: set name and birth date information
FullName: FullName{
	FirstName: "Nils",
	LastName: "Steiner",
	},
		BirthDate: BirthDate{
		Day: 3,
		Month: 12,
		Year: 2007,
	},
		NumberOfSiblings: 0,   // TODO: adjust
		ZodiacSign:       'S', // TODO: adjust
	}
	fmt.Println(me)

	fmt.Println("Siblings Before:", me.NumberOfSiblings)
	// TODO: imagine, you get a little brother or sister
	me.NumberOfSiblings++
	fmt.Println("Siblings After:", me.NumberOfSiblings)
}
