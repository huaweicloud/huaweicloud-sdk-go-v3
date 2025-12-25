package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ModuleFieldVo struct {

	// UUID
	ConnectionModuleId *string `json:"connection_module_id,omitempty"`

	// 名称
	Name *string `json:"name,omitempty"`

	// 其他信息
	Other *string `json:"other,omitempty"`

	// UUID
	TemplateFieldId *string `json:"template_field_id,omitempty"`

	// 类型
	Type *string `json:"type,omitempty"`

	// 值
	Value *string `json:"value,omitempty"`
}

func (o ModuleFieldVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModuleFieldVo struct{}"
	}

	return strings.Join([]string{"ModuleFieldVo", string(data)}, " ")
}
