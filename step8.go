//lets start using os.FileInfo to start doing something useful for a change
package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

type state struct {
	files map[string]time.Time
	dirs  map[string]time.Time
}

func newState() *state {
	return &state{make(map[string]time.Time), make(map[string]time.Time)}
}

func (this *state) newWalker() filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			this.dirs[path] = info.ModTime()
		} else {
			this.files[path] = info.ModTime()
		}
		return nil
	}
}

func main() {
	s := newState()
	homeDir := os.ExpandEnv("$HOME")
	fn := s.newWalker()
	filepath.Walk(homeDir, fn)
	fmt.Printf("files: %v\n", s.files)
	fmt.Printf("dirs: %v\n", s.dirs)
}
