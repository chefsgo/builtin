package builtin

import (
	"fmt"
	"strconv"

	. "github.com/chefsgo/base"
	"github.com/chefsgo/chef"
)

func anyToString(val Any) string {
	sv := ""
	switch v := val.(type) {
	case string:
		sv = v
	case int:
		sv = strconv.Itoa(v)
	case int64:
		sv = strconv.FormatInt(v, 10)
	case bool:
		sv = strconv.FormatBool(v)
	case Map:
		d, e := chef.JSONEncode(v)
		if e == nil {
			sv = string(d.([]byte))
		} else {
			sv = "{}"
		}
	case []Map:
		d, e := chef.JSONEncode(v)
		if e == nil {
			sv = string(d.([]byte))
		} else {
			sv = "[]"
		}
	case []int, []int8, []int16, []int32, []int64, []float32, []float64, []string, []bool, []Any:
		d, e := chef.JSONEncode(v)
		if e == nil {
			sv = string(d.([]byte))
		} else {
			sv = "[]"
		}
	default:
		sv = fmt.Sprintf("%v", v)
	}

	return sv
}
