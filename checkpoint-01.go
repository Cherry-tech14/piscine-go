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

package main

import (
	"fmt"
)
func CountChar(str string, c rune) int {
	count := 0
	for _, ch := range str {
		if ch == c {
			count++ 
		}
	}
	return count
}

func main() {
	fmt.Println(CountChar("Hello World", 'l'))
	fmt.Println(CountChar("5  balloons", 5))
	fmt.Println(CountChar("   ", ' '))
	fmt.Println(CountChar("The 7 deadly sins", '7'))
}
