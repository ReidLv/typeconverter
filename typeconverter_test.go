package typeconverter

import (
	// "encoding/json"
	"fmt"
	"testing"
	"time"
)

func err(t *testing.T, msg string, is interface{}, shouldbe interface{}) {
	t.Errorf(msg+": is %#v, should be %#v\n", is, shouldbe)
}

var _ = fmt.Errorf
var ti, _ = time.Parse(time.RFC3339, "2011-01-26T18:53:18+01:00")

var toInt = map[interface{}]int{
	1:            1,
	int64(2):     2,
	float64(3.0): 3,
	float32(3.0): 3,
	Json(`3.0`):  3,
	Json(`3`):    3,
	`1`:          1,
	`1.0`:        1,
}

func TestToInt(t *testing.T) {
	for in, out := range toInt {
		if r := Convert(in).ToInt(); r != out {
			err(t, "ToInt", r, out)
		}
	}
}

var toFloat = map[interface{}]float64{
	1:            1.0,
	int64(2):     2.0,
	float64(3.5): 3.5,
	float32(3.5): 3.5,
	Json(`3.5`):  3.5,
	Json(`3`):    3.0,
	`1`:          1.0,
	`1.5`:        1.5,
}

func TestToFloat(t *testing.T) {
	for in, out := range toFloat {
		if r := Convert(in).ToFloat(); r != out {
			err(t, "ToFloat", r, out)
		}
	}
}

var toBool = map[interface{}]bool{
	true:          true,
	false:         false,
	Json(`true`):  true,
	Json(`false`): false,
	`true`:        true,
	`false`:       false,
}

func TestToBool(t *testing.T) {
	for in, out := range toBool {
		if r := Convert(in).ToBool(); r != out {
			err(t, "ToBool", r, out)
		}
	}
}

var toString = map[interface{}]string{
	1:            "1",
	int64(2):     "2",
	3.5:          "3.5",
	float32(3.5): "3.5",
	ti:           `2011-01-26T18:53:18+01:00`,
	true:         `true`,
	Json(`{}`):   `{}`,
	`hi`:         `hi`,
}

func TestToString(t *testing.T) {
	for in, out := range toString {
		if r := Convert(in).ToString(); r != out {
			err(t, "ToString", r, out)
		}
	}

	m := map[string]interface{}{"a": 3}
	out := `{"a":3}`
	if r := Convert(m).ToString(); r != out {
		err(t, "ToString", r, out)
	}

	a := []interface{}{"a", 3, 4.5}
	out = `["a",3,4.5]`
	if r := Convert(a).ToString(); r != out {
		err(t, "ToString", r, out)
	}

}

var toJson = map[interface{}]string{
	1:            "1",
	int64(2):     "2",
	3.5:          "3.5",
	float32(3.5): "3.5",
	`hi`:         `"hi"`,
	ti:           `"2011-01-26T18:53:18+01:00"`,
	true:         `true`,
	Json(`{}`):   `{}`,
}

func TestToJson(t *testing.T) {
	for in, out := range toJson {
		if r := Convert(in).ToJson(); r != out {
			err(t, "ToJson", r, out)
		}
	}

	m := map[string]interface{}{"a": 3}
	out := `{"a":3}`
	if r := Convert(m).ToJson(); r != out {
		err(t, "ToJson", r, out)
	}

	a := []interface{}{"a", 3, 4.5}
	out = `["a",3,4.5]`
	if r := Convert(a).ToJson(); r != out {
		err(t, "ToJson", r, out)
	}

}

