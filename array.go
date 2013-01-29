package typeconverter

import (
	js "encoding/json"
	xl "encoding/xml"
)

type Arrayer interface {
	Array() []interface{}
}

type Array []interface{}

func (ø Array) Array() (a []interface{}) {
	a = ø
	return
}
func (ø Array) String() string { return ø.Json() }

func (ø Array) Json() string {
	b, err := js.Marshal(ø)
	if err != nil {
		panic("can't convert " + ø.String() + " to json")
	}
	return string(b)
}

func (ø Array) Xml() string {
	b, err := xl.Marshal(ø)
	if err != nil {
		panic("can't convert " + ø.String() + " to xml")
	}
	return string(b)
}
