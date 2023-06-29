package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Ruleset 创建规则集信息
type Ruleset struct {

	// 项目ID
	ProjectId string `json:"project_id"`

	// 新规则集名称
	TemplateName string `json:"template_name"`

	// 规则集语言
	Language string `json:"language"`

	// 如果有基于的规则集则是1，没有基于的规则集则是0
	IsDefault string `json:"is_default"`

	// 新启用规则ids
	RuleIds string `json:"rule_ids"`

	// 新关闭规则id
	UncheckIds *string `json:"uncheck_ids,omitempty"`

	// 规则集ID
	TemplateId *string `json:"template_id,omitempty"`

	// 自定义规则参数项，支持修改规则阈值
	CustomAttributes *[]CustomAttributes `json:"custom_attributes,omitempty"`
}

func (o Ruleset) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Ruleset struct{}"
	}

	return strings.Join([]string{"Ruleset", string(data)}, " ")
}
