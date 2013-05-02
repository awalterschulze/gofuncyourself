//lets simplify things for a change
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
	walker := func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			dirs[path] = info.ModTime()
		} else {
			files[path] = info.ModTime()
		}
		return nil
	}
	filepath.Walk(homeDir, walker)
	fmt.Printf("files: %v\n", files)
	fmt.Printf("dirs: %v\n", dirs)
}
