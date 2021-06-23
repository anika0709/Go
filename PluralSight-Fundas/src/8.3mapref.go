package main

import (
	"fmt"
)

func main() {
	//var m map[int]int
	m := make(map[int]int)
	m = map[int]int{0: 1,
		1: 100, 2: 454,
		3: 8786545678,
	}

	// Initial State
	for key, val := range m {
		fmt.Println(key, "=>", val)
	}

	fmt.Println("-----------------------")

	sumofmap(m)

	// After Passing to the function as a pointer
	for key, val := range m {
		fmt.Println(key, "=>", val)
	}
}

func sumofmap(m map[int]int) {

	m[1] = 45
	m[2] = 65
	m[3] = 55

}
