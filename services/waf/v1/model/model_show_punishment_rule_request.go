package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPunishmentRuleRequest Request Object
type ShowPunishmentRuleRequest struct {

	// 您可以通过调用企业项目管理服务（EPS）的查询企业项目列表接口（ListEnterpriseProject）查询企业项目id
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 防护策略id，您可以通过调用查询防护策略列表（ListPolicy）获取策略id
	PolicyId string `json:"policy_id"`

	// 攻击惩罚规则id，通过查询攻击惩罚规则列表接口（ListPunishmentRules）获取
	RuleId string `json:"rule_id"`
}

func (o ShowPunishmentRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPunishmentRuleRequest struct{}"
	}

	return strings.Join([]string{"ShowPunishmentRuleRequest", string(data)}, " ")
}
