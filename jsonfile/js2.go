package main

import (
	"errors"
	"fmt"
	"os"
)

func checkFileExists(filePath string) bool {

	_, err := os.Stat(filePath)

	return !errors.Is(err, os.ErrNotExist)

}

func main() {
	path := "examplee.txt"

	if checkFileExists(path) {

		fmt.Println("file already exists")

	} else {

		file, err := os.Create("examplee.txt")
		if err != nil {

			fmt.Println("error creating file", err)
			return
		}
		defer file.Close()
		fmt.Println("file creted successfully")
	}
	//open an existing file

	file, err := os.Open("example.txt")

	if err != nil {
		fmt.Println("erroe opening file:", err)
		return
	}

	defer file.Close()

	fmt.Println("file opened successfully")

}
