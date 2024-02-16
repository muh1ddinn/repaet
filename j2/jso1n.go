package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Animal struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Class string `json:"class"`
}

func main() {

	fileee, errr := os.Create("animal.json")

	if errr != nil {

		fmt.Println("err while creating ", errr)

		return
	}
	defer fileee.Close()

	animals := []Animal{

		Animal{"rex", 4, "mannaml"},
		Animal{"rex2", 2, "mannaml"},
		Animal{"rex3", 3, "mannaml"},
	}

	encoder := json.NewEncoder(fileee)
	encoder.SetIndent("", " ")
	errr = encoder.Encode(animals)

	if errr != nil {

		fmt.Println("error while encoding animal to json ,err:", errr)
	}
}
