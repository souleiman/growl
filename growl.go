package main

import (
	"fmt"
	docopt "github.com/docopt/docopt-go"
	"github.com/souleiman/growl/model"
)

var usage string = `Usage:
	growl -u <url>

-u, --url   The resource URL`

func main() {
	args, _ := docopt.Parse(usage, nil, true, "GrOwl 1.0", true)
	if args["-u"].(bool) {
		//crawl.OutputURL(args["<url>"].(string))
	}

	modal, _ := model.NewModel(args["<url>"].(string))
	fmt.Println(modal)
}
