package main

import (
	"fmt"
)

func main() {

	bestFinish := bestLeagueFinishes(15, 23, 55, 76, 54, 64, 10, 2) //using a f(x) to declare a variable
	fmt.Println(bestFinish)
}

func bestLeagueFinishes(finishes ...int) int {

	best := finishes[0] // first item in the list
	for _, i := range finishes {
		if i < best {
			best = i
		}
	}
	return best
}
