// (c) 2011-2014 Alexander Solovyov
// under terms of ISC License

package main

import (
	"fmt"
	"github.com/wsxiaoys/terminal/color"
)

type Printer struct {
	NoColors bool
	NoGroup bool
	previous string
}

func (p *Printer) Printf(colorfmt, plainfmt string,
	args... interface{}) {

	if p.NoColors {
		fmt.Printf(plainfmt, args...)
	} else {
		color.Printf(colorfmt, args...)
	}
}

func (p *Printer) Sprintf(colorfmt, plainfmt string,
	args... interface{}) string {

	if p.NoColors {
		return fmt.Sprintf(plainfmt, args...)
	} else {
		return color.Sprintf(colorfmt, args...)
	}
}

func (p *Printer) FilePrintf(fn, colorfmt, plainfmt string,
	args... interface{}) {

	if p.NoGroup {
		p.Printf("@g%s:", "%s:", fn)
	} else if fn != p.previous {
		if p.previous != "" {
			fmt.Println("")
		}
		p.Printf("@g%s\n", "%s\n", fn)
		p.previous = fn
	}

	p.Printf(colorfmt, plainfmt, args...)
}
