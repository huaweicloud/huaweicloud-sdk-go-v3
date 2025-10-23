package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteComplianceRuleRequest Request Object
type DeleteComplianceRuleRequest struct {

	// 账号ID
	DomainId string `json:"domain_id"`

	// 合规id
	ComplianceId string `json:"compliance_id"`
}

func (o DeleteComplianceRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteComplianceRuleRequest struct{}"
	}

	return strings.Join([]string{"DeleteComplianceRuleRequest", string(data)}, " ")
}
