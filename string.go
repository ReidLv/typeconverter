package typeconverter

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

type String string

type Stringer interface {
	String() string
}

func (ø String) String() string { return string(ø) }

func (ø String) Int() int {
	ii, err := strconv.ParseInt(ø.String(), 10, 32)
	if err != nil {
		var f float64
		f, err = strconv.ParseFloat(ø.String(), 64)
		if err != nil {
			panic("can't convert " + ø.String() + "to int")
		}
		return Float(f).Int()
	}
	return int(ii)
}

func (ø String) Float() float64 {
	f, err := strconv.ParseFloat(ø.String(), 64)
	if err != nil {
		panic("can't convert " + ø.String() + "to float64")
	}
	return f
}

func (ø String) Time() time.Time {
	tt, err := time.Parse(time.RFC3339, ø.String())
	if err != nil {
		panic("can't convert " + ø.String() + "to time")
	}
	return tt
}

func (ø String) Bool() bool {
	b, err := strconv.ParseBool(ø.String())
	if err != nil {
		panic("can't convert " + ø.String() + "to bool")
	}
	return b
}

func (ø String) Json() string {
	b, err := json.Marshal(ø.String())
	if err != nil {
		panic("can't convert " + fmt.Sprintf("%v", ø.String()) + " to json")
	}
	return string(b)
}

// checks, if something is a string or Stringer
func isString(i interface{}) bool {
	if _, ok := i.(string); ok {
		return true
	}

	if _, ok := i.(fmt.Stringer); ok {
		return true
	}

	return false
}
