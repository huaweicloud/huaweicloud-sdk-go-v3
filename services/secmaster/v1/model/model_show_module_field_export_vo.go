package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowModuleFieldExportVo struct {

	// 所属租户名称
	Name *string `json:"name,omitempty"`

	// UUID
	TemplateFieldId *string `json:"template_field_id,omitempty"`

	// 类型
	Type *string `json:"type,omitempty"`

	// 值
	Value *string `json:"value,omitempty"`
}

func (o ShowModuleFieldExportVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowModuleFieldExportVo struct{}"
	}

	return strings.Join([]string{"ShowModuleFieldExportVo", string(data)}, " ")
}
