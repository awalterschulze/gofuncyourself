//Now lets call a method in a closure
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
	filepath.Walk(homeDir, func(path string, info os.FileInfo, err error) error {
		return s.walker(path, info, err)
	})
}
