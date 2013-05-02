//lets sort again, but first lets create a list of files rather than a map
package main

import (
	"fmt"
	"os"
	"path/filepath"
)

type filter func(path string, info os.FileInfo, err error) error

func (this filter) size(path string, info os.FileInfo, err error) error {
	if info.Size() > 1e7 {
		return this(path, info, err)
	}
	return nil
}

type file struct {
	path string
	os.FileInfo
}

func (this file) String() string {
	return fmt.Sprintf("%v(%v:%v)", this.path, this.ModTime(), this.Size())
}

func main() {
	files := make([]file, 0)
	dirs := make([]file, 0)
	homeDir := os.ExpandEnv("$HOME")
	fileWalk := func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, file{path, info})
		}
		return nil
	}
	dirWalk := func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			dirs = append(dirs, file{path, info})
		}
		return nil
	}
	fileFilter := filter(fileWalk)
	walkers := []filepath.WalkFunc{fileFilter.size, dirWalk}
	for _, w := range walkers {
		filepath.Walk(homeDir, w)
	}
	for _, f := range files {
		fmt.Printf("%v\n", f)
	}
}
