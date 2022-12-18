package main

// Import multiple packages
// You could use an alias like f "fmt"

import(
	"bufio"
	f "fmt"	
	"log"
	"os"
)


// Create alias to long function names

var pl = f.Println

func main(){
	// Prints text and a newline
	// List package name followed by a period and the function name
	pl("Hello Go")

	// Get user input (To run this in the terminal go run hellogo.go)
	pl("Where do you live ?")
	// Setup buffered reader that gets text from the keyboard
	reader := bufio.NewReader(os.Stdin)
	// Copy text up to the newline
	// The blank identifier _ will get err and ignore it (Bad Practice)
	// name, _ := reader.ReadString('\n')
	// It is better to handle it
	place, err := reader.ReadString('\n')
	if err == nil {
		pl("Welcome To", place)
	} else {
		// Log this error
		log.Fatal(err)
	}
}


// go run gotraining.go