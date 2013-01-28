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

func (ø Special) Int() int {
	return 42
}

func dispatchToSpecial(out interface{}, in interface{}) (err error) {
	*out.(*Special) = Special(in.(Stringer).String())
	return
}

// convert time to string
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

func ExampleNew_ownType() {
	/* we defined
	type Special string

	func (ø Special) Int() int {
		return 42
	}

	*/
	c := New()
	sp := Special("")
	c.Input.AddType(sp)
	c.Output.AddType(&sp)
	c.Output.SetHandler("*typeconverter.Special", func(out interface{}, in interface{}) (err error) {
		*out.(*Special) = Special(in.(Stringer).String())
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
	// to special: 4.5
	// interface conversion: typeconverter.Special is not typeconverter.Timer: missing method Time
}

func ExampleNew_overwrite() {
	c := New()
	// if input should be transformed to string
	// change the output and add " was the answer" to normal string conversion
	c.Output.SetHandler("*string",
		func(out interface{}, in interface{}) (err error) {
			*out.(*string) = in.(Stringer).String() + " was the answer"
			return
		})
	var s string
	c.Convert(42, &s)
	fmt.Println(s)
	// Output: 42 was the answer
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
	ti:           1296064398,
}

/*
func TestError(t *testing.T) {
	var r time.Time
	if e := Convert(Special(""), &r); e == nil {
		err(t, "Convert error", false, true)
	}
}
*/

func TestToInt(t *testing.T) {
	for in, out := range toInt {
		var r int
		if Convert(in, &r); r != out {
			err(t, "Convert to int", r, out)
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
		var r float64
		if Convert(in, &r); r != out {
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
	var r bool
	for in, out := range toBool {
		if Convert(in, &r); r != out {
			err(t, "ToBool2", r, out)
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

var toArray = map[interface{}][]interface{}{
	Json(`["a",4]`): []interface{}{"a", 4},
}

func TestToArray(t *testing.T) {
	for in, out := range toArray {
		var r []interface{}
		if Convert(in, &r); r[0].(string) != out[0].(string) || toInt32(r[1].(float64)) != out[1].(int) {
			err(t, "ToArray", r, out)
		}
	}

	out := []interface{}{"a", 3}
	var r []interface{}
	if Convert(out, &r); r[0].(string) != out[0].(string) || r[1].(int) != out[1].(int) {
		err(t, "ToArray", r, out)
	}
}

var toMap = map[interface{}]map[string]interface{}{
	Json(`{"a":"b"}`): map[string]interface{}{"a": "b"},
}

func TestToMap(t *testing.T) {
	for in, out := range toMap {
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

var toTime = map[interface{}]string{
	1010000000:          "2002-01-02T20:33:20+01:00",
	int64(1010000000):   "2002-01-02T20:33:20+01:00",
	float32(1010000000): "2001-09-09T03:46:40+02:00",
	float64(1010000000): "2001-09-09T03:46:40+02:00",
	ti:                  `2011-01-26T18:53:18+01:00`,
	Json(`"2011-01-26T18:53:18+01:00"`): `2011-01-26T18:53:18+01:00`,
	`2011-01-26T18:53:18+01:00`:         `2011-01-26T18:53:18+01:00`,
}

func TestToTime(t *testing.T) {
	for in, out := range toTime {
		var r time.Time
		if Convert(in, &r); Time(r).String() != out {
			err(t, "ToTime", Time(r).String(), out)
		}
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
