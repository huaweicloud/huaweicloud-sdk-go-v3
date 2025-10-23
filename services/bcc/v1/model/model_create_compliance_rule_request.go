package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateComplianceRuleRequest Request Object
type CreateComplianceRuleRequest struct {

	// 账号ID
	DomainId string `json:"domain_id"`

	Body *CreateComplianceRuleRequestBody `json:"body,omitempty"`
}

func (o CreateComplianceRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateComplianceRuleRequest struct{}"
	}

	return strings.Join([]string{"CreateComplianceRuleRequest", string(data)}, " ")
}
