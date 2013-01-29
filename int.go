package typeconverter

import (
	xl "encoding/xml"
	"fmt"
	"time"
)

type Inter interface {
	Int() int
}

type Int int

func (ø Int) String() string  { return fmt.Sprintf("%v", ø.Int()) }
func (ø Int) Int() int        { return int(ø) }
func (ø Int) Float() float64  { return float64(ø) }
func (ø Int) Json() string    { return ø.String() }
func (ø Int) Time() time.Time { return time.Unix(int64(ø), 0) }

func (ø Int) Xml() string {
	b, err := xl.Marshal(int(ø))
	if err != nil {
		panic("can't convert " + ø.String() + " to xml")
	}
	return string(b)
}

type Int64 int64

func (ø Int64) String() string  { return fmt.Sprintf("%v", ø.Int()) }
func (ø Int64) Int() int        { return int(ø) }
func (ø Int64) Float() float64  { return float64(ø) }
func (ø Int64) Json() string    { return ø.String() }
func (ø Int64) Time() time.Time { return time.Unix(int64(ø), 0) }

func (ø Int64) Xml() string {
	b, err := xl.Marshal(int(ø))
	if err != nil {
		panic("can't convert " + ø.String() + " to xml")
	}
	return string(b)
}
