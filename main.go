package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/atotto/clipboard"
)

func main() {
	// CLI flags
	showFiles := flag.Bool("f", false, "Include files in the tree")
	dir := flag.String("d", ".", "Directory to print")
	ignore := flag.String("i", "", "Relative paths to ignore (comma separated)")
	copyClipboard := flag.Bool("c", false, "Copy output to clipboard")
	flag.Parse()

	ignores := strings.Split(*ignore, ",")
	for i, v := range ignores {
		ignores[i] = strings.TrimSpace(v)
	}

	// Generate tree into builder
	tree := writeTree(*dir, *showFiles, ignores...)

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

func writeTree(path string, showFiles bool, ignore ...string) string {
	var builder strings.Builder
	ignore = cleanIgnores(path, ignore)
	writeTreeRecurse(path, "", showFiles, ignore, &builder)
	return builder.String()

}

func writeTreeRecurse(path string, prefix string, showFiles bool, ignore []string, builder *strings.Builder) {
	entries, err := os.ReadDir(path)
	if err != nil {
		fmt.Println(err)
		return
	}

	filtered := filterEntries(path, entries, ignore, showFiles)

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
			writeTreeRecurse(filepath.Join(path, entry.Name()), newPrefix, showFiles, ignore, builder)
		}
	}
}

func filterEntries(path string, dirs []os.DirEntry, ignore []string, includeFiles bool) []os.DirEntry {
	var filtered []os.DirEntry
	for _, e := range dirs {
		if !e.IsDir() && !includeFiles {
			continue
		}
		if slices.Contains(ignore, filepath.Join(path, e.Name())) {
			continue
		}
		filtered = append(filtered, e)
	}
	return filtered
}

func cleanIgnores(path string, ignores []string) []string {
	var newIgnores []string
	for _, v := range ignores {
		newIgnores = append(newIgnores, strings.TrimPrefix(v, path+string(os.PathSeparator)))
	}
	return newIgnores
}
