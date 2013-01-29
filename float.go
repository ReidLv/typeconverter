package typeconverter

import (
	xl "encoding/xml"
	"fmt"
	"math"
	"strconv"
	"time"
)

type Floater interface {
	Float() float64
}

// stolen from https://groups.google.com/forum/?fromgroups=#!topic/golang-nuts/ITZV08gAugI
// return rounded version of x with prec precision.
func roundViaFloat(x float64, prec int) float64 {
	frep := strconv.FormatFloat(x, 'g', prec, 64)
	f, _ := strconv.ParseFloat(frep, 64)
	return f
}

func toInt(x float64) int {
	return int(math.Floor(roundViaFloat(x, 0)))
}

func toInt64(x float64) int64 {
	return int64(math.Floor(roundViaFloat(x, 0)))
}

type Float float64

func (ø Float) String() string  { return fmt.Sprintf("%v", ø.Float()) }
func (ø Float) Float() float64  { return float64(ø) }
func (ø Float) Json() string    { return ø.String() }
func (ø Float) Int() int        { return toInt(ø.Float()) }
func (ø Float) Time() time.Time { return time.Unix(toInt64(ø.Float()), 0) }

func (ø Float) Xml() string {
	b, err := xl.Marshal(float64(ø))
	if err != nil {
		panic("can't convert " + ø.String() + " to xml")
	}
	return string(b)
}

type Float32 float32

func (ø Float32) String() string  { return fmt.Sprintf("%v", ø.Float()) }
func (ø Float32) Float() float64  { return float64(ø) }
func (ø Float32) Json() string    { return ø.String() }
func (ø Float32) Int() int        { return toInt(ø.Float()) }
func (ø Float32) Time() time.Time { return time.Unix(toInt64(ø.Float()), 0) }

func (ø Float32) Xml() string {
	b, err := xl.Marshal(float32(ø))
	if err != nil {
		panic("can't convert " + ø.String() + " to xml")
	}
	return string(b)
}
