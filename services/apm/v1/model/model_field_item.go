package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 所需展示的字段列表模型
type FieldItem struct {

	// 表达式
	Function *string `json:"function,omitempty" xml:"function"`

	// 作为
	As *string `json:"as,omitempty" xml:"as"`

	// 默认值
	DefaultValue *string `json:"default_value,omitempty" xml:"default_value"`

	// 是否是trace
	Trace *bool `json:"trace,omitempty" xml:"trace"`

	// 百分比
	Precision *int32 `json:"precision,omitempty" xml:"precision"`

	// 单位
	Unit *string `json:"unit,omitempty" xml:"unit"`

	// 是否可见
	Visible *bool `json:"visible,omitempty" xml:"visible"`
}

func (o FieldItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FieldItem struct{}"
	}

	return strings.Join([]string{"FieldItem", string(data)}, " ")
}
