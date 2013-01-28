package typeconverter

import (
	"encoding/json"
	"time"
)

type Json string
type Jsoner interface {
	Json() string
}

func (ø Json) Int() int {
	var i float64
	err := json.Unmarshal([]byte(ø), &i)
	if err != nil {
		panic("can't convert " + ø + " to int")
	}
	return Float(i).Int()
}

func (ø Json) Float() float64 {
	var i float64
	err := json.Unmarshal([]byte(ø), &i)
	if err != nil {
		panic("can't convert " + ø + " to float")
	}
	return i
}

func (ø Json) Time() time.Time {
	var i time.Time
	err := json.Unmarshal([]byte(ø), &i)
	if err != nil {
		panic("can't convert " + ø + " to time")
	}
	return i
}

func (ø Json) String() string { return string(ø) }
func (ø Json) Json() string   { return string(ø) }

func (ø Json) Bool() bool {
	var i bool
	err := json.Unmarshal([]byte(ø), &i)
	if err != nil {
		panic("can't convert " + ø + " to bool")
	}
	return i
}

func (ø Json) Array() []interface{} {
	ia := []interface{}{}
	err := json.Unmarshal([]byte(ø), &ia)
	if err != nil {
		panic("can't convert " + ø + " to array")
	}
	return ia
}

func (ø Json) Map() map[string]interface{} {
	ia := map[string]interface{}{}
	err := json.Unmarshal([]byte(ø), &ia)
	if err != nil {
		panic("can't convert " + ø + " to array")
	}
	return ia
}
