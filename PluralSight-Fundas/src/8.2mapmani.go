package main

import (
	"fmt"
)

func main() {
	// Add multiple values
	mm := make(map[string]int)
	mm["route"] = 80
	fmt.Println("mm map is", mm)

	mymap := map[string]int{"route": 1,
		"route1": 2,
		"route3": 3}

	fmt.Println(mymap)
	mm["route"] = mymap["route"]
	fmt.Println("mm[route]=", mm["route"])

	//length of Map
	fmt.Println("length of Map", len(mymap))

	//delete
	fmt.Println("Deleting mm[routing")
	delete(mm, "route")
	fmt.Println("Now mm[route]=", mm["route"])

	commits := map[string]int{
		"rsc": 3711,
		"r":   2138,
		"gri": 1908,
		"adg": 912,
	}

	// Update the item
	commits["r"] = 46457
	fmt.Println("commits[r] = ", commits["r"])

	val, found := commits["r"] // found is true, as the value exist for key
	fmt.Println("vallue and Found =", val, found)

	_, found = commits["rew"] // found is false, as the value doesn't exist
	fmt.Println("Found =", found)

	for key, value := range commits {
		fmt.Println(key, value)

	}

	//sum
	int_map := map[int]int{0: 1,
		1: 100, 2: 454,
		3: 8786545678}
	sum := 0
	for _, num := range int_map {
		sum = sum + num
	}
	fmt.Println("Sum is", sum)
}
