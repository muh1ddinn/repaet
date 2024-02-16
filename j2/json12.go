package main

import (
	"fmt"
	"os"
)

func main() {
	/*
		file, err := os.Create("person.txt")

		if err != nil {

			fmt.Println("error while creating file,err: ", err)
		}

		defer file.Close()

		person := []Person{}

		n := 0

		fmt.Println("enter the count of person : ")

		fmt.Scan(&n)

		var name, gmail, job string
		for i := 0; i <= n; i++ {

			var persons = Person{}
			fmt.Printf("enter the %v animal's person:\n", i)
			fmt.Scanln(&name)
			fmt.Printf("enter the %v animal's person:\n", i)
			fmt.Scanln(&gamil)
			fmt.Printf("enter the %v animal's person:\n", i)
			fmt.Scanln(&job)

			person = append(person, persons)
		}


		encoder :

	*/

	file, err := os.Create("person.txt")

	if err != nil {

		fmt.Println("error while creating file,err: ", err)
	}

	defer file.Close()

	n := 0

	fmt.Println("enter the count of person : ")

	fmt.Scan(&n)

	var name, gmail, job string

	file, err := os.OpenFile("person.txt")

	if err != nil {
		panic(err)
	}
	defer file.Close()

	words := []string{}
	for i := 0; i <= n; i++ {

		fmt.Printf("enter the %v person's name:\n", i)
		fmt.Scanln(&name)

		fmt.Printf("enter the %v person's gmail:\n", i)
		fmt.Scanln(&gmail)
		fmt.Printf("enter the %v person's job:\n", i)
		fmt.Scanln(&job)

		_, err := file.WriteString(fmt.Sprintf("Person %d: Name: %s, Gmail: %s, Job: %s\n", i, name, gmail, job))
		if err != nil {
			fmt.Println("Error writing to file: ", err)
			return
		}
	}

	fmt.Println("Data has been written to person.txt")
}
