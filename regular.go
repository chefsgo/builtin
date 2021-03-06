package builtin

import (
	. "github.com/chefsgo/base"
	"github.com/chefsgo/chef"
)

func init() {

	regulars := Map{
		"password": `^[0-9A-Fa-f]{40}$`,
		"number":   `^[0-9]+$`,
		"float":    `^[0-9]+$`,
		"date": []string{
			`^(\d{4})(\d{2})(\d{2})$`,
			`^(\d{4})-(\d{2})-(\d{2})$`,
			`^(\d{10, 15})$`,
		},
		"datetime": []string{
			`^(\d{4})-(\d{2})-(\d{2})$`,
			`^(\d{4})-(\d{2})-(\d{2}) (\d{2}):(\d{2}):(\d{2})$`,
			`^(\d{4})-(\d{2})-(\d{2})T(\d{2}):(\d{2}):(\d{2})Z$`,
			`^(\d{4})-(\d{2})-(\d{2}) (\d{2}):(\d{2}):(\d{2})\.(\d{3})$`,
			`^(\d{4})-(\d{2})-(\d{2})T(\d{2}):(\d{2}):(\d{2})\.(\d{3})Z$`,
			`^(\d{4})-(\d{2})-(\d{2})T(\d{2}):(\d{2}):(\d{2})\.(\d{6})\+(\d{2}):(\d{2})$`,
			//2017-03-18T17:20:26.689501+08:00
			`^(\d{10, 15})$`,
		},
		"mobile": `^1[0-9]{10}$`,
		"idcard": `(^\d{15}$)|(^\d{18}$)|(^\d{17}(\d|X|x)$)`,
		"email":  `^[0-9A-Za-z][_.0-9A-Za-z-]{0,31}@([0-9A-Za-z][0-9A-Za-z-]{0,30}[0-9A-Za-z]\.){1,4}[A-Za-z]{2,4}$`,
		"array": []string{
			"^asdfasfasf$",
		},
		"test": `^test$`,
	}

	for k, v := range regulars {
		if vv, ok := v.(string); ok {
			chef.Register(k, chef.Regular{vv}, false)
		}
		if vvs, ok := v.([]string); ok {
			chef.Register(k, chef.Regular(vvs), false)
		}
	}

}
