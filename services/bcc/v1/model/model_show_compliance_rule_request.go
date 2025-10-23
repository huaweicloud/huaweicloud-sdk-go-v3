package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowComplianceRuleRequest Request Object
type ShowComplianceRuleRequest struct {

	// 账号ID
	DomainId string `json:"domain_id"`

	// 合规Id
	ComplianceId string `json:"compliance_id"`
}

func (o ShowComplianceRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowComplianceRuleRequest struct{}"
	}

	return strings.Join([]string{"ShowComplianceRuleRequest", string(data)}, " ")
}
