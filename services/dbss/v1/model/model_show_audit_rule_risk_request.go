package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAuditRuleRiskRequest Request Object
type ShowAuditRuleRiskRequest struct {

	// **参数解释**： 实例ID。可通过查询实例列表接口ID字段获取 **约束限制**： 不涉及 **取值范围**： 以查询实例列表接口值为准，字符长度32-64。 **默认取值**： 不涉及
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
