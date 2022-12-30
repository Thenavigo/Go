package main

import(
	f "fmt"
)

var pl = f.Println

// func say hello
func sayHello() {
	pl("Hello")
}


// func getSum
func getSum(x int, y int) int {
	return x + y
}


// func getQuotient
func getQuotient(x float64, y float64) (ans float64, err error) {
	if y == 0 {
		return 0, f.Errorf("You can't divide by zero")
	} else {
		return x / y, nil
	}
}


// func getTwo
func getTwo(x int) (int, int) {
	return x + 1, x + 2	
}


// func getSum2
func getSum2(nums ...int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}


// func getArraySum
func getArraySum(arr []int) int {
	sum := 0
	for _, num := range arr {
		sum += num
	}
	return sum
}


// func changeVal
func changeVal(f3 int) int {
	f3 += 1
	return f3
}


func main() {
	// print message
	sayHello()

	// print sum
	pl(getSum(5, 4))

	// print ans
	ans, err := getQuotient(5, 4)
		if err == nil {
			f.Printf("%f / %f = %f\n", 5.0, 4.0, ans)
		} else {
			pl(err)
		}

	// print sum
	f1, f2 := getTwo(5)
	f.Printf("%d\n  %d\n", f1, f2)

	
	// print getSum2
	pl("Unknow Sum :", getSum2(1, 2, 3, 4))

	
	// print getArraySum
	vArr := []int{1,2,3,4}
	pl("Array Sum :", getArraySum(vArr))


	// print changeVal
	f3 := 5
	pl("f3 before func :", f3)
	changeVal(f3)
	pl("f3 after func :", f3)
}