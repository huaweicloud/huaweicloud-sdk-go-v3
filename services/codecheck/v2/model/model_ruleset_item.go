package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RulesetItem struct {
	// 规则集id

	TemplateId *string `json:"template_id,omitempty"`
	// 规则集语言

	Language *string `json:"language,omitempty"`
	// 规则集名称

	TemplateName *string `json:"template_name,omitempty"`
	// 创建人ID

	CreatorId *string `json:"creator_id,omitempty"`
	// 创建人名称

	CreatorName *string `json:"creator_name,omitempty"`
	// 创建人时间

	TemplateCreateTime *string `json:"template_create_time,omitempty"`
	// 使用状态1使用中，0空闲中

	IsUsed *string `json:"is_used,omitempty"`
	// 规则集包含的规则id

	RuleIds *string `json:"rule_ids,omitempty"`
	// 是否默认规则集，0不是，1是

	IsDefault *string `json:"is_default,omitempty"`
}

func (o RulesetItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RulesetItem struct{}"
	}

	return strings.Join([]string{"RulesetItem", string(data)}, " ")
}
