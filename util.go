package builtin

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"strconv"

	. "github.com/chefsgo/base"
	"github.com/chefsgo/chef"
)

//密码加密格式
func password(str string) string {
	return sha1String(str)
}

//sha1加密
func sha1String(str string) string {
	sha1Ctx := sha1.New()
	sha1Ctx.Write([]byte(str))
	cipherStr := sha1Ctx.Sum(nil)
	return hex.EncodeToString(cipherStr)
}

//sha1加密文件
func sha1File(file string) string {
	if f, e := os.Open(file); e == nil {
		defer f.Close()

		h := sha1.New()
		if _, e := io.Copy(h, f); e == nil {
			return fmt.Sprintf("%x", h.Sum(nil))
		}
	}
	return ""
}

//md5加密
func md5String(str string) string {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(str))
	cipherStr := md5Ctx.Sum(nil)
	return hex.EncodeToString(cipherStr)
}

//md5加密文件
func md5File(file string) string {
	if f, e := os.Open(file); e == nil {
		defer f.Close()

		h := md5.New()
		if _, e := io.Copy(h, f); e == nil {
			return fmt.Sprintf("%x", h.Sum(nil))
		}
	}
	return ""
}

// func dataTargetConfig(extends ...Vars) Vars {
// 	config := Vars{
// 		"data":  chef.Define("string", false, nil, "数据库+表"),
// 		"base":  chef.Define("string", false, nil, "数据库"),
// 		"table": chef.Define("string", false, nil, "表"),
// 		"view":  chef.Define("string", false, nil, "视图"),
// 	}
// 	return chef.VarsExtend(config, extends...)
// }

// func dataTargetTable(args Map) (string, string) {
// 	base, table := "", ""
// 	if args["data"] != nil {
// 		if vv, ok := args["data"].(string); ok {
// 			i := strings.Index(vv, ".")
// 			base = vv[:i]
// 			table = vv[i+1:]
// 		}
// 	} else if args["base"] != nil && args["table"] != nil {
// 		if vv, ok := args["base"].(string); ok {
// 			base = vv
// 		}
// 		if vv, ok := args["table"].(string); ok {
// 			table = vv
// 		}
// 	}
// 	return base, table

// }

// func dataConfigFilters(name string, bases ...string) (groups, arrays []string, ranges Map, orders Map, querys []string) {
// 	base := ""
// 	if len(bases) > 0 {
// 		base = bases[0]
// 	}

// 	groups = make([]string, 0)
// 	arrays = make([]string, 0)
// 	ranges = Map{}
// 	orders = Map{}
// 	querys = make([]string, 0)

// 	fields := chef.Fields(name, nil)
// 	if len(fields) == 0 && base != "" {
// 		fields = chef.Fields(base+"."+name, nil)
// 	}

// 	for key, field := range fields {
// 		if field.Setting != nil {

// 			if field.Setting["group"] != nil {
// 				groups = append(groups, key)
// 			}
// 			if field.Setting["array"] != nil {
// 				arrays = append(arrays, key)
// 			}
// 			if field.Setting["range"] != nil {
// 				ranges[key] = true
// 			}
// 			if field.Setting["order"] != nil {
// 				orders[key] = true
// 			}
// 			if field.Setting["query"] != nil {
// 				querys = append(querys, key)
// 			}
// 		}
// 	}
// 	return
// }

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
		d, e := chef.MarshalJSON(v)
		if e == nil {
			sv = string(d)
		} else {
			sv = "{}"
		}
	case []Map:
		d, e := chef.MarshalJSON(v)
		if e == nil {
			sv = string(d)
		} else {
			sv = "[]"
		}
	case []int, []int8, []int16, []int32, []int64, []float32, []float64, []string, []bool, []Any:
		d, e := chef.MarshalJSON(v)
		if e == nil {
			sv = string(d)
		} else {
			sv = "[]"
		}
	default:
		sv = fmt.Sprintf("%v", v)
	}

	return sv
}
