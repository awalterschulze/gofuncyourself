//since functions are types we can use it as a struct member
package main

import (
	"fmt"
)

type pair struct {
	do func(a, b int)
}

func main() {
	pairs := []pair{
		pair{do: func(a, b int) { fmt.Printf("%v + %v = %v\n", a, b, a+b) }},
		pair{do: func(a, b int) { fmt.Printf("%v - %v = %v\n", a, b, a-b) }},
		pair{func(a, b int) { fmt.Printf("%v * %v = %v\n", a, b, a*b) }},
	}
	a := 5
	b := 7
	for _, p := range pairs {
		p.do(a, b)
	}
}
