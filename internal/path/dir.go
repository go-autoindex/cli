package path

import "path"

// JoinDir is a wrapper around [path.Join], and adds a trailing slash.
func JoinDir(elems ...string) string {
	return path.Join(elems...) + "/"
}
