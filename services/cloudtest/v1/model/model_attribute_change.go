package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 对象，包括编号与名称
type AttributeChange struct {

	// 变更后的取值
	NewValue *string `json:"new_value,omitempty"`

	// 变更前的取值
	OldValue *string `json:"old_value,omitempty"`

	// 发生变更的测试计划属性
	AttributeType *string `json:"attribute_type,omitempty"`
}

func (o AttributeChange) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttributeChange struct{}"
	}

	return strings.Join([]string{"AttributeChange", string(data)}, " ")
}
