package main

import (
	"github.com/docopt/docopt-go"
	"github.com/pkg/errors"
	"github.com/weblair/lair/generator"
	"os"
)

const VERSION = "Lair 0.1.0"
const USAGE = `Lair

Usage:
  lair new <owner> <project_name>

Options:
  -h --help   Show this screen.
  --version   Show version.
`

func main() {
	args, err := docopt.Parse(USAGE, os.Args[1:], true, VERSION, false)
	if err != nil {
		panic(errors.WithMessage(err, "failed to parse doc"))
	}

	if args["new"].(bool) {
		generator.NewGinProject(args["<owner>"].(string), args["<project_name>"].(string))
	}
}
