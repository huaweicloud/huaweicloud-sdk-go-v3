package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateModuleVo struct {

	// 相关描述信息
	Children *[]CreateModuleVo `json:"children,omitempty"`

	// UUID
	ConnectionModuleId *string `json:"connection_module_id,omitempty"`

	// 相关描述信息
	Fields *[]ModuleFieldVo `json:"fields,omitempty"`

	// 名称
	Name *string `json:"name,omitempty"`

	// UUID
	TemplateId *string `json:"template_id,omitempty"`
}

func (o CreateModuleVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateModuleVo struct{}"
	}

	return strings.Join([]string{"CreateModuleVo", string(data)}, " ")
}
