//a function can be a type
package main

import (
	"fmt"
	"os"
	"path/filepath"
)

//in the filepath package http://golang.org/pkg/path/filepath/#WalkFunc
//their is a type declared WalkFunc
//type WalkFunc func(path string, info os.FileInfo, err error) error

type state struct {
}

func (this *state) newWalker() filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		fmt.Printf("%v\n", path)
		return nil
	}
}

func main() {
	s := &state{}
	homeDir := os.ExpandEnv("$HOME")
	fn := s.newWalker()
	filepath.Walk(homeDir, fn)
}
