//lets sort again, now lets create a sortable list and sort it alphabetically
package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
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

type fileList struct {
	list []file
}

func (this *fileList) Len() int {
	return len(this.list)
}

func (this *fileList) Swap(i, j int) {
	this.list[i], this.list[j] = this.list[j], this.list[i]
}

func (this *fileList) Less(i, j int) bool {
	return this.list[i].path < this.list[j].path
}

func (this file) String() string {
	return fmt.Sprintf("%v(%v:%v)", this.path, this.ModTime(), this.Size())
}

func main() {
	files := &fileList{}
	dirs := &fileList{}
	homeDir := os.ExpandEnv("$HOME")
	fileWalk := func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files.list = append(files.list, file{path, info})
		}
		return nil
	}
	dirWalk := func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			dirs.list = append(dirs.list, file{path, info})
		}
		return nil
	}
	fileFilter := filter(fileWalk)
	walkers := []filepath.WalkFunc{fileFilter.size, dirWalk}
	for _, w := range walkers {
		filepath.Walk(homeDir, w)
	}
	sort.Sort(files)
	for _, f := range files.list {
		fmt.Printf("%v\n", f)
	}
}