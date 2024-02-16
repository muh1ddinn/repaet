
func main() {
	// Open or create the file for writing
	file, err := os.OpenFile("person.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("error while opening file: ", err)
		return
	}
	defer file.Close()

	// Prompt the user to enter the count of persons
	var n int
	fmt.Println("Enter the count of persons: ")
	fmt.Scan(&n)

	// Write person data to the file
	for i := 1; i <= n; i++ {
		var name, gmail, job string

		fmt.Printf("Enter the %v person's name: ", i)
		fmt.Scanln(&name)
		fmt.Printf("Enter the %v person's gmail: ", i)
		fmt.Scanln(&gmail)
		fmt.Printf("Enter the %v person's job: ", i)
		fmt.Scanln(&job)

		// Write the data to the file
		_, err := file.WriteString(fmt.Sprintf("Person %d: Name: %s, Gmail: %s, Job: %s\n", i, name, gmail, job))
		if err != nil {
			fmt.Println("Error writing to file: ", err)
			return
		}
	}

	fmt.Println("Data has been written to person.txt")
}