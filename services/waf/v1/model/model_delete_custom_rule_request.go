package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteCustomRuleRequest struct {

	// 您可以通过调用企业项目管理服务（EPS）的查询企业项目列表接口（ListEnterpriseProject）查询企业项目id
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 防护策略id，通过指定防护策略id来指明查询该防护策略下的防护规则，您可以通过调用查询防护策略列表（ListPolicy）获取策略id
	PolicyId string `json:"policy_id"`

	// 精准防护规则id，通过查询精准防护规则列表接口（ListCustomRules）获取
	RuleId string `json:"rule_id"`
}

func (o DeleteCustomRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteCustomRuleRequest struct{}"
	}

	return strings.Join([]string{"DeleteCustomRuleRequest", string(data)}, " ")
}
