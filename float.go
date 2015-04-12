package typeconverter

import (
	xl "encoding/xml"
	"fmt"
	"strconv"
	"time"
)

type Floater interface {
	Float() float64
}

// stolen from https://groups.google.com/forum/?fromgroups=#!topic/golang-nuts/ITZV08gAugI
// return rounded version of x with prec precision.
func roundviaFloat(x float64, prec int) float64 {
	frep := strconv.FormatFloat(x, 'g', prec, 64)
	f, _ := strconv.ParseFloat(frep, 64)
	return f
}

func numDig(f float64) int {
	return len(fmt.Sprintf("%0.0f", f))
}

func FloatToInt(x float64) int {
	return int(roundviaFloat(x, numDig(x)))
}

func RoundFloat(x float64, decimals int) float64 {
	return roundviaFloat(x, numDig(x)+decimals)
}

func toInt(x float64) int {
	return FloatToInt(x)
}

func toInt64(x float64) int64 {
	return int64(FloatToInt(x))
}

func Float(f float64) FloatType { return FloatType(f) }

type FloatType float64

func (ø FloatType) String() string  { return fmt.Sprintf("%v", ø.Float()) }
func (ø FloatType) Float() float64  { return float64(ø) }
func (ø FloatType) Json() string    { return ø.String() }
func (ø FloatType) Int() int        { return toInt(ø.Float()) }
func (ø FloatType) Time() time.Time { return time.Unix(toInt64(ø.Float()), 0) }

func (ø FloatType) Xml() string {
	b, err := xl.Marshal(float64(ø))
	if err != nil {
		panic("can't convert " + ø.String() + " to xml")
	}
	return string(b)
}

func Float32(f float32) FloatType32 { return FloatType32(f) }

type FloatType32 float32

func (ø FloatType32) String() string  { return fmt.Sprintf("%v", ø.Float()) }
func (ø FloatType32) Float() float64  { return float64(ø) }
func (ø FloatType32) Json() string    { return ø.String() }
func (ø FloatType32) Int() int        { return toInt(ø.Float()) }
func (ø FloatType32) Time() time.Time { return time.Unix(toInt64(ø.Float()), 0) }

func (ø FloatType32) Xml() string {
	b, err := xl.Marshal(float32(ø))
	if err != nil {
		panic("can't convert " + ø.String() + " to xml")
	}
	return string(b)
}
