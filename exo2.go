package main

import (
	f "fmt"
	  "bufio"
	  "log"
	  "os"
	  "strings"
	  "strconv"
)

// Create alias to long function names
var pl = f.Println

func main() {
	// Setup buffered reader that gets text from the keyboard
	reader := bufio.NewReader(os.Stdin)


	f.Print("Enter Number 1 : ")
	num1, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	// Trims whitespace from guess
	num1 = strings.TrimSpace(num1)
	iNum1, err := strconv.Atoi(num1)
	if err != nil {
		log.Fatal(err)
	}


	f.Print("Enter Number 2 : ")
	num2, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}


	// Trims whitespace from guess
	num2 = strings.TrimSpace(num2)
	iNum2, err := strconv.Atoi(num2)
	if err != nil {
		log.Fatal(err)
	}

	// Use Printf to create formatted string
	f.Printf("%s + %s = %d\n", num1, num2, (iNum1 + iNum2))

}