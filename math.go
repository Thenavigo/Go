package main

import (
	f "fmt"
	  "math/rand"
	  "time"
	  "math"
)

var pl = f.Println 

func main() {
	pl("5 + 4 = ", 5+4)
	pl("5 - 4 = ", 5-4)
	pl("5 * 4 = ", 5*4)
	pl("5 / 4 = ", 5/4)
	pl("5 % 4 = ", 5%4)
	mInt := 1
	mInt++
	pl(mInt)

	seedSecs := time.Now().Unix()
	rand.Seed(seedSecs)
	randNum := rand.Intn(50) + 1
	pl("Random :", randNum)


	pl("Abs(-10) =", math.Abs(-10))
	pl("Pow(4, 2) =", math.Pow(4,2))
	pl("Sqrt(16) =", math.Sqrt(16))
	pl("Cbrt(8) =", math.Cbrt(8))
	pl("Ceil(4.4) =", math.Ceil(4.4))
	pl("Floor(4.4) =", math.Floor(4.4))
	pl("Round(4.4) =", math.Round(4.4))
	pl("Log2(8) =", math.Log2(8))
	pl("Log10(100) =", math.Log10(100))
	pl("Log(7.389) =", math.Log(7.389))
	pl("Max(5,4) =", math.Max(5,4))
	pl("Min(5,4) =", math.Min(5,4))

	r90 := 90 * math.Pi / 180
	d90 := r90 * (180 / math.Pi)
	f.Printf("%f radians = %f degree\n", r90, d90)
	pl("Sin(90) =", math.Sin(r90))
}