package typeconverter

import (
	"encoding/json"
	"time"
)

type Timer interface {
	Time() time.Time
}

type Time time.Time

func (ø Time) String() string  { return time.Time(ø).Format(time.RFC3339) }
func (ø Time) Time() time.Time { return time.Time(ø) }
func (ø Time) Json() string {
	b, err := json.Marshal(time.Time(ø))
	if err != nil {
		panic("can't convert " + ø.String() + " to json")
	}
	return string(b)
}
