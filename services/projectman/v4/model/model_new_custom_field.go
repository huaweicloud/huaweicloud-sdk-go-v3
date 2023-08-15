package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NewCustomField 自定义字段
type NewCustomField struct {

	// 自定义字段
	CustomField *string `json:"custom_field,omitempty"`

	// 自定义字段名称
	FieldName *string `json:"field_name,omitempty"`

	// 自定义属性对应的值，多个值以英文逗号区分开
	Value *string `json:"value,omitempty"`
}

func (o NewCustomField) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NewCustomField struct{}"
	}

	return strings.Join([]string{"NewCustomField", string(data)}, " ")
}
