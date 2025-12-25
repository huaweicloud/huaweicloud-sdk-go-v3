package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateConnectionDto struct {

	// 描述信息
	Description *string `json:"description,omitempty"`

	// 相关描述信息
	Fields *[]ConnectionModuleFieldUpdateVo `json:"fields,omitempty"`

	// UUID
	ModuleId string `json:"module_id"`

	// 所属租户名称
	Name string `json:"name"`

	// UUID
	TemplateId string `json:"template_id"`

	// 相关标题
	Title string `json:"title"`
}

func (o UpdateConnectionDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateConnectionDto struct{}"
	}

	return strings.Join([]string{"UpdateConnectionDto", string(data)}, " ")
}
