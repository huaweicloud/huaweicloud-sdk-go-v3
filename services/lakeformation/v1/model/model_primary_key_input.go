package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PrimaryKeyInput 主键信息
type PrimaryKeyInput struct {

	// 列名称
	ColumnName string `json:"column_name"`

	// 主键名称
	PrimaryKeyName string `json:"primary_key_name"`

	// 是否启用主键
	EnableConstraint bool `json:"enable_constraint"`

	// 主键排序顺序
	KeySequence *int32 `json:"key_sequence,omitempty"`

	// 是否被外键依赖
	RelyConstraint bool `json:"rely_constraint"`

	// 限制条件是否可用
	ValidateConstraint bool `json:"validate_constraint"`
}

func (o PrimaryKeyInput) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PrimaryKeyInput struct{}"
	}

	return strings.Join([]string{"PrimaryKeyInput", string(data)}, " ")
}
