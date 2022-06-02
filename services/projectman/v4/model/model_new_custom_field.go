package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 自定义字段
type NewCustomField struct {

	// 自定义字段
	CustomField *string `json:"custom_field,omitempty"`

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
