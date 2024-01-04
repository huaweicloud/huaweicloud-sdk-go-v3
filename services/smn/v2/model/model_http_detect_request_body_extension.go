package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HttpDetectRequestBodyExtension Http/Https协议可以指定多个K/V对，在发送Http/Https协议消息时会携带这些K/V对作为请求header。extension字段允许为空，header字段允许为空。
type HttpDetectRequestBodyExtension struct {

	// header应满足如下要求： 1. key值限定为：包含英文字母([A-Za-z])、数字([0-9])、中划线(-)hyphens，中划线不得作为结尾，不得连续出现。 2. K/V不得超过10个 3. key需要以\"x-\"开头，不能以\"x-smn\"开头，正确示例：x-abc-cba,  x-abc 4. 所有K/V长度总和不得超过1024个字符 5. key不区分大小写 6. key值不可重复 7. value值限定为ASCII码，不支持中文或其他Unicode，支持空格
	Header map[string]string `json:"header,omitempty"`
}

func (o HttpDetectRequestBodyExtension) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HttpDetectRequestBodyExtension struct{}"
	}

	return strings.Join([]string{"HttpDetectRequestBodyExtension", string(data)}, " ")
}
