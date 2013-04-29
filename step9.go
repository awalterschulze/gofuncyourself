//since functions are types we can declare methods on functions
package main

import (
	"fmt"
)

type pair func(a, b int)

func (this pair) swap(a, b int) {
	this(b, a)
}

func main() {
	pairs := []pair{
		func(a, b int) { fmt.Printf("%v + %v = %v\n", a, b, a+b) },
		func(a, b int) { fmt.Printf("%v - %v = %v\n", a, b, a-b) },
		func(a, b int) { fmt.Printf("%v * %v = %v\n", a, b, a*b) },
	}
	for i := 0; i < 3; i++ {
		pairs = append(pairs, pairs[i].swap)
	}
	a := 5
	b := 7
	for _, p := range pairs {
		p(a, b)
	}
}
