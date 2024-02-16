package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	data, err := os.ReadFile("data.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Split the data into individual elements
	elements := strings.Fields(string(data))

	// Open files for writing
	numbersFile, err := os.Create("numbers.txt")
	if err != nil {
		fmt.Println("Error creating numbers.txt:", err)
		return
	}
	defer numbersFile.Close()

	lettersFile, err := os.Create("letters.txt")
	if err != nil {
		fmt.Println("Error creating letters.txt:", err)
		return
	}
	defer lettersFile.Close()

	var totalSum int
	var numbersData, lettersData strings.Builder

	// Loop through the elements
	for _, element := range elements {
		// Check if the element is a number
		if num, err := strconv.Atoi(element); err == nil {
			// Write the number to numbersData
			numbersData.WriteString(element + "\n")
			totalSum += num
		} else {
			// Check if the word contains only 'A', 'O', 'a', 'o'
			isAO := true
			for _, char := range element {
				if !unicode.In(char, unicode.Latin) || (char != 'A' && char != 'O' && char != 'a' && char != 'o') {
					isAO = false
					break
				}
			}

			// Write the word to lettersData if it contains only 'A', 'O', 'a', 'o'
			if isAO {
				lettersData.WriteString(element + "\n")
			}
		}
	}

	// Write the data to the files
	_, err = numbersFile.WriteString(numbersData.String())
	if err != nil {
		fmt.Println("Error writing to numbers.txt:", err)
		return
	}

	_, err = lettersFile.WriteString(lettersData.String())
	if err != nil {
		fmt.Println("Error writing to letters.txt:", err)
		return
	}

	// Write the total sum to numbers.txt
	_, err = numbersFile.WriteString(fmt.Sprintf("Total sum: %d\n", totalSum))
	if err != nil {
		fmt.Println("Error writing total sum to numbers.txt:", err)
		return
	}

	fmt.Println("Files written successfully.")
}
