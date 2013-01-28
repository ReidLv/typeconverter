package typeconverter

import (
	"encoding/json"
)

type Arrayer interface {
	Array() []interface{}
}

type Array []interface{}

func (ø Array) Array() []interface{} { return []interface{}{ø} }
func (ø Array) String() string       { return ø.Json() }

func (ø Array) Json() string {
	b, err := json.Marshal(ø)
	if err != nil {
		panic("can't convert " + ø.String() + " to json")
	}
	return string(b)
}
