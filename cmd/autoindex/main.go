package main

import (
	"github.com/go-autoindex/autoindex"
	"github.com/go-autoindex/cli/internal/path"
)

func main() {
	flags := parseFlags()
	autoindex.Opts.Root = flags.dir

	if err := path.WalkAndIndex(flags.dir); err != nil {
		panic(err)
	}
}
