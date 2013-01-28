package typeconverter

import (
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
func RoundViaFloat(x float64, prec int) float64 {
	frep := strconv.FormatFloat(x, 'g', prec, 64)
	f, _ := strconv.ParseFloat(frep, 64)
	return f
}

func ToInt(x float64) int {
	return int(math.Floor(RoundViaFloat(x, 0)))
}

func ToInt64(x float64) int64 {
	return int64(math.Floor(RoundViaFloat(x, 0)))
}

type Float float64

func (ø Float) String() string  { return fmt.Sprintf("%v", ø.Float()) }
func (ø Float) Float() float64  { return float64(ø) }
func (ø Float) Json() string    { return ø.String() }
func (ø Float) Int() int        { return ToInt(ø.Float()) }
func (ø Float) Time() time.Time { return time.Unix(ToInt64(ø.Float()), 0) }

type Float32 float32

func (ø Float32) String() string  { return fmt.Sprintf("%v", ø.Float()) }
func (ø Float32) Float() float64  { return float64(ø) }
func (ø Float32) Json() string    { return ø.String() }
func (ø Float32) Int() int        { return ToInt(ø.Float()) }
func (ø Float32) Time() time.Time { return time.Unix(ToInt64(ø.Float()), 0) }
