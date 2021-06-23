package main

import (
	"fmt"
	"os"
)

func main() {
	name, course := "Anika", "Go"
	new_name := os.Getenv("USER")
	fmt.Println("\nHi", name, " and ", new_name, "you're currently watching", course)
	//fmt.Println("Name is set to", course, "and type is", reflect.TypeOf(course) )
	// fmt.Println("module is set to", module, "and type is", reflect.TypeOf(module))

	changeCourse(course) // Make copy of var and not pass actual var
	fmt.Println("\n The value of couse is ", course)
}

func changeCourse(course string) string {
	course = "Docker"
	fmt.Println("\n Trying to change course to ", course)
	return course

}
