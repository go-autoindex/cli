package path

import (
	"strings"

	"github.com/go-autoindex/autoindex"
)

// ignored reports whether path should be ignored
func ignored(path string) bool {
	return isIndex(path) || isDot(path)
}

// isDot checks if the given path is a dot file.
func isDot(path string) bool {
	return strings.HasPrefix(path, ".")
}

// isIndex checks if the given path is "index.html".
func isIndex(path string) bool {
	return path == autoindex.Index()
}
