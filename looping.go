package main

import (
	f "fmt"
)

var pl = f.Println

func main() {
	// for initialization; condition; postStatement {BODY}
	// INCREMENT
	for x:= 1; x <= 5; x++ {
		pl(x)
	}

	// DECREMENT
	for x:= 5; x >= 1; x-- {
		pl(x)
	}

	// WHILE LOOP
	fx := 0
	for fx < 5 {
		pl(fx)
		fx++
	}

	aNums := []int{1, 2, 3}
	for _, num := range aNums {
		pl(num)
	}

	xVal := 1
	for true{
		if xVal == 5{
			break
		}
		pl(xVal)
		xVal++
	}
}