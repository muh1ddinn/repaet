package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

func ceckfex(filep string) bool {

	_, err := os.Stat(filep)
	return !errors.Is(err, os.ErrNotExist)

}

func main() {

	filep := ""
	fmt.Scan(&filep)

	if ceckfex(filep + ".txt") {

		var rec1, rec2, rec3 string = filep, filep, filep

		rec1 += strconv.Itoa((rand.Intn(20))) + ".txt"
		rec2 += strconv.Itoa((rand.Intn(20))) + ".txt"
		rec3 += strconv.Itoa((rand.Intn(20))) + ".txt"

		fmt.Println("File already exist,you can create files with name like this")

		if !ceckfex(rec1) {
			fmt.Printf("%v", rec1)
		}
		if !ceckfex(rec2) {
			fmt.Printf("%v", rec2)

		}
		if !ceckfex(rec3) {

			fmt.Println("%v", rec3)
		}

		fmt.Println()

	} else {

		_, err := os.Create(filep + ".txt")
		if err != nil {
			panic(err)
		}

	}
}
