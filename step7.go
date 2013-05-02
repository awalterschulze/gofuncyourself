//if we declare our own function type we will need to cast it to the filepath.WalkFunc function type
package main

import (
	"fmt"
	"os"
	"path/filepath"
)

type WalkFunc func(path string, info os.FileInfo, err error) error

type state struct {
}

func (this *state) newWalker() WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		fmt.Printf("%v\n", path)
		return nil
	}
}

func main() {
	s := &state{}
	homeDir := os.ExpandEnv("$HOME")
	fn := filepath.WalkFunc(s.newWalker())
	filepath.Walk(homeDir, fn)
}
