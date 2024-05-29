package sort

import (
	"os"
	"sort"
	"strings"
)

// DirFirst sorts entries by directory first
func DirFirst(entries []os.DirEntry) {
	sort.Slice(entries, func(i, j int) bool {
		// Directories first
		if entries[i].IsDir() != entries[j].IsDir() {
			return entries[i].IsDir()
		}

		// Case insensitive
		return strings.ToLower(entries[i].Name()) < strings.ToLower(entries[j].Name())
	})
}
