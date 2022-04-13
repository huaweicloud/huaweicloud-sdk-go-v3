package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type RulesRemote struct {
	// 表示IdP断言中的属性。

	Type string `json:"type"`
	// 输入属性值中包含指定值才生效，并返回布尔值，返回值不能用于local块中的占位符。在同一个remote数组元素中，any_one_of与not_any_of互斥，两者至多填写一个，不能同时填写。

	AnyOneOf *[]string `json:"any_one_of,omitempty"`
	// 输入属性值中不包含指定值才生效，并返回布尔值，返回值不能用于local块中的占位符。在同一个remote数组元素中，any_one_of与not_any_of互斥，两者至多填写一个，不能同时填写。

	NotAnyOf *[]string `json:"not_any_of,omitempty"`
}

func (o RulesRemote) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RulesRemote struct{}"
	}

	return strings.Join([]string{"RulesRemote", string(data)}, " ")
}
