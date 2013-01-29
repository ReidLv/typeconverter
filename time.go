package typeconverter

import (
	js "encoding/json"
	xl "encoding/xml"
	"time"
)

type Timer interface {
	Time() time.Time
}

type Time time.Time

func (ø Time) String() string  { return time.Time(ø).Format(time.RFC3339) }
func (ø Time) Time() time.Time { return time.Time(ø) }
func (ø Time) Int() int        { return int(ø.Time().Unix()) }

func (ø Time) Json() string {
	b, err := js.Marshal(time.Time(ø))
	if err != nil {
		panic("can't convert " + ø.String() + " to json")
	}
	return string(b)
}

func (ø Time) Xml() string {
	b, err := xl.Marshal(time.Time(ø))
	if err != nil {
		panic("can't convert " + ø.String() + " to xml")
	}
	return string(b)
}
