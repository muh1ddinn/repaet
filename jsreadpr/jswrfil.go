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

	filee, err := os.Create("animalsss.json")

	if err != nil {
		fmt.Println("you have err with creating file:", err)

	}
	defer filee.Close()

	animalss := []Animal{}

	animalss = append(animalss,

		Animal{"rex", 35, "sjd"},
		Animal{"rex", 325, "454"},
		Animal{"rex", 33, "rrtr"},
	)

	encoder := json.NewEncoder(filee)
	encoder.SetIndent("", " ")
	err = encoder.Encode(animalss)

	if err != nil {
		fmt.Print("error", err)
	}

}
