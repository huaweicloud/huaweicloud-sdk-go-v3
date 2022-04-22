package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteAntitamperRuleRequest struct {

	// 您可以通过调用企业项目管理服务（EPS)的查询企业项目列表接口（ListEnterpriseProject）查询企业项目id
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 防护策略id，您可以通过调用查询防护策略列表（ListPolicy）获取策略id
	PolicyId string `json:"policy_id"`

	// 防篡改规则id，通过查询防篡改规则列表接口（ListAntitamperRule）获取
	RuleId string `json:"rule_id"`
}

func (o DeleteAntitamperRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAntitamperRuleRequest struct{}"
	}

	return strings.Join([]string{"DeleteAntitamperRuleRequest", string(data)}, " ")
}
