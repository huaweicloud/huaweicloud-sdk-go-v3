package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RulesetItem struct {

	// **参数解释:** 规则集ID。 **取值范围:** 32位UUID。
	TemplateId *string `json:"template_id,omitempty"`

	// **参数解释:** 规则集语言。 **取值范围:** 当前支持的语言有： - ARKTS - C# - CANGJIE - C++ - CSS - GO - HTML - JAVA - JAVASCRIPT - KOTLIN - LUA - PHP - PYTHON - RUST - SCALA - SHELL - SQL - TERRAFORM - TYPESCRIPT - AZURERESOURCEMANAGER - DOCKER - KUBERNETES
	Language *string `json:"language,omitempty"`

	// **参数解释:** 规则集语言。 **取值范围:** 任意中英文，数字，点，连字符-及下划线_。长度为[1~128]个字符。
	TemplateName *string `json:"template_name,omitempty"`

	// **参数解释:** 创建者的user_id，值来源于IAM的用户ID。 **取值范围:** 32位UUID。
	CreatorId *string `json:"creator_id,omitempty"`

	// **参数解释:** 创建人名称，值来源于IAM用户名称。 **取值范围:** IAM用户名称。
	CreatorName *string `json:"creator_name,omitempty"`

	// **参数解释:** 创建时间，时区统一为0时区。 **取值范围:** 0时区时间。
	TemplateCreateTime *string `json:"template_create_time,omitempty"`

	// **参数解释:** 使用状态。 **取值范围:** - 0：空闲中。 - 1：使用中。
	IsUsed *string `json:"is_used,omitempty"`

	// **参数解释:** 规则集包含的规则ID，每条规则都有唯一的规则ID。 **取值范围:** 32位UUID列表，多个UUID以英文逗号分隔。
	RuleIds *string `json:"rule_ids,omitempty"`

	// **参数解释:** 规则集是否是该语言默认规则集。 **取值范围:** - 0：不是。 - 1：是。
	IsDefault *string `json:"is_default,omitempty"`

	// **参数解释:** 规则集是否是项目下该语言默认规则集。 **取值范围:** - 0：不是。 - 1：是。
	IsDevcloudProjectDefault *string `json:"is_devcloud_project_default,omitempty"`

	// **参数解释:** 规则集是否是系统规则集。 **取值范围:** - 0：不是。 - 1：是。
	IsDefaultTemplate *string `json:"is_default_template,omitempty"`
}

func (o RulesetItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RulesetItem struct{}"
	}

	return strings.Join([]string{"RulesetItem", string(data)}, " ")
}
