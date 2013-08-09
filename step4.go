//Or use a method from a function
//This is very new for go
//I think some people call this a functor?
package main

import (
	"fmt"
	"os"
	"path/filepath"
)

type state struct {
}

func (this *state) walker(path string, info os.FileInfo, err error) error {
	fmt.Printf("%v\n", path)
	return nil
}

func main() {
	s := &state{}
	homeDir := os.ExpandEnv("$HOME")
	filepath.Walk(homeDir, s.walker)
}
