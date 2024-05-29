package path

import (
	"os"

	"github.com/go-autoindex/autoindex"
	"github.com/go-autoindex/cli/internal/sort"
)

// WalkAndIndex walks through the directory and generates the index.
func WalkAndIndex(dir string) error {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return err
	}
	sort.DirFirst(entries)

	return autoindex.Gen(walk(dir, entries))
}

// walk walks through the directory and generates info for the index.
func walk(dir string, entries []os.DirEntry) autoindex.Info {
	info := autoindex.Info{
		Dir:     dir,
		Entries: make([]string, 0, len(entries)),
	}

	for _, entry := range entries {
		name := entry.Name()
		if ignored(name) {
			continue
		}

		if entry.IsDir() {
			// Mark name as a directory
			name = JoinDir(name)

			WalkAndIndex(JoinDir(dir, name))
		}

		info.Entries = append(info.Entries, name)
	}

	return info
}
