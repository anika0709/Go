package main

import (
	"fmt"
)

func main() {
	//array declare buts not initialize

	var c [4]int
	fmt.Println("c=", c)
	c[3] = 100
	fmt.Println("c[3] and c are", c[3], c)

	//declare and initialize an array
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println(b[2])

	// make slice using make command but not initialize
	myslice := make([]string, 3, 5)
	fmt.Println("Length of myslice", len(myslice), "\nCapacity of myslice ", cap(myslice))
	fmt.Println(myslice)

	// need to check
	//myslice = "Item1", "Item2", "Item3"

	//assign value to slice
	// myslice[0], myslice[1], myslice[2] = "Item1", "Item2", "Item3"
	// fmt.Println(myslice)

	// // Declare and initialize
	// secSlice := []string{"Chec1", "Chec2", "Chec3", "Chec4", "Chec5"}
	// fmt.Println("Length of secSlice", len(secSlice), "\nCapacity of secSlice ", cap(secSlice))
	// fmt.Println(secSlice)

	// secSlice[2] = "Notcheck3"
	// fmt.Println(secSlice)

	// newSlice := secSlice[1:4]
	// fmt.Println(newSlice)

}
