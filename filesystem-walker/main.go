package main

import (
	"fmt"
	"os"
	"path"
)

// PrintPath recursively prints all path(s) and their sub-path(s) given a root path.
func PrintPath(rootPath string) {
	fi, err := os.Stat(rootPath)
	if os.IsNotExist(err) {
		return
	}

	fmt.Println(rootPath)
	if !fi.IsDir() {
		return
	}

	dirs, err := os.ReadDir(rootPath)
	if err != nil {
		return
	}

	for _, dir := range dirs {
		PrintPath(path.Join(rootPath, dir.Name()))
	}
}

func main() {}
