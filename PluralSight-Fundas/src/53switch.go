package main

import (
	"fmt"
)

func main() {
	switch "docker" {
	case "linux":
		fmt.Println("Try docker")
	case "docker":
		fmt.Println("Found it")
	case "DOCKER":
		fmt.Println("Try docker")
	default:
		fmt.Println("Please come later")
	}
}
