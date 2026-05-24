// package main

// import (
// 	"fmt"
	
// )
// func main() {
// 	fmt.Println("1")
// }


// checknumber

// Write a function that takes a string as an argument and returns true if the string contains any number, otherwise return false.

// package main

// import (
// 	"fmt"
// )
// func CheckNumber(arg string) bool {
// 	for _, ch := range arg {
// 		if ch >= '0' && ch <= '9' {
// 			return true
// 		}
		
// 	}
// 	return false
// }
// func main() {
// 	fmt.Println(CheckNumber("100"))
// 	fmt.Println(CheckNumber("Hello1"))
// }


// countAlpha
// Write a function CountAlpha() that takes a string as an argument and returns the number of alphabetic characters in the string.

// package main

// import (
// 	"fmt"
// )
// func CountAlpha(s string) int {
// 	count := 0
// 	for _, ch := range s {
// 		if (ch >= 'a' && ch <= 'z') || (ch <= 'A' && ch <= 'Z') {
// 			count++
// 		}
// 	}
// 	return count

// }
// func main() {
// 	fmt.Println(CountAlpha("Hello world"))
// 	fmt.Println(CountAlpha("H e l l o"))
// 	fmt.Println(CountAlpha("H1e2l3l4o"))
// }

// countcharacter
// write a function that takes a string and a character as arguments and returns the number of times the character appears in the string.

    // if the character is not in the string return 0
    // if the string is empty return 0

// package main

// import (
// 	"fmt"
// )
// func CountChar(str string, c rune) int {
// 	count := 0
// 	for _, ch := range str {
// 		if ch == c {
// 			count++ 
// 		}
// 	}
// 	return count
// }

// func main() {
// 	fmt.Println(CountChar("Hello World", 'l'))
// 	fmt.Println(CountChar("5  balloons", 5))
// 	fmt.Println(CountChar("   ", ' '))
// 	fmt.Println(CountChar("The 7 deadly sins", '7'))
// }

// printif
// Write a function that takes a string as an argument and returns the letter G followed by a newline \n if the argument length is more or equal than 3, otherwise returns Invalid Input followed by a newline \n.

    // If it's an empty string return G followed by a newline \n.

	
	
// 	package main

// 	import (
// 		"fmt"
// 	)
// 	func PrintIf(str string) string {
// 		if len(str) >= 3 || len(str) == 0 {
// 			return "G\n"
// 		} else {
// 			return "Invalid input"
// 		}
// 	}

// func main() {
// 	fmt.Print(PrintIf("abcdefz"))
// 	fmt.Print(PrintIf("abc"))
// 	fmt.Print(PrintIf(""))
// 	fmt.Print(PrintIf("14"))
// }


// printifnot
// Write a function that takes a string and returns:

    // "G\n" if the string's length is less than 3 (including empty string).

    // "Invalid Input\n" otherwise.

	// package main

	// import (
	// 	"fmt"
	// )
	// func PrintIfNot(str string) string {
	// 	if len(str) < 3 {
	// 		return "G"
	// 	}else {
	// 		return "invalid input"
	// 	}
	// }
	// func main() {
	// 	fmt.Println(PrintIfNot("abcdefz"))
	// 	fmt.Println(PrintIfNot("abc"))
	// 	fmt.Println(PrintIfNot(""))
	// 	fmt.Println(PrintIfNot("14"))
	// }

	// rectperimeter
	// Write a function that takes two int's as arguments, representing the length of width and height of a rectangle and returning the perimeter of the rectangle.

    // If one of the arguments is negative it should return -1.

// 	package main

// 	import (
// 		"fmt"
// 	)
// 	func RectPerimeter(w, h int) int {
// 		if w < 0 || h < 0 {
// 			return -1
// 		}
// 		return 2* (w + h)

// 	}
// 	func main() {
// 		fmt.Println(RectPerimeter(10, 2))
// 		fmt.Println(RectPerimeter(434343, 898989))
// 		fmt.Println(RectPerimeter(10, -2))
// }
	

// retain firsthalf
// Write a function called RetainFirstHalf() that takes a string as an argument and returns the first half of this string.

//     If the length of the string is odd, round it down.
//     If the string is empty, return an empty string.
//     If the string length is equal to one, return the string.

package main

import (
	"fmt"
)
func RetainFirstHalf(str string) string {
	if len(str) == 1 {
		return str
	}
	half := len(str) / 2

	return str[:half]
}
func main() {
	fmt.Println(RetainFirstHalf("This is the 1st halfThis is the 2nd half"))
	fmt.Println(RetainFirstHalf("A"))
	fmt.Println(RetainFirstHalf(""))
	fmt.Println(RetainFirstHalf("Hello World"))
}
	










