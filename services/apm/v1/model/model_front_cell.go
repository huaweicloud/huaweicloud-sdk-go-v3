package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FrontCell 数据单元。
type FrontCell struct {

	// 数据类型。
	DataType *string `json:"data_type,omitempty"`

	// 函数。
	Function *string `json:"function,omitempty"`

	// 是否调用链。
	Trace *bool `json:"trace,omitempty"`

	// 是否是span信息，如果是就跳到调用链搜索页面。
	Span *bool `json:"span,omitempty"`

	// span字段。
	SpanField *string `json:"span_field,omitempty"`

	// 小数点位数。
	Precision *int32 `json:"precision,omitempty"`

	// 文本信息。
	Text *string `json:"text,omitempty"`

	// 单位。
	Unit *string `json:"unit,omitempty"`

	// 是否可见。
	Visible *bool `json:"visible,omitempty"`
}

func (o FrontCell) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FrontCell struct{}"
	}

	return strings.Join([]string{"FrontCell", string(data)}, " ")
}
