package main

import (
	"fmt"
	"os"

	"github.com/deadbird11/tldr/downloading"
	"github.com/deadbird11/tldr/render"
)

func main() {
	args := os.Args[1:]
	desc, err := downloading.GetCommandDesc(args[0])
	if err != nil {
		fmt.Println(err.Error())
	}
	render.MD(desc)
}
