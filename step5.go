//Or return a function from a method, there are many ways to func a cat
package main

import (
	"fmt"
	"os"
	"path/filepath"
)

type state struct {
}

func (this *state) newWalker() func(path string, info os.FileInfo, err error) error {
	return func(path string, info os.FileInfo, err error) error {
		fmt.Printf("%v\n", path)
		return nil
	}
}

func main() {
	s := &state{}
	homeDir := os.ExpandEnv("$HOME")
	filepath.Walk(homeDir, s.newWalker())
}
