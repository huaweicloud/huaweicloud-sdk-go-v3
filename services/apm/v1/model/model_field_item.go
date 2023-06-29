package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FieldItem 所需展示的字段列表模型。
type FieldItem struct {

	// 表达式。
	Function string `json:"function"`

	// 作为。
	As *string `json:"as,omitempty"`

	// 默认值。
	DefaultValue *string `json:"default_value,omitempty"`

	// 是否是trace。
	Trace *bool `json:"trace,omitempty"`

	// 百分比。
	Precision *int32 `json:"precision,omitempty"`

	// 单位。
	Unit *string `json:"unit,omitempty"`

	// 是否可见。
	Visible *bool `json:"visible,omitempty"`
}

func (o FieldItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FieldItem struct{}"
	}

	return strings.Join([]string{"FieldItem", string(data)}, " ")
}
