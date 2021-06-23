package main

import (
	"fmt"
)



func main() {
	x := 1

	if x < 2 {
		fmt.Printf ("x is less than 2\n")
	} else if x > 2 {
		fmt.Printf ("x is greater than 2\n")
	} else {
		fmt.Printf ("x is equal to 2\n")
	}
	for x := 0; x < 3; x++ {
		fmt.Println(x)
	}
    
    y := 0
	for y < 3 {
		fmt.Println (y)
		y++
	}

}

