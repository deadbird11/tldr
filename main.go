package main

import (
	"fmt"
	"os"

	"github.com/deadbird11/tldr/downloading"
)

func main() {
	args := os.Args[1:]
	desc, _ := downloading.GetCommandDesc(args[0])
	fmt.Println(desc)
}
