package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ObjectArrayPatterns 对象数组
type ObjectArrayPatterns struct {

	// 键
	Key *string `json:"key,omitempty"`

	// 对象类型
	Type *string `json:"type,omitempty"`

	// 最大长度
	MaxLength *int32 `json:"max_length,omitempty"`

	// 最大值
	MaxValue *float64 `json:"max_value,omitempty"`

	// 最小值
	MinValue *float64 `json:"min_value,omitempty"`

	// 是否可以为空值
	Nullable *bool `json:"nullable,omitempty"`

	// 提示信息
	Hint *string `json:"hint,omitempty"`
}

func (o ObjectArrayPatterns) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ObjectArrayPatterns struct{}"
	}

	return strings.Join([]string{"ObjectArrayPatterns", string(data)}, " ")
}
