package main

import (
	"fmt"
	"strings"
)

func main() {

	module := "function basics"
	author := "Anika Rathi"

	fmt.Println(module)
	fmt.Println(author)

	auth, mod := converter(module, author)
	fmt.Println(auth, "\n", mod)
	//fmt.Println(converter(module, author))
}

func converter(mod, auth string) (s1, s2 string) {
	mod = strings.Title(mod)
	auth = strings.ToUpper(auth)

	return mod, auth
}
