package main

import (
	"fmt"
)

func main() {

	courseInProg := []string{"Docker", "Go", "CKAD", "K8s"}
	courseComplete := []string{"AWS", "CKA", "K8s"}

	for _, i := range courseInProg {
		//fmt.Println(i)
		for _, j := range courseComplete {
			if i == j {
				fmt.Println("Hey found a clash", i)
			}
		}
	}
}
