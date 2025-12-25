package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Tag 资源标签。
type Tag struct {

	// 标签键，取值可以包含任意语种字母、数字、空格，以及_.:=+-@特殊字符，但首尾不能含有空格，不能以_sys_开头。
	Key string `json:"key"`

	// 标签值，格式为取值可以包含任意语种字母、数字、空格，以及_.:/=+-@特殊字符。
	Value *string `json:"value,omitempty"`
}

func (o Tag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Tag struct{}"
	}

	return strings.Join([]string{"Tag", string(data)}, " ")
}
