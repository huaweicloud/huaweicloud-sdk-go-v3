package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Restriction struct {

	// 逻辑条件
	Logic *bool `json:"logic,omitempty"`

	// 相关标题
	Title *string `json:"title,omitempty"`

	// 规则类型
	Type *string `json:"type,omitempty"`

	// 规则名称
	Value *string `json:"value,omitempty"`
}

func (o Restriction) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Restriction struct{}"
	}

	return strings.Join([]string{"Restriction", string(data)}, " ")
}
