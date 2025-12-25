package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowModuleExportVo struct {

	// 相关描述信息
	Children *[]ShowModuleExportVo `json:"children,omitempty"`

	// 相关描述信息
	Fields *[]ShowModuleFieldExportVo `json:"fields,omitempty"`

	// UUID
	ModuleId *string `json:"module_id,omitempty"`

	// 所属租户名称
	Name *string `json:"name,omitempty"`

	// UUID
	TemplateId *string `json:"template_id,omitempty"`
}

func (o ShowModuleExportVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowModuleExportVo struct{}"
	}

	return strings.Join([]string{"ShowModuleExportVo", string(data)}, " ")
}
