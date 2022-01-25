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
	// 是否该语言默认规则集，0不是，1是

	IsDefault *string `json:"is_default,omitempty"`
	// 是否是项目下语言默认规则集，0不是，1是

	IsDevcloudProjectDefault *string `json:"is_devcloud_project_default,omitempty"`
	// 是否是系统规则集，0不是，1是

	IsDefaultTemplate *string `json:"is_default_template,omitempty"`
}

func (o RulesetItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RulesetItem struct{}"
	}

	return strings.Join([]string{"RulesetItem", string(data)}, " ")
}
