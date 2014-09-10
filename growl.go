package main

import (
	docopt "github.com/docopt/docopt-go"
	"github.com/souleiman/growl/crawl"
)

var usage string = `Usage:
	growl -u <url>

-u, --url   The resource URL`

func main() {
	args, _ := docopt.Parse(usage, nil, true, "GrOwl 1.0", true)
	if args["-u"].(bool) {
		crawl.OutputURL(args["<url>"].(string))
	}
}
