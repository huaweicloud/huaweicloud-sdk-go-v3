package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RulesetItem struct {

	// 规则集id
	TemplateId *string `json:"template_id,omitempty" xml:"template_id"`

	// 规则集语言
	Language *string `json:"language,omitempty" xml:"language"`

	// 规则集名称
	TemplateName *string `json:"template_name,omitempty" xml:"template_name"`

	// 创建人ID
	CreatorId *string `json:"creator_id,omitempty" xml:"creator_id"`

	// 创建人名称
	CreatorName *string `json:"creator_name,omitempty" xml:"creator_name"`

	// 创建人时间
	TemplateCreateTime *string `json:"template_create_time,omitempty" xml:"template_create_time"`

	// 使用状态1使用中，0空闲中
	IsUsed *string `json:"is_used,omitempty" xml:"is_used"`

	// 规则集包含的规则id
	RuleIds *string `json:"rule_ids,omitempty" xml:"rule_ids"`

	// 是否该语言默认规则集，0不是，1是
	IsDefault *string `json:"is_default,omitempty" xml:"is_default"`

	// 是否是项目下语言默认规则集，0不是，1是
	IsDevcloudProjectDefault *string `json:"is_devcloud_project_default,omitempty" xml:"is_devcloud_project_default"`

	// 是否是系统规则集，0不是，1是
	IsDefaultTemplate *string `json:"is_default_template,omitempty" xml:"is_default_template"`
}

func (o RulesetItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RulesetItem struct{}"
	}

	return strings.Join([]string{"RulesetItem", string(data)}, " ")
}
