package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// resource_tag字段说明
type ResourceTag struct {
	// 标签的键。  最大长度36个字符。   key不能为空，不能包含非打印字符ASCII(0-31)，“=”,“*”,“<”,“>”,“\\”,“,”,“|”,“/”。只能包含大写字母、小写字母、数字，特殊字符\"-\"和\"_\"。

	Key string `json:"key"`
	// 标签的值。  每个值最大长度43个字符，可以为空字符串。  不能包含非打印字符ASCII(0-31)，“=”,“*”,“<”,“>”,“\\”,“,”,“|”,“/”。只能包含大写字母、小写字母、数字，特殊字符\"-\"和\"_\"。

	Value string `json:"value"`
}

func (o ResourceTag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceTag struct{}"
	}

	return strings.Join([]string{"ResourceTag", string(data)}, " ")
}
