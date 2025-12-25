package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateConnectionDto struct {

	// 描述信息
	Description *string `json:"description,omitempty"`

	// 相关描述信息
	Fields []ConnectionModuleFieldCreateVo `json:"fields"`

	// 所属租户名称
	Name string `json:"name"`

	// UUID
	TemplateId string `json:"template_id"`

	// 相关标题
	Title string `json:"title"`
}

func (o CreateConnectionDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateConnectionDto struct{}"
	}

	return strings.Join([]string{"CreateConnectionDto", string(data)}, " ")
}
