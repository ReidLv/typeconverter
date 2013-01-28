package typeconverter

import (
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

var toArray = map[interface{}][]interface{}{
	Json(`["a",4]`): []interface{}{"a", 4},
}

func TestToArray(t *testing.T) {
	for in, out := range toArray {
		if r := Convert(in).ToArray(); r[0].(string) != out[0].(string) || ToInt(r[1].(float64)) != out[1].(int) {
			err(t, "ToArray", r, out)
		}
	}

	out := []interface{}{"a", 3}
	if r := Convert(out).ToArray(); r[0].(string) != out[0].(string) || r[1].(int) != out[1].(int) {
		err(t, "ToArray", r, out)
	}
}

var toMap = map[interface{}]map[string]interface{}{
	Json(`{"a":"b"}`): map[string]interface{}{"a": "b"},
}

func TestToMap(t *testing.T) {
	for in, out := range toMap {
		if r := Convert(in).ToMap(); r["a"] != out["a"] {
			err(t, "ToMap", r["a"], out["a"])
		}
	}

	out := map[string]interface{}{"a": "b"}
	if r := Convert(out).ToMap(); r["a"] != out["a"] {
		err(t, "ToMap", r["a"], out["a"])
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
		if r := Time(Convert(in).ToTime()).String(); r != out {
			err(t, "ToTime", r, out)
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
