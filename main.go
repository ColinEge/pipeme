package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/atotto/clipboard"
)

func main() {
	// CLI flags
	showFiles := flag.Bool("f", false, "Include files in the tree")
	dir := flag.String("d", ".", "Directory to print")
	copyClipboard := flag.Bool("c", false, "Copy output to clipboard")
	flag.Parse()

	// Generate tree into builder
	tree := writeTree(*dir, *showFiles)

	// Output to CLI
	fmt.Print(tree)

	// Copy to clipboard if requested
	if *copyClipboard {
		err := clipboard.WriteAll(tree)
		if err != nil {
			fmt.Println("Failed to copy to clipboard:", err)
		} else {
			fmt.Println("\n[Tree copied to clipboard]")
		}
	}
}

func writeTree(path string, showFiles bool) string {
	var builder strings.Builder
	writeTreeRecurse(path, "", showFiles, &builder)
	return builder.String()
}

func writeTreeRecurse(path string, prefix string, showFiles bool, builder *strings.Builder) {
	entries, err := os.ReadDir(path)
	if err != nil {
		fmt.Println(err)
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
			builder.WriteString(fmt.Sprintf("%s└─ %s\n", prefix, entry.Name()))
			newPrefix = prefix + "  "
		} else {
			builder.WriteString(fmt.Sprintf("%s├─ %s\n", prefix, entry.Name()))
			newPrefix = prefix + "│ "
		}

		if entry.IsDir() {
			writeTreeRecurse(filepath.Join(path, entry.Name()), newPrefix, showFiles, builder)
		}
	}
}
