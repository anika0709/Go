package main

import (
	"fmt"
)

func main() {
	name, course := "Anika", "Go"
	fmt.Println("\nHi", name, "you're currently watching", course)
	//fmt.Println("Name is set to", course, "and type is", reflect.TypeOf(course) )
	// fmt.Println("module is set to", module, "and type is", reflect.TypeOf(module))

	changeCourse(&course) // pass as reference - pass location mem
	fmt.Println("\n The value of couse is ", course)
}

//Try without declaring in f(x)
func changeCourse(courses *string) string {
	*courses = "Docker"
	fmt.Println("\n Trying to change course to ", *courses)
	return *courses

}
