package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateComplianceRuleRequest Request Object
type UpdateComplianceRuleRequest struct {

	// 账号ID
	DomainId string `json:"domain_id"`

	// 合规id
	ComplianceId string `json:"compliance_id"`

	Body *UpdateComplianceRuleRequestBody `json:"body,omitempty"`
}

func (o UpdateComplianceRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateComplianceRuleRequest struct{}"
	}

	return strings.Join([]string{"UpdateComplianceRuleRequest", string(data)}, " ")
}
