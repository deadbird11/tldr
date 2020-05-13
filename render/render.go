package render

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"strings"

	"github.com/fatih/color"
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
			char := string(c)
			switch char {
			case "#":
				// skipping space after #
				r.ReadRune()
			case "`":
				renderCode(r)
			default:
				fmt.Print(char)
			}
		}

	}
}

// renderCode - renders chars between `` as magenta
// TODO: make this more complex
func renderCode(r *bufio.Reader) {
	codeStr, err := r.ReadString([]byte("`")[0])
	if err != nil {
		return
	}
	codeStr = codeStr[:len(codeStr)-1]
	color.Cyan(codeStr)
}
