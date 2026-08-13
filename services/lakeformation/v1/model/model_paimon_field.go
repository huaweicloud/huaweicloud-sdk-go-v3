package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PaimonField Paimon Schema中的字段定义，包含字段名、类型、是否可空及元数据。
type PaimonField struct {

	// 字段名称
	Name string `json:"name"`

	Type *PaimonType `json:"type"`

	// 字段是否允许为null
	Nullable *bool `json:"nullable,omitempty"`

	// 子字段
	Children *[]PaimonField `json:"children,omitempty"`

	// 字段描述
	Description *string `json:"description,omitempty"`
}

func (o PaimonField) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PaimonField struct{}"
	}

	return strings.Join([]string{"PaimonField", string(data)}, " ")
}
