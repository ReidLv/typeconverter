package typeconverter

import (
	"encoding/json"
	"fmt"
	"time"
)

type convert struct {
	value interface{}
}

func Convert(i interface{}) *convert {
	switch t := i.(type) {
	case Int, Int64, Float, Float32, Bool, String, Time, Json, Map, Array, Default:
		return &convert{t}
	case int:
		return &convert{Int(t)}
	case int32:
		return &convert{Int(t)}
	case int64:
		return &convert{Int64(t)}
	case float64:
		return &convert{Float(t)}
	case float32:
		return &convert{Float32(t)}
	case string:
		return &convert{String(t)}
	case bool:
		return &convert{Bool(t)}
	case time.Time:
		return &convert{Time(t)}
	case map[string]interface{}:
		return &convert{Map(t)}
	case []interface{}:
		return &convert{Array(t)}
	default:
		if IsString(t) {
			return &convert{String(fmt.Sprintf("%s", t))}
		} else {
			b, err := json.Marshal(t)
			if err != nil {
				panic("can't convert " + fmt.Sprintf("%#v", t) + " to anything reasonable or jsonable")
			}
			return &convert{String(string(b))}
		}
	}
	return &convert{Default(0)}
}

func (ø convert) ToInt() int                    { return ø.value.(Inter).Int() }
func (ø convert) ToString() string              { return ø.value.(Stringer).String() }
func (ø convert) ToFloat() float64              { return ø.value.(Floater).Float() }
func (ø convert) ToBool() bool                  { return ø.value.(Booler).Bool() }
func (ø convert) ToTime() time.Time             { return ø.value.(Timer).Time() }
func (ø convert) ToJson() string                { return ø.value.(Jsoner).Json() }
func (ø convert) ToMap() map[string]interface{} { return ø.value.(Mapper).Map() }
func (ø convert) ToArray() []interface{}        { return ø.value.(Arrayer).Array() }
