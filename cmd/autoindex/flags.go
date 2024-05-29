package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/go-autoindex/autoindex"
	"github.com/go-autoindex/cli/internal/path"
)

// flags represents the command line flags
type flags struct {
	// dir is the directory to index
	dir string
}

var (
	dir = flag.String("dir", ".", "directory to index")
	dry = flag.Bool("dry", false, "dry run")
)

// parseFlags parses the command line flags.
func parseFlags() *flags {
	flag.Parse()

	if *dry {
		fmt.Fprintln(os.Stderr, "dry run")
		autoindex.Opts.Dry = true
	}

	return &flags{
		dir: path.JoinDir(*dir),
	}
}
