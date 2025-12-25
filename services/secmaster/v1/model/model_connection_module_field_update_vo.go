package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ConnectionModuleFieldUpdateVo struct {

	// UUID
	FieldId *string `json:"field_id,omitempty"`

	// 所属租户名称
	Name *string `json:"name,omitempty"`

	// UUID
	TemplateFieldId *string `json:"template_field_id,omitempty"`

	// 相关标题
	Title *string `json:"title,omitempty"`

	// value的数据类型
	Type *string `json:"type,omitempty"`

	// 相关值
	Value *string `json:"value,omitempty"`
}

func (o ConnectionModuleFieldUpdateVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConnectionModuleFieldUpdateVo struct{}"
	}

	return strings.Join([]string{"ConnectionModuleFieldUpdateVo", string(data)}, " ")
}
