package main

import (
	"fmt"
	"reflect"
	//	"time"
)

func main() {
	a := 2
	b := &a

	fmt.Println(a, b)
	fmt.Println(reflect.TypeOf(a), reflect.TypeOf(b))

	c := "Anika"
	d := &c
	h := *d

	fmt.Println(c, d, "h=", h)
	fmt.Println(reflect.TypeOf(c), reflect.TypeOf(d))

	e := []int{1, 2, 3, 4}
	f := &e
	fmt.Println(e, reflect.TypeOf(e))
	fmt.Println("Address of slice", f)

}
