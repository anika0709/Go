package main

import (
	"fmt"
)

func main() {

	// for i := 10; i < 100; i = i + 10 {
	// 	fmt.Println("Value of i is ", i)
	// Brake:
	// 	for j := 0; j < 10; j++ {
	// 		fmt.Println("Value of J is", j)
	// 		if j == 5 {
	// 			fmt.Println("Value of J is", j)
	// 			break
	// 		}
	// 		for k := 0; k < 10; k++ {
	// 			fmt.Println("k is now", k)
	// 			if k == 5 {
	// 				fmt.Println("k is now", k)
	// 				break Brake
	// 			}
	// 		}
	// 	}
	// }

	// Continue Example
	// for i := 10; i < 100; i = i + 10 {
	// 	fmt.Println("Value of i is ", i)
	// Cont:
	// 	for j := 0; j < 10; j++ {
	// 		fmt.Println("Value of J is", j)
	// 		if j == 5 {
	// 			fmt.Println("Value of J is", j)
	// 			break
	// 		}
	// 		for k := 0; k < 10; k++ {
	// 			fmt.Println("k is now", k)
	// 			if k == 5 {
	// 				fmt.Println("k is now", k)
	// 				continue Cont
	// 			}
	// 		}
	// 	}
	// }
	// Print Even Numbers between 0-20

	for i := 0; i <= 20; i++ {
		if i%2 == 1 {
			continue
		}
		fmt.Println(i)
	}
}
