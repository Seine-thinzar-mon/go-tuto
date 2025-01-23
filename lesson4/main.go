package main

import "fmt"

func main() {
	age := 23
	name := "Seine"

	// Print, don't break a new line. use \n to break a new line
	fmt.Print("Hello, ")
	fmt.Print("world! \n")
	fmt.Print("This is new line \n")

	// Println, break a new line
	fmt.Println("Hay yo!")
	fmt.Println("Bye!")
	fmt.Println("My name is", name, "and I am", age, "years old.")

	//Printf(formatted string), %v refers format specifier, doesn't add new line automatically
	fmt.Printf("My name is %v and I am %v years old.\n", name, age) // %v is for default format
	fmt.Printf("My name is %q.\n", name)                            // %q is to place quotes around the variable, but it can be used for string type variable
	fmt.Printf("Age is of type: %T\n", age)                         // %T is to get the type of variable
	fmt.Printf("Float number is %f\n", 23.773477)                   // %f is to output the float number
	fmt.Printf("Fix decimal point %0.2f\n", 746.398372)

	// Sprintf(save formatted string)
	var str = fmt.Sprintf("My name is %v and I am %v years old.\n", name, age)
	fmt.Println("Saved string is:", str)
}
