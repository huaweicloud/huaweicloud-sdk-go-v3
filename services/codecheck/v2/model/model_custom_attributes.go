package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CustomAttributes struct {
	// 配置项属性，severity：为问题级别

	Attribute *string `json:"attribute,omitempty"`
	// 规则详细

	Rules *[]CustomAttributesRule `json:"rules,omitempty"`
}

func (o CustomAttributes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CustomAttributes struct{}"
	}

	return strings.Join([]string{"CustomAttributes", string(data)}, " ")
}
