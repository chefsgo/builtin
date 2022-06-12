package builtin

import (
	"fmt"
	"html/template"
	"strconv"
	"strings"
	"time"

	. "github.com/chefsgo/base"
	"github.com/chefsgo/chef"
	"github.com/chefsgo/view"
)

func init() {

	chef.Register("hashid", view.Helper{
		Name: "hashid", Desc: "hashid",
		Action: func(val int64) string {
			id, _ := chef.EncryptDIGIT(val)
			return id
		},
	}, false)

	chef.Register("percent", view.Helper{
		Name: "百分比", Desc: "百分比",
		Action: func(val float64) string {
			return fmt.Sprintf("%.2f", val*float64(100))
		},
	}, false)

	chef.Register("round", view.Helper{
		Name: "四舍六入", Desc: "四舍六入",
		Action: func(val float64, precisions ...Any) string {
			precision := 2

			if len(precisions) > 0 {
				if vv, ok := precisions[0].(int); ok {
					precision = vv
				} else if vv, ok := precisions[0].(int64); ok {
					precision = int(vv)
				} else if vv, ok := precisions[0].(string); ok {
					if ii, ee := strconv.ParseInt(vv, 10, 64); ee == nil {
						precision = int(ii)
					}
				}
			}

			if precision > 0 {
				format := fmt.Sprintf("%%0.%vf", precision)
				return fmt.Sprintf(format, val)
			}

			return fmt.Sprintf("%f", val)
		},
	}, false)

	chef.Register("raw", view.Helper{
		Name: "原始输出", Desc: "原始输出",
		Action: func(sss Any) template.HTML {
			if sss != nil {
				return template.HTML(fmt.Sprintf("%v", sss))
			}
			return template.HTML("")
		},
	}, false)

	chef.Register("html", view.Helper{
		Name: "输出HTML", Desc: "输出HTML，和raw一个意思",
		Action: func(sss Any) template.HTML {
			if sss != nil {
				return template.HTML(fmt.Sprintf("%v", sss))
			}
			return template.HTML("")
		},
	}, false)

	chef.Register("attr", view.Helper{
		Name: "HTML属性输出", Desc: "HTML属性输出，因为GO模板会自动转义html",
		Action: func(text Any) template.HTMLAttr {
			if text != nil {
				return template.HTMLAttr(fmt.Sprintf("%v", text))
			}
			return template.HTMLAttr("")
		},
	}, false)

	chef.Register("url", view.Helper{
		Name: "Url输出", Desc: "Url输出，因为GO模板会自动转义html",
		Action: func(text Any) template.URL {
			if text != nil {
				return template.URL(fmt.Sprintf("%v", text))
			}
			return template.URL("")
		},
	}, false)

	chef.Register("join", view.Helper{
		Name: "数组join输出", Desc: "数组join输出",
		Action: func(a Any, s string) template.HTML {
			strs := []string{}

			if a != nil {

				switch vv := a.(type) {
				case []string:
					for _, v := range vv {
						strs = append(strs, v)
					}
				case []Any:
					for _, v := range vv {
						strs = append(strs, fmt.Sprintf("%v", v))
					}
				case []int:
					for _, v := range vv {
						strs = append(strs, fmt.Sprintf("%v", v))
					}
				case []int8:
					for _, v := range vv {
						strs = append(strs, fmt.Sprintf("%v", v))
					}
				case []int16:
					for _, v := range vv {
						strs = append(strs, fmt.Sprintf("%v", v))
					}
				case []int32:
					for _, v := range vv {
						strs = append(strs, fmt.Sprintf("%v", v))
					}
				case []int64:
					for _, v := range vv {
						strs = append(strs, fmt.Sprintf("%v", v))
					}
				case []float32:
					for _, v := range vv {
						strs = append(strs, fmt.Sprintf("%v", v))
					}
				case []float64:
					for _, v := range vv {
						strs = append(strs, fmt.Sprintf("%v", v))
					}
				}
			}

			html := strings.Join(strs, s)
			return template.HTML(html)
		},
	}, false)

	chef.Register("json", view.Helper{
		Name: "json输出", Desc: "json输出",
		Action: func(data Any) template.HTML {
			if data != nil {
				bytes, err := chef.MarshalJSON(data)
				if err == nil {
					return template.HTML(string(bytes))
				}
			}
			return template.HTML("")
		},
	}, false)

	chef.Register("mapping", view.Helper{
		Name: "生成Map", Desc: "生成Map",
		Action: func(args ...Any) Map {
			m := Map{}

			k := ""
			for i, v := range args {
				if (i+1)%2 == 1 {
					k = fmt.Sprintf("%v", v)
				} else {
					m[k] = v
				}
			}

			return m
		},
	}, false)

	chef.Register("now", view.Helper{
		Name: "当前时间", Desc: "当前时间",
		Action: func() time.Time {
			return time.Now()
		},
	}, false)

	chef.Register("in", view.Helper{
		Name: "判断某个值是否在数组", Desc: "判断某个值是否在数组",
		Action: func(val Any, arrs ...Any) bool {

			strVal := fmt.Sprintf("%v", val)
			strArr := []string{}

			if len(arrs) > 1 {
				for _, vv := range arrs {
					strArr = append(strArr, fmt.Sprintf("%v", vv))
				}
			} else {
				switch vv := arrs[0].(type) {
				case []Any:
					{
						for _, v := range vv {
							strArr = append(strArr, fmt.Sprintf("%v", v))
						}
					}
				case []string:
					for _, v := range vv {
						strArr = append(strArr, v)
					}
				case []int:
					for _, v := range vv {
						strArr = append(strArr, fmt.Sprintf("%v", v))
					}
				case []int8:
					for _, v := range vv {
						strArr = append(strArr, fmt.Sprintf("%v", v))
					}
				case []int16:
					for _, v := range vv {
						strArr = append(strArr, fmt.Sprintf("%v", v))
					}
				case []int32:
					for _, v := range vv {
						strArr = append(strArr, fmt.Sprintf("%v", v))
					}
				case []int64:
					for _, v := range vv {
						strArr = append(strArr, fmt.Sprintf("%v", v))
					}
				default:
					strArr = append(strArr, fmt.Sprintf("%v", vv))
				}
			}

			for _, v := range strArr {
				if v == strVal {
					return true
				}
			}

			return false
		},
	}, false)

	chef.Register("out", view.Helper{
		Name: "输出输组某下标元素", Desc: "输出输组某下标元素",
		Action: func(arr Any, i int) string {

			strArr := []string{}

			switch vv := arr.(type) {
			case []string:
				for _, v := range vv {
					strArr = append(strArr, v)
				}
			case []int:
				for _, v := range vv {
					strArr = append(strArr, fmt.Sprintf("%v", v))
				}
			case []int8:
				for _, v := range vv {
					strArr = append(strArr, fmt.Sprintf("%v", v))
				}
			case []int16:
				for _, v := range vv {
					strArr = append(strArr, fmt.Sprintf("%v", v))
				}
			case []int32:
				for _, v := range vv {
					strArr = append(strArr, fmt.Sprintf("%v", v))
				}
			case []int64:
				for _, v := range vv {
					strArr = append(strArr, fmt.Sprintf("%v", v))
				}
			}

			if len(strArr) > i {
				return strArr[i]
			}

			return ""
		},
	}, false)

}
