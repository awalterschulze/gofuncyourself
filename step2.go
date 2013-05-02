//Here we see our first function pass as a parameter, we also saw this last time https://code.google.com/p/nogotovogo/source/browse/step2.go
package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func walker(path string, info os.FileInfo, err error) error {
	fmt.Printf("%v\n", path)
	return nil
}

func main() {
	homeDir := os.ExpandEnv("$HOME")
	fmt.Printf("Your home directory is %v\n", homeDir)
	//This is basically just an ls command, since walker just prints out the paths
	filepath.Walk(homeDir, walker)
	//We could also have done it with a closure
	//filepath.Walk(homeDir, func(path string, info os.FileInfo, err error) error {
	//	fmt.Printf("%v\n", path)
	//	return nil
	//})
}
