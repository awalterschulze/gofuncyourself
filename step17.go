//lets add flags
package main

import (
	"flag"
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
	var filterType string
	flag.StringVar(&filterType, "filter", "files", "files or dirs or size")
	var sortType string
	flag.StringVar(&sortType, "sort", "alpha", "alpha or reverse or time")
	flag.Parse()

	files := []file{}
	dirs := []file{}
	homeDir := os.ExpandEnv("$HOME")

	walkers := make(map[string]filepath.WalkFunc)

	walkers["files"] = func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, file{path, info})
		}
		return nil
	}
	walkers["dirs"] = func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			dirs = append(dirs, file{path, info})
		}
		return nil
	}
	walkers["size"] = filter(walkers["files"]).size

	filepath.Walk(homeDir, walkers[filterType])

	sorters := make(map[string]By)

	sorters["alpha"] = func(f1, f2 *file) bool {
		return f1.path < f2.path
	}
	sorters["time"] = func(f1, f2 *file) bool {
		return f1.ModTime().UnixNano() < f2.ModTime().UnixNano()
	}
	sorters["reverse"] = func(f1, f2 *file) bool {
		return f1.path > f2.path
	}

	By(sorters[sortType]).Sort(files)

	for _, f := range files {
		fmt.Printf("%v\n", f)
	}
}
