//let's start slowly
package main

import (
	"fmt"
	"os"
)

func main() {
	homeDir := os.ExpandEnv("$HOME")
	fmt.Printf("Your home directory is %v\n", homeDir)
}
