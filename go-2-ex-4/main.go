package main

func main() {
	// declare a type for Student (with first and last name)
	type Student struct {
		FirstName string
		LastName  string
	}
	
	// declare a type for Class (consisting of multiple students)
	type Class struct {
		Students []Student
	}

	// declare a map of modules being attended by multiple classes
	modules := map[int]Class{
		104: {
			Students: []Student{
				{FirstName: "Nils", LastName: "Steiner"},
				{FirstName: "Lars", LastName: "Maler"},
			},
		},
		117: {
			Students: []Student{
				{FirstName: "Gustaf", LastName: "Gans"},
				{FirstName: "Donald", LastName: "Duck"},
			}
		},
		346: {
			Students: []Student{
				{FirstName: "Apo", LastName: "R E D"},
				{FirstName: "Leon", LastName: "Machère"},
			}
		}
	}

// TODO: Aufgabe muss ich nochmals überarbeiten weil ich ein Fehler gemacht habe


	// output everything using fmt.Println()
	fmt.Println("Modul 104:", modules[104])
	fmt.Println("Modul 117:", modules[117])
	fmt.Println("Modul 346:", modules[346])

}
