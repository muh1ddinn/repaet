package main

import "fmt"

type author struct {
	name   string
	brach  string
	salart int
	age    int
}

// create method
// show is method name
func (a author) sho
w() {

	fmt.Println("author's name : ", a.name)
	fmt.Println("brach name: ", a.brach)

}

func main() {

	res := author{
		name:   "sona",
		brach:  "cse",
		salart: 550,
	}

	res.show()

}
