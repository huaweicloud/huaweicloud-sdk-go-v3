package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 删除单个标签资源结构
type DeleteResourceTag struct {

	// 键。同一资源的key值不能重复。最大长度为36个UNICODE字符。key不能为空，不允许为空字符串。不能包含以下字符：非打印字符ASCII(0-31)特殊字符“*”,“<”,“>”,“\\”,“=”,“,”,“|”,“/”键。不能为空。对于同一资源键值唯一。
	Key string `json:"key" xml:"key"`

	// 值。最大长度为43个UNICODE字符。删除时如果value有值按照key/value删除，如果value没值，则按照key删除。当value存在时，不能为空，可以为空字符串。不能包含以下字符：非打印字符ASCII(0-31)特殊字符“*”,“<”,“>”,“\\”,“=”,“,”,“|”,“/”。长度不超过43个字符。
	Value *string `json:"value,omitempty" xml:"value"`
}

func (o DeleteResourceTag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteResourceTag struct{}"
	}

	return strings.Join([]string{"DeleteResourceTag", string(data)}, " ")
}
