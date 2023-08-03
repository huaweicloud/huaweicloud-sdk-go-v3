package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ConformancePackComplianceSummary 合规结果概览。
type ConformancePackComplianceSummary struct {

	// 合规规则包ID。
	Id *string `json:"id,omitempty"`

	// 合规规则包名称。
	Name *string `json:"name,omitempty"`

	// 合规规则包合规结果。
	Compliance *string `json:"compliance,omitempty"`
}

func (o ConformancePackComplianceSummary) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConformancePackComplianceSummary struct{}"
	}

	return strings.Join([]string{"ConformancePackComplianceSummary", string(data)}, " ")
}
