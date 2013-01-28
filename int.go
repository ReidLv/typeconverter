package typeconverter

import (
	"fmt"
)

type Inter interface {
	Int() int
}

type Int int

func (ø Int) String() string { return fmt.Sprintf("%v", ø.Int()) }
func (ø Int) Int() int       { return int(ø) }
func (ø Int) Float() float64 { return float64(ø) }
func (ø Int) Json() string   { return ø.String() }

type Int64 int64

func (ø Int64) String() string { return fmt.Sprintf("%v", ø.Int()) }
func (ø Int64) Int() int       { return int(ø) }
func (ø Int64) Float() float64 { return float64(ø) }
func (ø Int64) Json() string   { return ø.String() }
