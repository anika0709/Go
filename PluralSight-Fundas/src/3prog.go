package main

import (
	"fmt"
	"reflect"
)

func main(){
	name, course, module := "Anika", "Go", 3.2

	fmt.Println("Name is set to", name, "and type is", reflect.TypeOf(name) )
	fmt.Println("Name is set to", course, "and type is", reflect.TypeOf(course) )
	fmt.Println("module is set to", module, "and type is", reflect.TypeOf(module))
	a := 3.00000
	b := 10
	fmt.Println("Mem addr of module is set to", &module)
	ptr := &module
	fmt.Println("module is set to", *ptr, "and mem addr is", ptr )
	fmt.Println("Type of a and b is", reflect.TypeOf(a), reflect.TypeOf(b))
	c := int(a) + b
	fmt.Println("Type of c", reflect.TypeOf(c))
}