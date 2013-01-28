package typeconverter

import (
	"encoding/json"
	"fmt"
)

type Mapper interface {
	Map() map[string]interface{}
}

type Map map[string]interface{}

func (ø Map) Map() map[string]interface{} { return map[string]interface{}(ø) }
func (ø Map) String() string              { return ø.Json() }

func (ø Map) Json() string {
	b, err := json.Marshal(ø)
	if err != nil {
		panic("can't convert " + fmt.Sprintf("%v", ø) + " to json")
	}
	return string(b)
}
