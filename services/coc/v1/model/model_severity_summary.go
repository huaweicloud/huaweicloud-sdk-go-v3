package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SeveritySummary struct {

	// 重大合规性报告数量
	CriticalCount *int32 `json:"critical_count,omitempty"`

	// 高合规性报告数量
	HighCount *int32 `json:"high_count,omitempty"`

	// 信息性合规性报告数量
	InformationalCount *int32 `json:"informational_count,omitempty"`

	// 低合规性报告数量
	LowCount *int32 `json:"low_count,omitempty"`

	// 中级合规性报告数量
	MediumCount *int32 `json:"medium_count,omitempty"`

	// 未指定合规性报告数量
	UnspecifiedCount *int32 `json:"unspecified_count,omitempty"`
}

func (o SeveritySummary) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SeveritySummary struct{}"
	}

	return strings.Join([]string{"SeveritySummary", string(data)}, " ")
}
