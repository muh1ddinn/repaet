package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Product struct {
	Id                 int     `json:"id"`
	Title              string  `json:"title"`
	Description        string  `json:"description"`
	Price              float64 `json:"price"`
	DiscountPercentage float64 `json:"discountPercentage"`
	Rating             float64 `json:"rating"`
	Stock              float64 `json:"stock"`
	Brand              string  `json:"brand"`
	Category           string  `json:"category"`
	Thumbnail          string  `json:"thumbnail"`
}

func readfuncfromjson(filename string) []Product {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("error with reading file:", err)
		return nil
	}

	var products []Product
	err = json.Unmarshal(data, &products)
	if err != nil {
		fmt.Println("error unmarshaling json:", err)
		return nil
	}
	return products
}

func writeresultfile(filename, result string) {
	fil, err := os.Create(filename)
	if err != nil {
		fmt.Println("err with creating text file ", err)
		return
	}
	defer fil.Close()

	_, err = fil.WriteString(result)
	if err != nil {
		fmt.Println("err with writing to file", err)
	}
}

func main() {

	

	var max float64

	max = 0
	prod := readfuncfromjson("dataa.json")

	for i := 0; i < len(prod); i++ {

		if prod[i].DiscountPercentage > max {
			max = prod[i].DiscountPercentage
		}
	}
	fmt.Println("Maximum discount percentage:", max)

	a := fmt.Sprint(max)

	writeresultfile("output.txt", a)

}
