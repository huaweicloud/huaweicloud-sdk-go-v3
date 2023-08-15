package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAntileakageRuleRequest Request Object
type DeleteAntileakageRuleRequest struct {

	// 您可以通过调用企业项目管理服务（EPS）的查询企业项目列表接口（ListEnterpriseProject）查询企业项目id
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 防护策略id，您可以通过调用查询防护策略列表（ListPolicy）获取策略id
	PolicyId string `json:"policy_id"`

	// 防敏感信息泄露规则id，通过查询防敏感信息泄露规则列表接口（ListAntileakageRules）获取
	RuleId string `json:"rule_id"`
}

func (o DeleteAntileakageRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAntileakageRuleRequest struct{}"
	}

	return strings.Join([]string{"DeleteAntileakageRuleRequest", string(data)}, " ")
}
