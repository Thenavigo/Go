package main

import (
	f "fmt"	
)

var pl = f.Println

func main() {
	// If Conditional
	// Conditional Operators : > < >= <= == !=
	// Logical Operators : && || !

	iAge := 8
	if (iAge >= 1) && (iAge <= 18){
		pl("Important Birthday")
	} else if (iAge == 21) || (iAge == 50){
		pl("Important Birthday")
	} else if iAge >= 65 {
		pl("Important Birthday")
	} else {
		pl("Not an Important Birthday")
	}


	// ! turns bools into their opposite value
	pl("!true = ", !true)


	// ----- FORMATTED PRINT -----
	// Go has its own version of C's printf
	// %d : Integer
	// %c : Character
	// %f : Float
	// %t : Boolean
	// %s : String
	// %o : Base 8
	// %x : Base 16
	// %v : Guesses based on data type
	// %T : Type of supplied value

	f.Printf("%s %d %c %f %t %o %x\n", "Stuff", 1, 'A', 3.14, true, 1,1)

	// Float formatting
	f.Printf("%9f\n", 3.14)      // Width 9
	f.Printf("%.2f\n", 3.141592) // Decimal precision 2
	f.Printf("%9.f\n", 3.141592) // Width 9 no precision

	// Sprintf returns a formatted string instead of printing
	sp1 := f.Sprintf("%9.f\n", 3.141592)
	pl(sp1)
}