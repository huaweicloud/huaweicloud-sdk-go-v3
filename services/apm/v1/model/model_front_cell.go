package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 数据单元
type FrontCell struct {

	// 数据类型
	DataType *string `json:"data_type,omitempty" xml:"data_type"`

	// 函数
	Function *string `json:"function,omitempty" xml:"function"`

	// 是否调用链
	Trace *bool `json:"trace,omitempty" xml:"trace"`

	// 是否是span信息，如果是就跳到调用链搜索页面
	Span *bool `json:"span,omitempty" xml:"span"`

	// span字段
	SpanField *string `json:"span_field,omitempty" xml:"span_field"`

	// 小数点位数
	Precision *int32 `json:"precision,omitempty" xml:"precision"`

	// 文本信息
	Text *string `json:"text,omitempty" xml:"text"`

	// 单位
	Unit *string `json:"unit,omitempty" xml:"unit"`

	// 是否可见
	Visible *bool `json:"visible,omitempty" xml:"visible"`
}

func (o FrontCell) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FrontCell struct{}"
	}

	return strings.Join([]string{"FrontCell", string(data)}, " ")
}
