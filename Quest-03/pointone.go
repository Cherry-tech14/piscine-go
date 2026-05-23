package main

import (
	"fmt"
)

func PointOne(n *int) {

}
func main() {
	n := 1
	PointOne(&n)
	fmt.Println(n)
}
