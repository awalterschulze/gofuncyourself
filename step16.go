//Ideas stolen from rob pike's example before methods could be assigned like functions
//https://code.google.com/p/go/source/browse/src/pkg/sort/example_keys_test.go?spec=svn2bf8d07c14c178039b3220beb5d76fc07b56358a&name=2bf8d07c14c1&r=2bf8d07c14c178039b3220beb5d76fc07b56358a

//cast functions: words, forward, reverse to type By
//call their Sort methods
//Sort creates a struct with a function member
package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
)

type By func(q1, q2 *file) bool

func (by By) Sort(files []file) {
	fs := &fileSorter{
		files: files,
		by:    by,
	}
	sort.Sort(fs)
}

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

type fileSorter struct {
	files []file
	by    func(f1, f2 *file) bool
}

func (s *fileSorter) Len() int {
	return len(s.files)
}

func (s *fileSorter) Swap(i, j int) {
	s.files[i], s.files[j] = s.files[j], s.files[i]
}

func (s *fileSorter) Less(i, j int) bool {
	return s.by(&s.files[i], &s.files[j])
}

func (this file) String() string {
	return fmt.Sprintf("%v(%v:%v)", this.path, this.ModTime(), this.Size())
}

func main() {
	files := []file{}
	dirs := []file{}
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

	alpha := func(f1, f2 *file) bool {
		return f1.path < f2.path
	}
	modTime := func(f1, f2 *file) bool {
		return f1.ModTime().UnixNano() < f2.ModTime().UnixNano()
	}
	reverse := func(f1, f2 *file) bool {
		return f1.path > f2.path
	}

	By(reverse).Sort(files)
	fmt.Printf("files: %v\n", files)
	By(alpha).Sort(files)
	fmt.Printf("files: %v\n", files)
	By(modTime).Sort(files)
	for _, f := range files {
		fmt.Printf("%v\n", f)
	}
}
