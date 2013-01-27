package typeconverter

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// checks, if something is a string or Stringer
func IsString(i interface{}) bool {
	if _, ok := i.(string); ok {
		return true
	}

	if _, ok := i.(fmt.Stringer); ok {
		return true
	}

	return false
}

func ToString(i interface{}) (r string) {
	switch v := i.(type) {
	case *time.Time:
		r = v.Format(time.RFC3339)
	case []int:
		s := []string{}
		for _, i := range v {
			s = append(s, ToString(i))
		}
		r = strings.Join(s, ",")
	case []bool:
		s := []string{}
		for _, i := range v {
			s = append(s, ToString(i))
		}
		r = strings.Join(s, ",")
	case []float64:
		s := []string{}
		for _, i := range v {
			s = append(s, ToString(i))
		}
		r = strings.Join(s, ",")
	case []*time.Time:
		s := []string{}
		for _, i := range v {
			s = append(s, ToString(i))
		}
		r = strings.Join(s, ",")
	case []time.Time:
		s := []string{}
		for _, i := range v {
			s = append(s, ToString(i))
		}
		r = strings.Join(s, ",")
	case []string:
		r = strings.Join(v, ",")
	case []fmt.Stringer:
		s := []string{}
		for _, i := range v {
			s = append(s, ToString(i))
		}
		r = strings.Join(s, ",")
	default:
		if IsString(v) {
			r = fmt.Sprintf("%v", v)
		} else {
			b, _ := json.Marshal(v)
			r = string(b)
		}

	}
	return
}

func ToInt(o interface{}) (i int, err error) {
	s := ToString(o)
	ii, err := strconv.ParseInt(s, 10, 32)
	i = int(ii)
	return
}

func ToBool(o interface{}) (b bool, err error) {
	s := ToString(o)
	b, err = strconv.ParseBool(s)
	return
}

func ToFloat(o interface{}) (f float64, err error) {
	s := ToString(o)
	f, err = strconv.ParseFloat(s, 64)
	return
}

func ToTime(o interface{}) (t *time.Time, err error) {
	s := ToString(o)
	// example for time.RFC3339 format: 2006-01-02T15:04:05Z07:00
	tt, err := time.Parse(time.RFC3339, s)
	t = &tt
	return
}

func ToIntArr(o interface{}) (ia []int, err error) {
	s := ToString(o)
	a := strings.Split(s, ",")
	ia = []int{}
	for _, ii := range a {
		iii, e := strconv.ParseInt(ii, 10, 32)
		if e != nil {
			err = e
			return
		}
		ia = append(ia, int(iii))
	}
	return
}

func ToBoolArr(o interface{}) (ia []bool, err error) {
	s := ToString(o)
	a := strings.Split(s, ",")
	ia = []bool{}
	for _, ii := range a {
		b, e := strconv.ParseBool(ii)
		if e != nil {
			err = e
			return
		}
		ia = append(ia, b)
	}
	return
}

func ToFloatArr(o interface{}) (ia []float64, err error) {
	s := ToString(o)
	a := strings.Split(s, ",")
	ia = []float64{}
	for _, ii := range a {
		iii, e := strconv.ParseFloat(ii, 32)
		if e != nil {
			err = e
			return
		}
		ia = append(ia, float64(iii))
	}
	return
}

func ToStringArr(o interface{}) (ia []string) {
	s := ToString(o)
	a := strings.Split(s, ",")
	ia = []string{}
	for _, ii := range a {
		iii := ToString(ii)
		ia = append(ia, iii)
	}
	return
}

func ToTimeArr(o interface{}) (ia []*time.Time, err error) {
	s := ToString(o)
	a := strings.Split(s, ",")
	ia = []*time.Time{}
	for _, ii := range a {
		iii, e := time.Parse(time.RFC3339, ii)
		if e != nil {
			err = e
			return
		}
		ia = append(ia, &iii)
	}
	return
}

func ToArr(o interface{}) (ia []interface{}, err error) {
	ia = []interface{}{}
	s := ToString(o)
	err = json.Unmarshal([]byte(s), &ia)
	return
}

func ToStringMap(o interface{}) (m map[string]interface{}, err error) {
	m = map[string]interface{}{}
	s := ToString(o)
	err = json.Unmarshal([]byte(s), &m)
	return
}

func ToIntMap(o interface{}) (m map[int]interface{}, err error) {
	m = map[int]interface{}{}
	switch t := o.(type) {
	case []interface{}:
		for k, v := range t {
			m[k] = v
		}
	default:
		var a []interface{}
		a, err = ToArr(o)
		for k, v := range a {
			m[k] = v
		}
	}
	return
}

func ToStringStringMap(o interface{}) (m map[string]string, err error) {
	m = map[string]string{}
	s, err := ToStringMap(o)
	for k, v := range s {
		m[k] = ToString(v)
	}
	return
}
