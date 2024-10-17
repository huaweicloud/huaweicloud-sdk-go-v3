package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAuditRuleRiskRequest Request Object
type ShowAuditRuleRiskRequest struct {

	// 实例ID。可在查询实例列表接口的ID字段获取。
	InstanceId string `json:"instance_id"`

	// 风险规则ID。可在查询风险规则策略接口的ID字段获取。
	RiskId string `json:"risk_id"`
}

func (o ShowAuditRuleRiskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAuditRuleRiskRequest struct{}"
	}

	return strings.Join([]string{"ShowAuditRuleRiskRequest", string(data)}, " ")
}
