package main

import (
	"fmt"
)

func main() {

	// apend()
	firstSlice := make([]int, 2, 4)

	for i := 0; i <= 17; i++ {
		firstSlice = append(firstSlice, i)
		fmt.Println("Length and Capacity is", len(firstSlice), "and", cap(firstSlice))
	}

	//Appending Slice to Slice

	mySlice := []int{1, 2, 3, 4, 5}
	fmt.Println("mySlice is", mySlice)

	for j, i := range mySlice {
		fmt.Sprintf("The value index %d is %d", j, i)
	}

	newSlice := []int{10, 20, 30}
	mySlice = append(mySlice, newSlice...)
	fmt.Println("MY new slice is", mySlice)
}
