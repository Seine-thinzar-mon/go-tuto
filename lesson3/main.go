package main

import "fmt"

func main() {
	// string
	var strOne string = "Hello" // declare string
	var strTwo = "Hola"         // automatically refer the string type
	var strThree string
	strFour := "Konnichiwa" // shortern, automatically refer the string type

	// strThree = 23 // cannot assign integer in string type variable
	strThree = "Mingalapr"

	fmt.Println("string =>", strOne, strTwo, strThree, strFour)

	//int
	var valOne int = 10
	var valTwo = 20
	var valThree int8     // can assign only between its bit range
	valFour := 40         // automatically refer the int type
	var valFive uint = 20 // unassigned int type, cannot assign minus value

	// valThree = 267 // cannot assign beacause out of 8 bit range
	valThree = 30

	fmt.Println("integer =>", valOne, valTwo, valThree, valFour, valFive)

	// float
	var deciOne float32 = 34.33 // can assign only between its bit range
	var deciTwo float64 = 47377.78
	deciThree := 837373.03 // automatically referred to float64

	fmt.Println("float =>", deciOne, deciTwo, deciThree)
}
