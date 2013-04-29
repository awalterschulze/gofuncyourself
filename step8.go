//since a function is a type it can be placed in a slice (list)
//Thank you for the idea http://jordanorelli.tumblr.com/post/42369331748/function-types-in-go-golang
package main

import (
	"fmt"
)

type pair func(a, b int)

func main() {
	pairs := []pair{
		func(a, b int) { fmt.Printf("%v + %v = %v\n", a, b, a+b) },
		func(a, b int) { fmt.Printf("%v - %v = %v\n", a, b, a-b) },
		func(a, b int) { fmt.Printf("%v * %v = %v\n", a, b, a*b) },
	}
	a := 5
	b := 7
	for _, p := range pairs {
		p(a, b)
	}
}
