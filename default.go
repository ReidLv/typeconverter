package typeconverter

import (
	"time"
)

type Default int

func (ø Default) String() string              { return "" }
func (ø Default) Int() int                    { return 0 }
func (ø Default) Float() float64              { return float64(0) }
func (ø Default) Time() time.Time             { return time.Now() }
func (ø Default) Json() string                { return "{}" }
func (ø Default) Array() []interface{}        { return []interface{}{} }
func (ø Default) Bool() bool                  { return false }
func (ø Default) Map() map[string]interface{} { return map[string]interface{}{} }
