package typeconverter

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

/*
	basic types are:

		- int
		- bool
		- float32
		- *time.Time
		- string/Stringer
		- []bool
		- []int
		- []float32
		- []string
		- []*time.Time
*/

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
	case []float32:
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
		r = fmt.Sprintf("%v", v)
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

func ToFloat(o interface{}) (f float32, err error) {
	s := ToString(o)
	ff, err := strconv.ParseFloat(s, 32)
	f = float32(ff)
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

func ToFloatArr(o interface{}) (ia []float32, err error) {
	s := ToString(o)
	a := strings.Split(s, ",")
	ia = []float32{}
	for _, ii := range a {
		iii, e := strconv.ParseFloat(ii, 32)
		if e != nil {
			err = e
			return
		}
		ia = append(ia, float32(iii))
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
