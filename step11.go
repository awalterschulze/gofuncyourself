//since functions are types we can declare methods on them
package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

type filter func(path string, info os.FileInfo, err error) error

func (this filter) size(path string, info os.FileInfo, err error) error {
	if info.Size() > 1e7 {
		return this(path, info, err)
	}
	return nil
}

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
	fileFilter := filter(fileWalk)
	walkers := []filepath.WalkFunc{fileFilter.size, dirWalk}
	for _, w := range walkers {
		filepath.Walk(homeDir, w)
	}
	fmt.Printf("files: %v\n", files)
	//fmt.Printf("dirs: %v\n", dirs)
}
