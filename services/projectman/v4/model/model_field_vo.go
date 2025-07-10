package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FieldVo 字段参数返回体
type FieldVo struct {

	// 字段名称
	DisplayName *string `json:"display_name,omitempty"`

	// 添加人
	CreatedBy *string `json:"created_by,omitempty"`

	// 字段类型
	FieldType *string `json:"field_type,omitempty"`

	// 是否显示在新建页
	ShowOnCard *bool `json:"show_on_card,omitempty"`

	// 是否必填
	Optional *bool `json:"optional,omitempty"`

	// 字段选项
	AllOptions *[]OptionEntity `json:"all_options,omitempty"`

	// 默认值
	DefaultValue *string `json:"default_value,omitempty"`
}

func (o FieldVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FieldVo struct{}"
	}

	return strings.Join([]string{"FieldVo", string(data)}, " ")
}
