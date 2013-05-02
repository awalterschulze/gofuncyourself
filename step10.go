//lets make different walkers, less efficient, but that is not why we are here.
//Since functions are just types we can make a slice (list) of them.
package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func main() {
	files := make(map[string]time.Time)
	dirs := make(map[string]time.Time)
	homeDir := os.ExpandEnv("$HOME")
	fileWalk := func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files[path] = info.ModTime()
		}
		return nil
	}
	dirWalk := func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			dirs[path] = info.ModTime()
		}
		return nil
	}
	walkers := []filepath.WalkFunc{fileWalk, dirWalk}
	for _, w := range walkers {
		filepath.Walk(homeDir, w)
	}
	fmt.Printf("files: %v\n", files)
	fmt.Printf("dirs: %v\n", dirs)
}
