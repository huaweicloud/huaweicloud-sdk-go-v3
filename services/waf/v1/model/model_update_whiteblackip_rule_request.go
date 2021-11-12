package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateWhiteblackipRuleRequest struct {
	// 企业项目id

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
	// 策略id（策略id从查询防护策略列表接口获取）

	PolicyId string `json:"policy_id"`
	// 黑白名单规则ID（从查询黑白名单规则列表ListWhiteblackipRule接口获取）

	RuleId string `json:"rule_id"`

	Body *UpdateWhiteBlackIpRuleRequestBody `json:"body,omitempty"`
}

func (o UpdateWhiteblackipRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateWhiteblackipRuleRequest struct{}"
	}

	return strings.Join([]string{"UpdateWhiteblackipRuleRequest", string(data)}, " ")
}
