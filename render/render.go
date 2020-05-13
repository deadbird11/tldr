package render

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"strings"
)

var buffer []byte

// MD - takes a string written in github markdown
// and prints it out in a prettier way, with colors
func MD(content *string) {
	r := bufio.NewReader(strings.NewReader(*content))

	for {
		if c, _, err := r.ReadRune(); err != nil {
			if err == io.EOF {
				break
			} else {
				log.Fatal(err)
			}
		} else {
			fmt.Printf("%s", string(c))
		}

	}
}
