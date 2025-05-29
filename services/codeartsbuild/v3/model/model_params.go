package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Params 构建参数
type Params struct {

	// 参数名
	Name *string `json:"name,omitempty"`

	// 名称
	Title *string `json:"title,omitempty"`

	// 类型
	Type *string `json:"type,omitempty"`

	// 必填
	Required *bool `json:"required,omitempty"`

	// 简要构建信息列表
	Constraints *[]Constraints `json:"constraints,omitempty"`

	// 删除
	Deletion *bool `json:"deletion,omitempty"`

	// 默认
	Defaults *bool `json:"defaults,omitempty"`
}

func (o Params) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Params struct{}"
	}

	return strings.Join([]string{"Params", string(data)}, " ")
}
