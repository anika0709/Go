package main

import (
	"fmt"
	"reflect"
)

func main() {

	var mm map[string]int //not initialize, can read but ---
	// if you try to write to empty map, it will lead to Panic error

	fmt.Println("Initial value of mm map = ", mm)

	// mm["route"] = 80 can't write to nill map

	mm = make(map[string]int) // allocates and initializes a hash map

	mm["route"] = 80

	fmt.Println("mm[route]=", mm["route"])

	hw := mm["route"]
	//test := mm["route1"]
	//MyMap[1,2,3] = {1,2,3}
	fmt.Println("New value of mm map", mm)
	fmt.Println(mm["route"])
	fmt.Println("mm[real] is not assigned so it's value is ", mm["real"])
	fmt.Println("Value of HW is mm[route] and its value is", hw, &hw)
	fmt.Println("Type of HW is mm[route] and its mem address is", reflect.TypeOf(hw), reflect.TypeOf(&hw))

	//Another way to declare map
	commits := map[string]int{
		"rsc": 3711,
		"r":   2138,
		"gri": 1908,
		"adg": 912,
	}

	for key, value := range commits {
		fmt.Println(key, value)

	}

}
