package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateComplianceRuleResponse Response Object
type CreateComplianceRuleResponse struct {

	// 合规id
	ComplianceId   *string `json:"compliance_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateComplianceRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateComplianceRuleResponse struct{}"
	}

	return strings.Join([]string{"CreateComplianceRuleResponse", string(data)}, " ")
}
