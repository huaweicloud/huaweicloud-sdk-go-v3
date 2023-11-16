package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RuleInstanceProperty struct {

	// 属性键
	Key string `json:"key"`

	// 类型
	Type string `json:"type"`

	// 展示名称
	Name string `json:"name"`

	// 比较运算符
	Operator *string `json:"operator,omitempty"`

	// 属性值
	Value string `json:"value"`

	// 数据类型
	ValueType string `json:"value_type"`
}

func (o RuleInstanceProperty) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RuleInstanceProperty struct{}"
	}

	return strings.Join([]string{"RuleInstanceProperty", string(data)}, " ")
}
