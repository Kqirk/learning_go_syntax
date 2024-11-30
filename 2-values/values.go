package main

import "fmt"

func main() {
	var prt = fmt.Printf
	fmt.Println("go" + "lang")
	prt("go" + "lang using prt") // Possible but not recommended for readability

	fmt.Println("1+1 =", 1 + 1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)
	fmt.Println("7/3 =", 7/3)
	fmt.Println("7/3.0", 7/3.0)
	// fmt.Println("-" * 26) not possible 

	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}