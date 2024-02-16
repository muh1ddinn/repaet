package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Producct struct {
	Id                 int     `json:"id"`
	Title              string  `json:"title"`
	Description        string  `json:"description"`
	Price              float64 `json:"price"`
	DiscountPercentage float64 `json:"discountpercentage"`
	Rating             float64 `json:"rating"`
	Stock              float64 `json:"stock"`
	Brand              string  `json:"brand"`
	Category           string  `json:"category"`
	Thumbnail          string  `json:"thumbnail"`
}

func reead(filename string) []Producct {

	filee, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("you have err while reading jsonfile ", err)
		return nil
	}

	var prod []Producct

	err = json.Unmarshal(filee, &prod)
	if err != nil {

		fmt.Println("you have err while unmarshaling time ", err)
		return nil
	}

	return prod

}

func writefuntojs(file, reslut string) {

	fil, err := os.Create(file)
	if err != nil {

		fmt.Println(" you have err while creating file ")

		return
	}
	defer fil.Close()

	_, err = fil.WriteString(reslut)

	if err != nil {

		fmt.Println("err while writing to file ")
	}

}

func main() {

	var max float64
	max = 0

	proo := reead("dataa.json")

	for i := 0; i < len(proo); i++ {

		if proo[i].DiscountPercentage > max {

			max = proo[i].DiscountPercentage

		}

	}
	fmt.Println("max vale of DiscountPercentage: ", max)

}
