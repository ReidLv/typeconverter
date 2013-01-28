package typeconverter

import (
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

type Int64 int64

func (ø Int64) String() string  { return fmt.Sprintf("%v", ø.Int()) }
func (ø Int64) Int() int        { return int(ø) }
func (ø Int64) Float() float64  { return float64(ø) }
func (ø Int64) Json() string    { return ø.String() }
func (ø Int64) Time() time.Time { return time.Unix(int64(ø), 0) }
