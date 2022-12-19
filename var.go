package main 

import(
	f "fmt"
	"reflect"
	"strconv"
)

var pl= f.Println

func main() {
	// int, float64, bool, string, rune
	// Default type 0, 0.0, false, ""
	pl(reflect.TypeOf(25))
	pl(reflect.TypeOf(3.14))
	pl(reflect.TypeOf(true))
	pl(reflect.TypeOf("Hello"))
	pl(reflect.TypeOf('✨'))

	cV1 := 1.5
	cV2 := int(cV1)
	pl(cV2)

	cV3 := "1000000"
	cV4, err := strconv.Atoi(cV3)
	pl(cV4, err, reflect.TypeOf(cV4))


	cV5 := 1000000
	cV6 := strconv.Itoa(cV5)
	pl(cV6)

	cV7 := "3.14"
	if cV8, err := strconv.ParseFloat(cV7, 64); err == 
	nil {
		pl(cV8)
	}

	cV9 := f.Sprintf("%f", 3.14)
	pl(cV9)
}