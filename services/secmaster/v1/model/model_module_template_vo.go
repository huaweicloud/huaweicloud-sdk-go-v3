package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModuleTemplateVo 模板具体内容
type ModuleTemplateVo struct {

	// 描述信息
	Description *string `json:"description,omitempty"`

	// 模板名称
	Name *string `json:"name,omitempty"`

	// UUID
	TemplateId *string `json:"template_id,omitempty"`

	// 名称
	Title *string `json:"title,omitempty"`
}

func (o ModuleTemplateVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModuleTemplateVo struct{}"
	}

	return strings.Join([]string{"ModuleTemplateVo", string(data)}, " ")
}
