package typeconverter

import (
	"fmt"
)

type Booler interface {
	Bool() bool
}

type Bool bool

func (ø Bool) String() string { return fmt.Sprintf("%v", ø.Bool()) }
func (ø Bool) Json() string   { return ø.String() }
func (ø Bool) Bool() bool     { return bool(ø) }
