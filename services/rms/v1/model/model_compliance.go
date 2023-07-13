package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Compliance 合规规则合规性
type Compliance struct {

	// 合规结果。
	ComplianceState *string `json:"compliance_state,omitempty"`

	ResourceDetails *PolicyComplianceSummaryUnit `json:"resource_details,omitempty"`
}

func (o Compliance) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Compliance struct{}"
	}

	return strings.Join([]string{"Compliance", string(data)}, " ")
}
