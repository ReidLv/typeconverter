package typeconverter

import (
	"encoding/json"
	"fmt"
	"github.com/metakeule/dispatch"
	"time"
)

type Converter struct {
	Output *dispatch.Dispatcher
	Input  *dispatch.Dispatcher
}

func (ø *Converter) Convert(from interface{}, to interface{}) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%s", r)
		}
	}()
	ø.Input.Dispatch(from, to)
	return
}

var basicConvert = New()

func Convert(from interface{}, to interface{}) (err error) { return basicConvert.Convert(from, to) }

func New() (ø *Converter) {
	ø = &Converter{dispatch.New(), dispatch.New()}

	ø.Input.AddType(Int(0))
	ø.Input.AddType(Int64(0))
	ø.Input.AddType(Float(0.0))
	ø.Input.AddType(Float32(0.0))
	ø.Input.AddType(Bool(true))
	ø.Input.AddType(String(""))
	ø.Input.AddType(Time(time.Time{}))
	ø.Input.AddType(Json(""))
	ø.Input.AddType(Map(map[string]interface{}{}))
	ø.Input.AddType(Array([]interface{}{}))
	ø.Input.AddType(Default(0))

	ø.Input.AddType(int(0))
	ø.Input.AddType(int64(0))
	ø.Input.AddType(int32(0))
	ø.Input.AddType(float64(0.0))
	ø.Input.AddType(float32(0.0))
	ø.Input.AddType(bool(true))
	ø.Input.AddType(string(""))
	ø.Input.AddType(time.Time{})
	ø.Input.AddType(map[string]interface{}{})
	ø.Input.AddType([]interface{}{})

	inSwitcher := func(from interface{}, to interface{}) (err error) {
		switch t := from.(type) {
		case int:
			ø.Output.Dispatch(to, Int(t))
		case int32:
			ø.Output.Dispatch(to, Int(t))
		case int64:
			ø.Output.Dispatch(to, Int64(t))
		case float64:
			ø.Output.Dispatch(to, Float(t))
		case float32:
			ø.Output.Dispatch(to, Float32(t))
		case string:
			ø.Output.Dispatch(to, String(t))
		case bool:
			ø.Output.Dispatch(to, Bool(t))
		case time.Time:
			ø.Output.Dispatch(to, Time(t))
		case map[string]interface{}:
			ø.Output.Dispatch(to, Map(t))
		case []interface{}:
			ø.Output.Dispatch(to, Array(t))
		default:
			ø.Output.Dispatch(to, t)
		}
		return
	}

	ø.Input.SetHandler("int", inSwitcher)
	ø.Input.SetHandler("int32", inSwitcher)
	ø.Input.SetHandler("int64", inSwitcher)
	ø.Input.SetHandler("float64", inSwitcher)
	ø.Input.SetHandler("float32", inSwitcher)
	ø.Input.SetHandler("string", inSwitcher)
	ø.Input.SetHandler("bool", inSwitcher)
	ø.Input.SetHandler("time.Time", inSwitcher)
	ø.Input.SetHandler("map[string]interface {}", inSwitcher)
	ø.Input.SetHandler("[]interface {}", inSwitcher)

	ø.Input.AddFallback(func(in interface{}, out interface{}) (didHandle bool, err error) {
		didHandle = true
		if ø.Input.HasTypeForInstance(in) {
			// make a default dispatch for known types
			ø.Output.Dispatch(out, in)
			return
		}
		if IsString(in) {
			ø.Output.Dispatch(out, String(fmt.Sprintf("%s", in)))
		} else {
			b, err := json.Marshal(in)
			if err != nil {
				panic("can't convert " + fmt.Sprintf("%#v", in) + " to anything reasonable or jsonable")
			}
			ø.Output.Dispatch(out, String(string(b)))
		}
		return
	})

	i := int(0)
	ø.Output.AddType(&i)
	i1 := int32(0)
	ø.Output.AddType(&i1)
	i2 := int64(0)
	ø.Output.AddType(&i2)
	fl := float64(0)
	ø.Output.AddType(&fl)
	fl2 := float32(0)
	ø.Output.AddType(&fl2)
	st := string("")
	ø.Output.AddType(&st)
	js := Json("")
	ø.Output.AddType(&js)
	t := bool(true)
	ø.Output.AddType(&t)
	tm := time.Time{}
	ø.Output.AddType(&tm)
	mp := map[string]interface{}{}
	ø.Output.AddType(&mp)
	arr := []interface{}{}
	ø.Output.AddType(&arr)

	outSwitcher := func(out interface{}, in interface{}) (err error) {
		switch t := out.(type) {
		case *bool:
			*out.(*bool) = in.(Booler).Bool()
		case *int:
			*out.(*int) = in.(Inter).Int()
		case *int64:
			*out.(*int64) = int64(in.(Inter).Int())
		case *string:
			*out.(*string) = in.(Stringer).String()
		case *float64:
			*out.(*float64) = in.(Floater).Float()
		case *float32:
			*out.(*float32) = float32(in.(Floater).Float())
		case *time.Time:
			*out.(*time.Time) = in.(Timer).Time()
		case *Json:
			*out.(*Json) = Json(in.(Jsoner).Json())
		case *map[string]interface{}:
			*out.(*map[string]interface{}) = in.(Mapper).Map()
		case *[]interface{}:
			*out.(*[]interface{}) = in.(Arrayer).Array()
		default:
			return fmt.Errorf("can't convert to %#v: no converter found", t)
		}
		return
	}

	ø.Output.SetHandler("*bool", outSwitcher)
	ø.Output.SetHandler("*int", outSwitcher)
	ø.Output.SetHandler("*int32", outSwitcher)
	ø.Output.SetHandler("*int64", outSwitcher)
	ø.Output.SetHandler("*string", outSwitcher)
	ø.Output.SetHandler("*float64", outSwitcher)
	ø.Output.SetHandler("*float32", outSwitcher)
	ø.Output.SetHandler("*time.Time", outSwitcher)
	ø.Output.SetHandler("*typeconverter.Json", outSwitcher)
	ø.Output.SetHandler("*map[string]interface {}", outSwitcher)
	ø.Output.SetHandler("*[]interface {}", outSwitcher)
	return
}
