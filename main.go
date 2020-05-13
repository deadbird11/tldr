package main

import (
	"fmt"
	"os"

	"github.com/deadbird11/tldr/downloading"
	"github.com/deadbird11/tldr/render"
)

func main() {
	args := os.Args[1:]
	desc := downloading.GetCommandDesc(args[0])
	if desc == nil {
		fmt.Printf("could not provide description for command '%s'", args[0])
	}
	render.MD(desc)
}
