package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	showFiles := flag.Bool("f", false, "Include files in the tree")
	dir := flag.String("d", ".", "Directory to print")
	flag.Parse()

	printTree(*dir, *showFiles)
}

func printTree(path string, showFiles bool) {
	printTreeRecurse(path, "", showFiles)
}

func printTreeRecurse(path string, prefix string, showFiles bool) {
	entries, err := os.ReadDir(path)
	if err != nil {
		return
	}

	var filtered []os.DirEntry
	for _, e := range entries {
		if e.IsDir() || showFiles {
			filtered = append(filtered, e)
		}
	}

	for i, entry := range filtered {
		isLast := i == len(filtered)-1
		var newPrefix string

		if isLast {
			fmt.Printf("%s└─ %s\n", prefix, entry.Name())
			newPrefix = prefix + "  "
		} else {
			fmt.Printf("%s├─ %s\n", prefix, entry.Name())
			newPrefix = prefix + "│ "
		}

		if entry.IsDir() {
			printTreeRecurse(filepath.Join(path, entry.Name()), newPrefix, showFiles)
		}
	}
}
