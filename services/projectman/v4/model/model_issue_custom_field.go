package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 自定义字段
type IssueCustomField struct {

	// 自定义字段
	Name *string `json:"name,omitempty" xml:"name"`

	// 自定义字段
	CustomField *string `json:"custom_field,omitempty" xml:"custom_field"`

	// 自定义字段的可选值，多个值以英文逗号区分
	Options *string `json:"options,omitempty" xml:"options"`

	// 自定义字段类型， textArea 多行文本，只能包含汉字、英文大小写字母、数字、下划线和连接符，不能超过500字符； text 单行文本， 只能包含汉字、英文大小写字母、数字、下划线和连接符，不能超过500字符； select 下拉框，只能包含汉字、英文大小写字母、数字、下划线和连接符，每个选项最大长度40个字符，最多可定义60个选项； number 数字，取值范围由用户创建自定义字段时设置； date 日期 精确到年月日， time_date 日期 精确到时分秒， 长整型时间戳； checkbox 多选框，只能包含汉字、英文大小写字母、数字、下划线和连接符，每个选项最大长度40个字符，最多可定义60个选项； radio 单选框，只能包含汉字、英文大小写字母、数字、下划线和连接符，每个选项最大长度40个字符，最多可定义60个选项；
	Type *string `json:"type,omitempty" xml:"type"`

	// 自定义字段支持的工作项类型 2任务/Task,3缺陷/Bug,5Epic,6Feature,7Story
	TrackerIds *[]int32 `json:"tracker_ids,omitempty" xml:"tracker_ids"`
}

func (o IssueCustomField) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IssueCustomField struct{}"
	}

	return strings.Join([]string{"IssueCustomField", string(data)}, " ")
}
