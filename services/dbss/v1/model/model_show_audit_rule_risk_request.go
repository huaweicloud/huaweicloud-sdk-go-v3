package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAuditRuleRiskRequest Request Object
type ShowAuditRuleRiskRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 风险规则ID
	RiskId string `json:"risk_id"`
}

func (o ShowAuditRuleRiskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAuditRuleRiskRequest struct{}"
	}

	return strings.Join([]string{"ShowAuditRuleRiskRequest", string(data)}, " ")
}
