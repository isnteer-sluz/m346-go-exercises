package main

import "fmt"

const (
	Aries       = '\u2648' // Widder
	Taurus      = '\u2649' // Stier
	Gemini      = '\u264a' // Zwillinge
	Cancer      = '\u264b' // Krebs
	Leo         = '\u264c' // Löwe
	Virgo       = '\u264d' // Jungfrau
	Libra       = '\u264e' // Waage
	Scorpius    = '\u264f' // Skorpion
	Sagittarius = '\u2650' // Schütze
	Capricornus = '\u2651' // Steinbock
	Aquarius    = '\u2652' // Wassermann
	Pisces      = '\u2653' // Fische
)

func outputDateRange(zodiacSign rune) {
	fmt.Printf("%c: ", zodiacSign)
	// TODO: Replace if, else if branching with switch/case.
	switch {
		case (month == 3 && day >= 21) || (month == 4 && day <= 19):
			return "Aries" // Widder
		case (month == 4 && day >= 20) || (month == 5 && day <= 20):
			return "Taurus" // Stier
		case (month == 5 && day >= 21) || (month == 6 && day <= 20):
			return "Gemini" // Zwillinge
		case (month == 6 && day >= 21) || (month == 7 && day <= 22):
			return "Cancer" // Krebs
		case (month == 7 && day >= 23) || (month == 8 && day <= 22):
			return "Leo" // Löwe
		case (month == 8 && day >= 23) || (month == 9 && day <= 22):
			return "Virgo" // Jungfrau
		case (month == 9 && day >= 23) || (month == 10 && day <= 22):
			return "Libra" // Waage
		case (month == 10 && day >= 23) || (month == 11 && day <= 21):
			return "Scorpius" // Skorpion
		case (month == 11 && day >= 22) || (month == 12 && day <= 21):
			return "Sagittarius" // Schütze
		case (month == 12 && day >= 22) || (month == 1 && day <= 19):
			return "Capricornus" // Steinbock
		case (month == 1 && day >= 20) || (month == 2 && day <= 18):
			return "Aquarius"
		case (month == 2 && day >= 19) || (month == 3 && day <= 20):
			return "Pisces"
		default:
			return "Invalid Date"
		}

	// TODO: Define all 12 cases...
	switch zodiacSign {
		case Aries:
			fmt.Println("21.03. - 20.04")
		case Taurus:
			fmt.Println("21.04. - 21.05")
		case Gemini:
			fmt.Println("21.04. - 21.05")
		case Cancer:
			fmt.Println("21.04. - 21.05")
		case Leo:
			fmt.Println("21.04. - 21.05")
		case Virgo:
			fmt.Println("21.04. - 21.05")
		case Libra:
			fmt.Println("21.04. - 21.05")
		case Scorpius:
			fmt.Println("21.04. - 21.05")
		case Sagittarius:
			fmt.Println("21.04. - 21.05")
		case Capricornus:
			fmt.Println("21.04. - 21.05")
		case Aquarius:
			fmt.Println("21.04. - 21.05")
		case Pisces:
			fmt.Println("21.04. - 21.05")
		default:
			fmt.Println("Error no zodiac sign found!")
		}
}

func main() {
	for zodiacSign := Aries; zodiacSign <= Pisces; zodiacSign++ {
		outputDateRange(zodiacSign)
	}
}
