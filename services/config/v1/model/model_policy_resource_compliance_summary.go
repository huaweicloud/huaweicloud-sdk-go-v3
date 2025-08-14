package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PolicyResourceComplianceSummary 资源合规总结
type PolicyResourceComplianceSummary struct {

	// 规则合规状态
	ComplianceState *string `json:"compliance_state,omitempty"`

	Resource *PolicyResource `json:"resource,omitempty"`

	Results *PolicyComplianceSummaryResults `json:"results,omitempty"`
}

func (o PolicyResourceComplianceSummary) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PolicyResourceComplianceSummary struct{}"
	}

	return strings.Join([]string{"PolicyResourceComplianceSummary", string(data)}, " ")
}
