package builtin

import (
	"bytes"
	"encoding/base64"
	"encoding/gob"
	"encoding/xml"
	"errors"
	"strconv"
	"strings"
	"time"

	jsoniter "github.com/json-iterator/go"
	"github.com/speps/go-hashids/v2"

	. "github.com/chefsgo/base"
	"github.com/chefsgo/chef"
)

var (
	errInvalidData = errors.New("Invalid data.")

	codecConfig = chef.CodecConfig()

	jsonCodec jsoniter.API = jsoniter.ConfigCompatibleWithStandardLibrary

	textCoder     = base64.NewEncoding(codecConfig.Strings)
	digitCoder, _ = hashids.NewWithData(&hashids.HashIDData{
		Alphabet: codecConfig.Numbers, Salt: codecConfig.Salt, MinLength: codecConfig.Length,
	})
)

func init() {

	gob.Register(time.Now())
	gob.Register(Map{})
	gob.Register([]Map{})

	chef.Register("json", chef.Codec{
		Name: "JSON编解码", Desc: "JSON编解码",
		Encode: func(value Any) (Any, error) {
			return jsonCodec.Marshal(value)
		},
		Decode: func(data Any, value Any) (Any, error) {
			if bytes, ok := data.([]byte); ok {
				err := jsonCodec.Unmarshal(bytes, value)
				if err != nil {
					return nil, err
				}
				return value, nil
			}
			return nil, errInvalidData
		},
	}, false)

	chef.Register("xml", chef.Codec{
		Name: "XML编解码", Desc: "XML编解码",
		Encode: func(value Any) (Any, error) {
			return xml.Marshal(value)
		},
		Decode: func(data Any, value Any) (Any, error) {
			if dataBytes, ok := data.([]byte); ok {
				err := xml.Unmarshal(dataBytes, value)
				if err != nil {
					return nil, err
				}
				return value, nil
			}
			return nil, errInvalidData
		},
	}, false)

	chef.Register("gob", chef.Codec{
		Name: "GOB编解码", Desc: "GOB编解码",
		Encode: func(value Any) (Any, error) {
			var buffer bytes.Buffer
			encoder := gob.NewEncoder(&buffer)
			err := encoder.Encode(value)
			if err != nil {
				return nil, err
			}
			return buffer.Bytes(), nil
		},
		Decode: func(data Any, value Any) (Any, error) {
			if dataBytes, ok := data.([]byte); ok {
				buffer := bytes.NewReader(dataBytes)
				decoder := gob.NewDecoder(buffer)
				err := decoder.Decode(value)
				if err != nil {
					return nil, err
				}

				return value, nil
			}
			return nil, errInvalidData
		},
	}, false)

	chef.Register("base64", chef.Codec{
		Name: "BASE64加解密", Desc: "BASE64加解密",
		Encode: func(value Any) (Any, error) {
			text := anyToString(value)
			return base64.StdEncoding.EncodeToString([]byte(text)), nil
		},
		Decode: func(data Any, value Any) (Any, error) {
			text := anyToString(value)
			bytes, err := base64.StdEncoding.DecodeString(text)
			if err != nil {
				return "", err
			}
			return string(bytes), nil
		},
	}, false)

	chef.Register("string", chef.Codec{
		Name: "文本加密", Desc: "文本加密，自定义字符表的base64编码，字典：" + codecConfig.Strings,
		Encode: func(value Any) (Any, error) {
			text := anyToString(value)
			return textCoder.EncodeToString([]byte(text)), nil
		},
		Decode: func(data Any, value Any) (Any, error) {
			text := anyToString(value)
			bytes, err := textCoder.DecodeString(text)
			if err != nil {
				return nil, err
			}
			return string(bytes), nil
		},
	}, false)
	chef.Register("strings", chef.Codec{
		Name: "文本数组加密", Desc: "文本数组加密，自定义字符表的base64编码，字典：" + codecConfig.Strings,
		Encode: func(value Any) (Any, error) {
			text := ""
			if vvs, ok := value.([]string); ok {
				text = strings.Join(vvs, "\n")
			} else {
				text = anyToString(value)
			}
			return textCoder.EncodeToString([]byte(text)), nil
		},
		Decode: func(data Any, value Any) (Any, error) {
			text := anyToString(value)
			bytes, err := textCoder.DecodeString(text)
			if err != nil {
				return nil, err
			}
			return strings.Split(string(bytes), "\n"), nil
		},
	}, false)

	chef.Register("number", chef.Codec{
		Name: "数字加密", Desc: "数字加密",
		Encode: func(value Any) (Any, error) {
			num := int64(0)
			if vv, ok := value.(int); ok {
				num = int64(vv)
			} else if vv, ok := value.(int64); ok {
				num = int64(vv)
			} else if vv, ok := value.(string); ok {
				if v, e := strconv.ParseInt(vv, 10, 64); e == nil {
					num = v
				} else {
					return "", errors.New("无效数字")
				}
			} else {
				return "", errors.New("无效数字")
			}
			return digitCoder.EncodeInt64([]int64{num})
		},
		Decode: func(data Any, value Any) (Any, error) {
			text := anyToString(value)
			digits, err := digitCoder.DecodeInt64WithError(text)
			if err != nil {
				return nil, err
			}
			if len(digits) == 0 {
				return nil, errors.New("无效结果")
			}
			return digits[0], nil
		},
	}, false)

	chef.Register("numbers", chef.Codec{
		Name: "数字数组加密", Desc: "数字数组加密",
		Encode: func(value Any) (Any, error) {
			nums := []int64{}
			if vv, ok := value.(int); ok {
				nums = append(nums, int64(vv))
			} else if vv, ok := value.(int64); ok {
				nums = append(nums, vv)
			} else if vvs, ok := value.([]int); ok {
				for _, num := range vvs {
					nums = append(nums, int64(num))
				}
			} else if vvs, ok := value.([]int64); ok {
				for _, num := range vvs {
					nums = append(nums, num)
				}
			} else if vv, ok := value.(string); ok {
				if v, e := strconv.ParseInt(vv, 10, 64); e == nil {
					nums = append(nums, int64(v))
				} else {
					return "", errors.New("无效数字")
				}
			} else {
				return "", errors.New("无效数字")
			}
			return digitCoder.EncodeInt64(nums)
		},
		Decode: func(data Any, value Any) (Any, error) {
			text := anyToString(value)
			return digitCoder.DecodeInt64WithError(text)
		},
	}, false)

}
