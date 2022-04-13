package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 单个标签资源结构
type ResourceTag struct {
	// 键。同一资源的key值不能重复。最大长度为36个UNICODE字符。key不能为空，不允许为空字符串。不能包含以下字符：非打印字符ASCII(0-31)特殊字符“*”,“<”,“>”,“\\”,“=”,“,”,“|”,“/”键。不能为空。对于同一资源键值唯一。

	Key string `json:"key"`
	// 值。最大长度为43个UNICODE字符。value不能为空，可以为空字符串。不能包含以下字符：非打印字符ASCII(0-31)特殊字符“*”,“<”,“>”,“\\”,“=”,“,”,“|”,“/”。长度不超过43个字符。

	Value string `json:"value"`
}

func (o ResourceTag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceTag struct{}"
	}

	return strings.Join([]string{"ResourceTag", string(data)}, " ")
}
