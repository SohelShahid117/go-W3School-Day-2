// Go Constants,output,data types
package main

import "fmt"

const A = 1
const PI = 3.14

/*
Constant Rules
    Constant names-->follow the same naming rules as variables
    Constant names-->are usually written in uppercase letters (for easy identification and differentiation from variables)
    Constants-->can be declared both inside and outside of a function
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

	fmt.Printf("value in the default format,%v\n", k)
	fmt.Printf("value in Go-syntax format,%#v\n", k)
	fmt.Printf(",Prints the % sign,%v%%\n", k)
	fmt.Printf("type of the value,%T\n", k)

	fmt.Printf("value in the default format,%v\n", txt)
	fmt.Printf("value in Go-syntax format,%#v\n", txt)
	fmt.Printf(",Prints the % sign,%v%%\n", txt)
	fmt.Printf("type of the value,%T\n", txt)

	// fmt.Printf("%v\n", txt)
	// fmt.Printf("%#v\n", txt)
	// fmt.Printf("%v%%\n", txt)
	// fmt.Printf("%T\n", txt)

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

	/*
		Verb 	Description
		%s 		Prints the value as plain string
		%q 		Prints the value as a double-quoted string
		%8s 	Prints the value as plain string (width 8, right justified)
		%-8s 	Prints the value as plain string (width 8, left justified)
		%x 		Prints the value as hex dump of byte values
		% x 	Prints the value as hex dump with spaces
	*/

	var txt2 = "Hello"

	fmt.Printf("plain string,%s\n", txt2)
	fmt.Printf("double-quoted string,%q\n", txt2)
	fmt.Printf("plain string (width 8, right justified),%8s\n", txt2)
	fmt.Printf("plain string (width 8, left justified),%-8s\n", txt2)
	fmt.Printf("hex dump of byte values,%x\n", txt2)
	fmt.Printf("hex dump with spaces,% x\n", txt2)

	//%t 	Value of the boolean operator in true or false format (same as using %v)

	var x5 = true
	var x6 = false

	fmt.Printf("%t\n", x5)
	fmt.Printf("%t\n", x6)
	fmt.Printf("%T\n", x6)

	/*
		Verb 	Description
		%e 		Scientific notation with 'e' as exponent
		%f 		Decimal point, no exponent
		%.2f 	Default width, precision 2
		%6.2f 	Width 6, precision 2
		%g 		Exponent as needed, only necessary digits
	*/

	var pii = 3.141

	fmt.Printf("%e\n", pii)
	fmt.Printf("%f\n", pii)
	fmt.Printf("%.2f\n", pii)
	fmt.Printf("%6.2f\n", pii)
	fmt.Printf("%g\n", pii)

	/*
			Go has three basic data types:

		    bool: represents a boolean value and is either true or false
		    Numeric: represents integer types, floating point values, and complex types
		    string: represents a string value

	*/

	var a5 bool = true    // Boolean
	var b5 int = 5        // Integer
	var c5 float32 = 3.14 // Floating point number
	var d5 string = "Hi!" // String

	fmt.Println("Boolean: ", a5)
	fmt.Println("Integer: ", b5)
	fmt.Println("Float:   ", c5)
	fmt.Println("String:  ", d5)

	var b11 bool = true // typed declaration with initial value
	var b21 = true      // untyped declaration with initial value
	var b31 bool        // typed declaration without initial value
	b41 := true         // untyped declaration with initial value

	fmt.Println(b11)
	fmt.Println(b21)
	fmt.Println(b31)
	fmt.Println(b41)

	/*
			The integer data type has two categories:

		    Signed integers - can store both positive and negative values
		    Unsigned integers - can only store non-negative values
	*/

	var x int = 500
	var y int = -4500
	fmt.Printf("Type: %T, value: %v\n", x, x)
	fmt.Printf("Type: %T, value: %v\n", y, y)

	var x1 uint = 500
	var y1 uint = 4500
	fmt.Printf("Type: %T, value: %v\n", x1, x1)
	fmt.Printf("Type: %T, value: %v\n", y1, y1)

	//Unsigned integers, declared with one of the uint keywords, can only store non-negative values:

	var x3 float32 = 123.78
	var y3 float64 = 3.4e+38
	fmt.Printf("Type: %T, value: %v\n", x3, x3)
	fmt.Printf("Type: %T, value: %v\n", y3, y3)

	var x4 float64 = 3.4e+39
	fmt.Println(x4)

	//The string data type is used to store a sequence of characters (text). String values must be surrounded by double quotes:

	var txt12 string = "Hello!"
	var txt22 string
	txt32 := "World 1"

	fmt.Printf("Type: %T, value: %v\n", txt12, txt12)
	fmt.Printf("Type: %T, value: %v\n", txt22, txt22)
	fmt.Printf("Type: %T, value: %v\n", txt32, txt32)

}
