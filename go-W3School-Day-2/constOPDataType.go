// Go Constants,
package main

import "fmt"

const A = 1
const PI = 3.14

/*
Constant Rules
    Constant names follow the same naming rules as variables
    Constant names are usually written in uppercase letters (for easy identification and differentiation from variables)
    Constants can be declared both inside and outside of a function
*/

//multi constant
const (
	A1 int = 1
	B1     = 3.14
	C1     = "Hi!"
)

/*
Go has three functions to output text:

    Print()
    Println()
    Printf()
*/

func main() {
	fmt.Println(A)
	fmt.Println(PI)
	fmt.Println(A1)
	fmt.Println(B1)
	fmt.Println(C1)

	var i, j string = "Hello", "World"

	fmt.Print(i)
	fmt.Print(j, "\n")
	fmt.Println(i)
	fmt.Println(j)
	fmt.Print(i, "\n", j)
	fmt.Print("\n")
	fmt.Print(i, " ", j)

	var i2, j2 = 10, 20

	fmt.Print("\n", i2, j2)

	fmt.Print("\n")
	fmt.Println(i, j)

	var i3 string = "Hello"
	var j3 int = 15

	fmt.Printf("i3 has value: %v and type: %T\n", i3, i3)
	fmt.Printf("j3 has value: %v and type: %T\n", j3, j3)

	/*
		Verb 	Description
		%v 		Prints the value in the default format
		%#v 	Prints the value in Go-syntax format
		%T 		Prints the type of the value
		%% 		Prints the % sign
	*/

	var k = 15.5
	var txt = "Hello World!"

	fmt.Printf("%v\n", k)
	fmt.Printf("%#v\n", k)
	fmt.Printf("%v%%\n", k)
	fmt.Printf("%T\n", k)

	fmt.Printf("%v\n", txt)
	fmt.Printf("%#v\n", txt)
	fmt.Printf("%v%%\n", txt)
	fmt.Printf("%T\n", txt)

	/*
		Verb 	Description
		%b 		Base 2
		%d 		Base 10
		%+d 	Base 10 and always show sign
		%o 		Base 8
		%O 		Base 8, with leading 0o
		%x 		Base 16, lowercase
		%X 		Base 16, uppercase
		%#x 	Base 16, with leading 0x
		%4d 	Pad with spaces (width 4, right justified)
		%-4d 	Pad with spaces (width 4, left justified)
		%04d 	Pad with zeroes (width 4)
	*/

	var i6 = 15

	fmt.Printf("Base 2,%b\n", i6)
	fmt.Printf("Base 10,%d\n", i6)
	fmt.Printf("Base 10 and always show sign,%+d\n", i6)
	fmt.Printf("Base 8,%o\n", i6)
	fmt.Printf("Base 8, with leading 0o,%O\n", i6)
	fmt.Printf("Base 16, lowercase,%x\n", i6)
	fmt.Printf("Base 16, uppercase,%X\n", i6)
	fmt.Printf("Base 16, with leading 0x,%#x\n", i6)
	fmt.Printf("Pad with spaces (width 4, right justified),%4d\n", i6)
	fmt.Printf("Pad with spaces (width 4, left justified),%-4d\n", i6)
	fmt.Printf("Pad with zeroes (width 4),%04d\n", i6)

	//String Formatting Verbs teke soro hbe
	//https://www.w3schools.com/go/go_formatting_verbs.php

}
