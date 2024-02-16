package main

import "fmt"

func main() {

	var m map[int]string

	if m == nil {
		fmt.Println("true")
	} else {
		fmt.Println("flase")
	}
	map2 := map[int]string{

		90: "ijef",
		88: "dfdg",
		87: "677",
		45: "hjhj",
	}

	fmt.Println("maps : ", map2)

}
