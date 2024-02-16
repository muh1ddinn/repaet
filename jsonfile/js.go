package main

import (
	"encoding/json"
	"fmt"
)

type Phone struct {
	Model    string `json:"mark"`
	Mark     string `json:"model"`
	Year     int    `json:"year"`
	Performs string ` json:"perfoms"`
}

func main() {

	phone := Phone{

		Model:    "iphone",
		Mark:     "15 pro",
		Year:     2023,
		Performs: "batter",
	}

	//json.Marshal
	marshelled, err := json.Marshal(phone)

	if err != nil {
		fmt.Println("error occured:", err)
	}

	fmt.Println("marshelled :", string(marshelled))

	unmarshel := Phone{}

	err = json.Unmarshal(marshelled, &unmarshel)

	if err != nil {
		fmt.Println("error occured:", err)
	}

	fmt.Printf("unmarshel %+v\n:", unmarshel)
	fmt.Println(phone)
}
