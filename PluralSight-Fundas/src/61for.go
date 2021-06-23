package main

import (
	"fmt"
	//	"time"
)

func main() {

	i := 0

	for timer := 10; timer >= 0; timer-- {
		if timer == 0 {
			fmt.Println("Boom!")
			break
		}
		fmt.Println(timer)
		//time.Sleep(1 * time.Second)
	}
	for tester := 0; tester <= 10; tester++ {
		fmt.Println(tester)
	}
	// Infinite look
	// for {
	// 	i := 0
	// 	i++

	// 	if i == 5 {
	// 		fmt.Println("Reached the limit of ", i)
	// 		break
	// 	}
	// }
	//Boolean exp
	for i < 10 {
		fmt.Println("Testing i", i)
		i++

	}
	// loop over range
	slice := []string{"Test1", "Value2", "Item3"}
	for _, j := range slice {
		fmt.Println(j)
	}
}
