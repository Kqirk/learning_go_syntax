// Variables are explicitly declared
// Go is a compiled language

package main

import (
	"encoding/base32"
	"fmt"
	"reflect"
)

func main() {
	var a = "initial" // Will be inferred as string type
	fmt.Println(reflect.TypeOf(a))
	fmt.Println(a)

	// Multiple declaration
	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int // initialized zero-valued 
	fmt.Println(e)

	// := shorthand for declaring and init variable 
	// only available inside functions
	f := "bad apple"
	fmt.Println(f)
}