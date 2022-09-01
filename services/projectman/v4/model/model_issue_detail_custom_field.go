package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type IssueDetailCustomField struct {

	// 自定义字段
	CustomField *string `json:"custom_field,omitempty" xml:"custom_field"`

	// 自定义字段名称
	FieldName *string `json:"field_name,omitempty" xml:"field_name"`

	// 自定义属性对应的值，多个值以英文逗号区分开
	Value *string `json:"value,omitempty" xml:"value"`

	// 自定义字段类型， textArea 多行文本，text 单行文本，select 下拉框，number 数字，time_date 日期，checkbox 多选框，radio 单选框
	FieldType *string `json:"field_type,omitempty" xml:"field_type"`

	// 自定义字段描述
	Description *string `json:"description,omitempty" xml:"description"`
}

func (o IssueDetailCustomField) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IssueDetailCustomField struct{}"
	}

	return strings.Join([]string{"IssueDetailCustomField", string(data)}, " ")
}
