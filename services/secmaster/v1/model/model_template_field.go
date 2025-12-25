package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TemplateField struct {

	// 默认值
	DefaultValue *string `json:"default_value,omitempty"`

	// 描述信息
	Description *string `json:"description,omitempty"`

	// 示例
	Example *string `json:"example,omitempty"`

	// 规则名称
	Name *string `json:"name,omitempty"`

	// 是否必填
	Required *bool `json:"required,omitempty"`

	// 相关描述信息
	Restrictions *[]Restriction `json:"restrictions,omitempty"`

	// UUID
	TemplateFieldId *string `json:"template_field_id,omitempty"`

	// 相关标题
	Title *string `json:"title,omitempty"`

	// 规则类型
	Type *string `json:"type,omitempty"`
}

func (o TemplateField) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TemplateField struct{}"
	}

	return strings.Join([]string{"TemplateField", string(data)}, " ")
}
