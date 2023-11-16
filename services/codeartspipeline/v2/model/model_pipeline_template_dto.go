package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PipelineTemplateDto struct {

	// 模板名称
	Name string `json:"name"`

	// 模板描述
	Description *string `json:"description,omitempty"`

	// 模板语言
	Language string `json:"language"`

	Variables *CustomVariable `json:"variables,omitempty"`

	// 模板编排json，包含stages
	Definition string `json:"definition"`

	// 是否系统模板
	IsSystem bool `json:"is_system"`

	// 所属租户ID
	DomainId string `json:"domain_id"`

	// 是否显示流水线源
	IsShowSource bool `json:"is_show_source"`
}

func (o PipelineTemplateDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PipelineTemplateDto struct{}"
	}

	return strings.Join([]string{"PipelineTemplateDto", string(data)}, " ")
}
