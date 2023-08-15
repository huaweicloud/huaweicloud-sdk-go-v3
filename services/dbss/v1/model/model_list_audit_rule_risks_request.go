package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAuditRuleRisksRequest Request Object
type ListAuditRuleRisksRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 风险名称
	Name *string `json:"name,omitempty"`

	// 风险级别[LOW,MEDIUM,HIGH,NO_RISK]
	RiskLevels *string `json:"risk_levels,omitempty"`
}

func (o ListAuditRuleRisksRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAuditRuleRisksRequest struct{}"
	}

	return strings.Join([]string{"ListAuditRuleRisksRequest", string(data)}, " ")
}
