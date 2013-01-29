package typeconverter

import (
	xl "encoding/xml"
	"fmt"
)

type Booler interface {
	Bool() bool
}

type Bool bool

func (ø Bool) String() string { return fmt.Sprintf("%v", ø.Bool()) }
func (ø Bool) Json() string   { return ø.String() }
func (ø Bool) Bool() bool     { return bool(ø) }

func (ø Bool) Xml() string {
	b, err := xl.Marshal(bool(ø))
	if err != nil {
		panic("can't convert " + ø.String() + " to xml")
	}
	return string(b)
}
