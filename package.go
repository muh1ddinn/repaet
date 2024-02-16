package main

import (
	"fmt"
)

/*
func div(a, b int) (int, error) {

	if b == 0 {

		return 0, errors.New("division by zero")

	}

	return a / b, nil

}

func main() {

	result, err := div(10, 2)

	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println("result:", result)
	}

	result, err = div(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}
}
*/

type Students struct {
	Name  string
	Cours int
}

func main() {

	Studentss := []Students{
		Students{

			Name:  "11",
			Cours: 3,
		},
		Students{
			Name:  "12",
			Cours: 2,
		},
		Students{
			Name:  "23",
			Cours: 1,
		},
		Students{
			Name:  "33",
			Cours: 3,
		},
		Students{
			Name:  "33",
			Cours: 4,
		},
		Students{
			Name:  "33",
			Cours: 1,
		},
	}

	k := len(Studentss)
	for i := 0; i < k; i++ {
		for j := 0; j < k-i-1; j++ {

			if Studentss[j].Cours > Studentss[j+1].Cours {

				Studentss[j], Studentss[j+1] = Studentss[j+1], Studentss[j]

			}
		}
	}
	// Print sorted Studentss slice

	for _, student := range Studentss {

		fmt.Printf("Name: %s, Cours: %d\n", student.Name, student.Cours)
	}

}
