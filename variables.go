package main

import "fmt"

func main() {
	var a = "initial"
	fmt.Println(a)

	// var declares one or more variables.
	var b, c int = 1, 2
	fmt.Println(b, c)

	// go will infer the type when the variable is initialized.
	var d = true
	fmt.Println(d)

	// zero-valued by default.
	var e int
	fmt.Println(e)

	var g string
	fmt.Println(g)

	f := "apple"
	fmt.Println(f)
}
