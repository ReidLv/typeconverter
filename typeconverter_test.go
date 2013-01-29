package typeconverter

import (
	"fmt"
	"testing"
	"time"
)

func err(t *testing.T, msg string, is interface{}, shouldbe interface{}) {
	t.Errorf(msg+": is %#v, should be %#v\n", is, shouldbe)
}

type Special string

func (ø Special) Int() int { return 42 }

func dispatchToSpecial(out interface{}, in interface{}) (err error) {
	*out.(*Special) = Special(in.(Stringer).String())
	return
}

// convert default to different types to get the defaults
func ExampleConvert_default() {
	var d Default
	var s string
	Convert(d, &s)
	fmt.Printf("default string: %#v\n", s)
	var i int
	Convert(d, &i)
	fmt.Printf("default int: %#v\n", i)
	var f float64
	Convert(d, &f)
	fmt.Printf("default float: %#v\n", f)
	var j Json
	Convert(d, &j)
	fmt.Printf("default json: %#v\n", j)
	var t time.Time
	Convert(d, &t)
	fmt.Printf("default time: %#v\n", Time(t).String())
	// Output: default string: ""
	// default int: 0
	// default float: 0
	// default json: "{}"
	// default time: "0001-01-01T00:00:00Z"
}

// convert time to string
func ExampleConvert() {
	var s string
	t1, _ := time.Parse(time.RFC3339, "2011-01-26T18:53:18+01:00")
	Convert(t1, &s)
	fmt.Println(s)
	// Output: 2011-01-26T18:53:18+01:00
}

// define your own type and integrate it to the converter
func ExampleNew_ownType() {
	/* we defined
	type Special string

	func (ø Special) Int() int {
		return 42
	}

	*/
	c := New()
	sp := Special("")
	c.Output.SetHandler(&sp, func(out interface{}, in interface{}) (err error) {
		*out.(*Special) = Special(in.(Stringer).String() + " as special")
		return
	})

	s := Special("")
	var r int
	c.Convert(s, &r)
	fmt.Printf("to int: %v\n", r)

	c.Convert(float64(4.5), &s)
	fmt.Printf("to special: %v\n", s)

	var t time.Time
	e := c.Convert(Special(""), &t)
	fmt.Println(e)
	// Output: to int: 42
	// to special: 4.5 as special
	// interface conversion: typeconverter.Special is not typeconverter.Timer: missing method Time
}

// overwrite builtin conversions
func ExampleNew_overwrite() {
	c := New()
	// if input should be transformed to string
	// change the output and add " was the answer" to normal string conversion
	var s string
	c.Output.SetHandler(&s,
		func(out interface{}, in interface{}) (err error) {
			*out.(*string) = in.(Stringer).String() + " was the answer"
			return
		})
	c.Convert(42, &s)
	fmt.Println(s)
	// Output: 42 was the answer
}

var _ = fmt.Errorf
var ti, _ = time.Parse(time.RFC3339, "2011-01-26T18:53:18+01:00")

var toIntTests = map[interface{}]int{
	1:            1,
	int64(2):     2,
	float64(3.0): 3,
	float32(3.0): 3,
	Json(`3.0`):  3,
	Json(`3`):    3,
	`1`:          1,
	`1.0`:        1,
	ti:           1296064398,
}

func TestToInt(t *testing.T) {
	for in, out := range toIntTests {
		var r int
		if Convert(in, &r); r != out {
			err(t, "Convert to int", r, out)
		}
	}
}

var toFloatTests = map[interface{}]float64{
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
	for in, out := range toFloatTests {
		var r float64
		if Convert(in, &r); r != out {
			err(t, "ToFloat", r, out)
		}
	}
}

var toBoolTests = map[interface{}]bool{
	true:          true,
	false:         false,
	Json(`true`):  true,
	Json(`false`): false,
	`true`:        true,
	`false`:       false,
}

func TestToBool(t *testing.T) {
	var r bool
	for in, out := range toBoolTests {
		if Convert(in, &r); r != out {
			err(t, "ToBool2", r, out)
		}
	}
}

var toStringTests = map[interface{}]string{
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
	for in, out := range toStringTests {
		var r string
		if Convert(in, &r); r != out {
			err(t, "ToString", r, out)
		}
	}

	m := map[string]interface{}{"a": 3}
	out := `{"a":3}`
	var r string
	if Convert(m, &r); r != out {
		err(t, "ToString", r, out)
	}

	a := []interface{}{"a", 3, 4.5}
	out = `["a",3,4.5]`
	if Convert(a, &r); r != out {
		err(t, "ToString", r, out)
	}

}

var toArrayTests = map[interface{}][]interface{}{
	Json(`["a",4]`): []interface{}{"a", 4},
}

func TestToArray(t *testing.T) {
	for in, out := range toArrayTests {
		var r []interface{}
		if Convert(in, &r); r[0].(string) != out[0].(string) || toInt(r[1].(float64)) != out[1].(int) {
			err(t, "ToArray", r, out)
		}
	}

	out := []interface{}{"a", 3}
	var r []interface{}
	if Convert(out, &r); r[0].(string) != out[0].(string) || r[1].(int) != out[1].(int) {
		err(t, "ToArray", r, out)
	}
}

var toMapTests = map[interface{}]map[string]interface{}{
	Json(`{"a":"b"}`): map[string]interface{}{"a": "b"},
}

func TestToMap(t *testing.T) {
	for in, out := range toMapTests {
		var r map[string]interface{}
		if Convert(in, &r); r["a"] != out["a"] {
			err(t, "ToMap", r, out)
		}
	}

	out := map[string]interface{}{"a": "b"}
	var r map[string]interface{}
	if Convert(out, &r); r["a"] != out["a"] {
		err(t, "ToMap", r, out)
	}

}

var toTimeTests = map[interface{}]string{
	int32(1010000000):   "Wed Jan  2 19:33:20 2002",
	int64(1010000000):   "Wed Jan  2 19:33:20 2002",
	float32(1010000000): "Sun Sep  9 01:46:40 2001",
	float64(1010000000): "Sun Sep  9 01:46:40 2001",
	ti:                  `Wed Jan 26 17:53:18 2011`,
	Json(`"2011-01-26T18:53:18+01:00"`): `Wed Jan 26 17:53:18 2011`,
	`2011-01-26T18:53:18+01:00`:         `Wed Jan 26 17:53:18 2011`,
}

func TestToTime(t *testing.T) {
	for in, out := range toTimeTests {
		var r time.Time
		r = r.UTC()
		Convert(in, &r)
		//r.UTC()
		if r.UTC().Format(time.ANSIC) != out {
			err(t, "ToTime", r.UTC().Format(time.ANSIC), out)
		}
	}
}

var toJsonTests = map[interface{}]string{
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
	for in, out := range toJsonTests {
		var r Json
		if Convert(in, &r); string(r) != out {
			err(t, "ToJson", string(r), out)
		}
	}

	m := map[string]interface{}{"a": 3}
	out := `{"a":3}`
	var r Json
	if Convert(m, &r); string(r) != out {
		err(t, "ToJson", string(r), out)
	}

	a := []interface{}{"a", 3, 4.5}
	out = `["a",3,4.5]`
	if Convert(a, &r); string(r) != out {
		err(t, "ToJson", string(r), out)
	}

}
