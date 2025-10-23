package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateComplianceRuleResponse Response Object
type UpdateComplianceRuleResponse struct {

	// 合规id
	ComplianceId   *string `json:"compliance_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateComplianceRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateComplianceRuleResponse struct{}"
	}

	return strings.Join([]string{"UpdateComplianceRuleResponse", string(data)}, " ")
}
