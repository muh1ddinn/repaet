package main

import "fmt"

type Student struct {
	name        string
	scholarship int
	course      int
}

func main() {
	students := []Student{
		{"Student1", 5000, 2},
		{"Student2", 5000, 2},
		{"Student3", 5000, 1},
		{"Student4", 5000, 3},
		{"Student5", 5000, 2},
		{"Student6", 5000, 1},
		{"Student7", 5000, 2},
		{"Student8", 5000, 3},
		{"Student9", 5000, 2},
		{"Student10", 5000, 1},
	}

	totalScholarship := 0
	for _, student := range students {
		if student.course == 2 {
			totalScholarship += student.scholarship

		}
	}

	fmt.Printf("Total scholarships paid to all 2nd year students: %d\n", totalScholarship/5000)
}
