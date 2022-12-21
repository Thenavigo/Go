package main 

import (
  f "fmt"
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

var pl = f.Println

func main() {
	// Receive customer data (Their age)
	reader := bufio.NewReader(os.Stdin)
	pl("What is your age : ") // What is your age : 
	age, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	// Google how to trim whitespace from input
	age = strings.TrimSpace(age)
	iAge, err := strconv.Atoi(age)
	if err != nil {
		log.Fatal(err)
	}
	
	// Age < 5 Too young for school

	if iAge < 5 {
		pl("Too young for school")
	} else if iAge == 5 {
		pl("Go to kindergarden") // Age == 5 Go to kindergarden
	} else if (iAge > 5) && (iAge <= 17) { // Age > 5 or Age <= 17 Go to grade GRADE
	
		grade := iAge - 5
		f.Printf("Go to grade %d\n", grade)
	} else {
		pl("Go to college") // Default Go to college
	}
}