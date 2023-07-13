package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowWhiteBlackIpRuleRequest Request Object
type ShowWhiteBlackIpRuleRequest struct {

	// 您可以通过调用企业项目管理服务（EPS）的查询企业项目列表接口（ListEnterpriseProject）查询企业项目id
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 防护策略id，您可以通过调用查询防护策略列表（ListPolicy）获取策略id
	PolicyId string `json:"policy_id"`

	// 黑白名单规则id，您可以通过调用查询黑白名单规则列表（ListWhiteblackipRule）获取黑白名单规则id
	RuleId string `json:"rule_id"`
}

func (o ShowWhiteBlackIpRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowWhiteBlackIpRuleRequest struct{}"
	}

	return strings.Join([]string{"ShowWhiteBlackIpRuleRequest", string(data)}, " ")
}
