package main

import (
	"fmt"
)

func main() {

	if firstRank, secondRank := 83957, 54614; firstRank < secondRank {

		fmt.Println("\n First Course is better" +
			" than second course")
	} else if firstRank > secondRank {
		fmt.Println("\n Your first course is better")
	} else {
		fmt.Println("both are either same or there is some issue")
	}
}
