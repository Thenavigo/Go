package main

import(
	f "fmt"
)

var pl = f.Println

func main() {
	// ----- SLICES -----
	// Slices are like arrays but they can grow
	// var name []dataType
	// Create a slice with make
	sl1 := make([]string, 6)


	// Assign values by index
	sl1[0] = "Society"
	sl1[1] = "on"
	sl1[2] = "off"
	sl1[3] = "the"
	sl1[4] = "Universe"


	// Size of slice
	pl("Slice size :", len(sl1))


	// Cycle with for
	for i := 0; i < len(sl1); i++ {
		pl(sl1[i])
	}


	// Cycle with range
	for _, x := range sl1 {
		pl(x)
	}


	// Create a slice literal
	sl2 := []int{12, 21, 1974}
	pl(sl2)


	// A slice points at an array and you can create a slice
	// of an array (A slice is a view of an underlying array)
	// You can have multiple slices point to the same array

	sArr := [5]int{1, 2, 3, 4, 5}
	// Start at 0 index up to but not including the 2nd index
	sl3 := sArr[:3]
	// Get slice from beginning
	pl(sl3)

	sl4 := sArr[0:2]
	pl(sl4)

	sArr2 := [5]int{1, 2, 3, 4, 5}
	sl5 := sArr2[0:2]

	// If you change the array the slice also changes
	sArr2[0] = 10
	pl("sl5 :", sl5)

	// Append a value to a slice (Also overwrites array)
	sl5 = append(sl5, 12)
	pl("sl5 :", sl5)
	pl("sArr2 :", sArr2)


	// Printing empty slices will return nils which show
	// as empty slices
	sl6 := make([]string, 6)
	pl("sl6 :", sl6)
	pl("sl6[0] :", sl6[0])

}