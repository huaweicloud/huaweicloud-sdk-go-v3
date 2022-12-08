package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 资源栈输出
type StackOutput struct {

	// 资源栈输出的name，由用户自己在模板中定义
	Name *string `json:"name,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`

	// 输出的类型
	Type *string `json:"type,omitempty"`

	// 输出的值(json字符串)
	Value *string `json:"value,omitempty"`

	// 是否为敏感信息
	Sensitive *bool `json:"sensitive,omitempty"`
}

func (o StackOutput) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StackOutput struct{}"
	}

	return strings.Join([]string{"StackOutput", string(data)}, " ")
}