/*
func TestConv(t *testing.T) {
	c := New()
	if r, _ := c.ToString(34); r != "34" {
		err(t, "Converter.ToString int", r, "34")
	}

	m := map[string]interface{}{"a": 34}

	if r, _ := c.ToString(m); r != `{"a":34}` {
		err(t, "Converter.ToString map[string]int", r, `{"a":34}`)
	}

	ti, _ := time.Parse(time.RFC3339, "2010-01-26T18:53:18+01:00")
	if r, _ := c.ToString(&ti); r != "2010-01-26T18:53:18+01:00" {
		err(t, "Converter.ToString time", r, "2010-01-26T18:53:18+01:00")
	}

	ti2, _ := time.Parse(time.RFC3339, "2011-01-26T18:53:18+01:00")
	times := []interface{}{&ti, &ti2}
	if r, _ := c.ToString(times); r != `[2010-01-26T18:53:18+01:00,2011-01-26T18:53:18+01:00]` {
		err(t, "Converter.ToString time", r, `[2010-01-26T18:53:18+01:00,2011-01-26T18:53:18+01:00]`)
	}

	ints := []interface{}{3, 2}

	if r, _ := c.ToString(ints); r != `[3,2]` {
		err(t, "Converter.ToString ints", r, `[3,2]`)
	}

	strs := []string{"a", "b"}
	if r, _ := c.ToString(strs); r != `["a","b"]` {
		err(t, "Converter.ToString strs", r, `["a","b"]`)
	}

	is := []int{2, 3}
	if r, _ := c.ToString(is); r != `[2,3]` {
		err(t, "Converter.ToString is", r, `[2,3]`)
	}

	fs := []float64{2.3, 3.5}
	if r, _ := c.ToString(fs); r != `[2.3,3.5]` {
		err(t, "Converter.ToString fs", r, `[2.3,3.5]`)
	}

	bs := []bool{true, false}
	if r, _ := c.ToString(bs); r != `[true,false]` {
		err(t, "Converter.ToString bs", r, `[true,false]`)
	}

	mm := map[string]string{"a": "b"}
	if r, _ := c.ToString(mm); r != `{"a":"b"}` {
		err(t, "Converter.ToString map[string]string", r, `{"a":"b"}`)
	}
}

type s int

func (ø s) String() string { return "special" }

type ii int

func (ø ii) String() string { return "2" }

var toStringTests = map[interface{}]string{
	2:      "2",
	2.3:    "2.3",
	"hiho": "hiho",
	s(2):   "special",
	true:   "true",
}

var toIntTests = map[interface{}]int{
	2:     2,
	"2":   2,
	2.0:   2,
	ii(3): 2,
}

var toBoolTests = map[interface{}]bool{
	true:    true,
	false:   false,
	"true":  true,
	"false": false,
}
var toFloatTests = map[interface{}]float64{
	2.3:   2.3,
	"2.3": 2.3,
	2:     2.0,
	"2":   2.0,
}
var toTimeTests = map[interface{}]*time.Time{}

func init() {
	_ = fmt.Sprintf
	ti, _ := time.Parse(time.RFC3339, "2010-01-26T18:53:18+01:00")
	toStringTests[&ti] = "2010-01-26T18:53:18+01:00"
	toTimeTests[&ti] = &ti
	toTimeTests["2010-01-26T18:53:18+01:00"] = &ti

}

func TestToString(t *testing.T) {
	for in, out := range toStringTests {
		if r := ToString(in); r != out {
			err(t, "ToString", r, out)
		}
	}
}

func TestMapToString(t *testing.T) {
	m := map[string]int{"a": 2}
	s := ToString(m)

	if s != `{"a":2}` {
		err(t, "IntMapToString", s, `{"a":2}`)
	}

	ms := map[string]string{"a": "2"}
	s = ToString(ms)

	if s != `{"a":"2"}` {
		err(t, "StringMapToString", s, `{"a":"2"}`)
	}

	mi := map[string]interface{}{"a": 2}
	s = ToString(mi)

	if s != `{"a":2}` {
		err(t, "InterfaceMapToString", s, `{"a":2}`)
	}
}

func TestToInt(t *testing.T) {
	for in, out := range toIntTests {
		if r, e := ToInt(in); r != out || e != nil {
			if e != nil {
				t.Error(e)
			} else {
				err(t, "ToInt", r, out)
			}

		}
	}
}

func TestToIntArr(t *testing.T) {
	if r, _ := ToIntArr("3,4,5"); r[0] != 3 || r[1] != 4 || r[2] != 5 {
		err(t, "ToIntArr", r, []int{3, 4, 5})
	}
}

func TestToFloatArr(t *testing.T) {
	if r, _ := ToFloatArr("3.5,4.5,5.5"); r[0] != 3.5 || r[1] != 4.5 || r[2] != 5.5 {
		err(t, "ToFloatArr", r, []float64{3.5, 4.5, 5.5})
	}
}

func TestToStringArr(t *testing.T) {
	if r := ToStringArr("a,b,c"); r[0] != "a" || r[1] != "b" || r[2] != "c" {
		err(t, "ToString", r, []string{"a", "b", "c"})
	}
}

func TestToBoolArr(t *testing.T) {
	if r, _ := ToBoolArr("true,false"); r[0] != true || r[1] != false {
		err(t, "ToBoolArr", r, []bool{true, false})
	}
}

func TestToTimeArr(t *testing.T) {
	if r, _ := ToTimeArr("2010-01-26T18:53:18+01:00,2009-01-26T18:53:18+01:00"); r[0].Format(time.RFC3339) != "2010-01-26T18:53:18+01:00" || r[1].Format(time.RFC3339) != "2009-01-26T18:53:18+01:00" {
		err(t, "ToTimeArr", r, []string{"2010-01-26T18:53:18+01:00", "2009-01-26T18:53:18+01:00"})
	}
}

func TestToBool(t *testing.T) {
	for in, out := range toBoolTests {
		if r, _ := ToBool(in); r != out {
			err(t, "ToBool", r, out)
		}
	}
}

func TestToFloat(t *testing.T) {
	for in, out := range toFloatTests {
		if r, _ := ToFloat(in); r != out {
			err(t, "ToFloat", r, out)
		}
	}
}

func TestToTime(t *testing.T) {
	for in, out := range toTimeTests {
		if r, _ := ToTime(in); r.Format(time.RFC3339) != out.Format(time.RFC3339) {
			err(t, "ToTime", r, out)
		}
	}
}

func TestToStringMap(t *testing.T) {
	s := `{"a": 2,"b": "3","c": 4.5}`
	m, _ := ToStringMap(s)

	// warning: json renders all numbers to float64
	if m["a"] != 2.0 || m["b"] != "3" || m["c"] != 4.5 {
		err(t, "ToStringMap", m, map[string]interface{}{"a": 2.0, "b": "3", "c": 4.5})
	}
}

func TestToIntMap(t *testing.T) {
	s := []interface{}{2, "3", 4.5}
	m, _ := ToIntMap(s)

	if m[0] != 2 || m[1] != "3" || m[2] != 4.5 {
		err(t, "ToIntMap []interface{} input", m, map[int]interface{}{0: 2, 1: "3", 2: 4.5})
	}

	ss := `[2,"3",4.5]`
	m, _ = ToIntMap(ss)
	// warning: json renders all numbers to float64
	if m[0] != 2.0 || m[1] != "3" || m[2] != 4.5 {
		err(t, "ToIntMap json string input", m, map[int]interface{}{0: 2.0, 1: "3", 2: 4.5})
	}
}

func TestToArr(t *testing.T) {
	s := `[2,"3",4.5]`
	m, _ := ToArr(s)

	// warning: json renders all numbers to float64
	if m[0] != 2.0 || m[1] != "3" || m[2] != 4.5 {
		err(t, "ToStringMap", m, []interface{}{2.0, "3", 4.5})
	}
}

func TestToStringStringMap(t *testing.T) {
	s := `{"a": 2,"b": "3","c": 4.5}`
	m, _ := ToStringStringMap(s)

	if m["a"] != "2" || m["b"] != "3" || m["c"] != "4.5" {
		err(t, "ToStringMap", m, map[string]string{"a": "2", "b": "3", "c": "4.5"})
	}
}
*/
