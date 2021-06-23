package main

import (
	"fmt"
)

func main() {
   
   var x int

   x = 30

   if y := x%3; y == 0 {
   	fmt.Printf ("x is divisble by 3. FIZZZ")
   } else if y := x%5; y == 0 {
   	fmt.Printf ("x is divisble by 5. BUZZ")
   	} else if y := x%5; y == 0 && y := x%3; y == 0 {
   	fmt.Printf ("X is Fizz Buzz")
   }
}