package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PolicyComplianceSummaryUnit 规则总结详情
type PolicyComplianceSummaryUnit struct {

	// 合规数量
	CompliantCount *int32 `json:"compliant_count,omitempty"`

	// 不合规数量
	NonCompliantCount *int32 `json:"non_compliant_count,omitempty"`
}

func (o PolicyComplianceSummaryUnit) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PolicyComplianceSummaryUnit struct{}"
	}

	return strings.Join([]string{"PolicyComplianceSummaryUnit", string(data)}, " ")
}
