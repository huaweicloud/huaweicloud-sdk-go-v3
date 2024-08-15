package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NonCompliantSummary 不合规总结
type NonCompliantSummary struct {

	// 不合规补丁数量
	NonCompliantCount *int32 `json:"non_compliant_count,omitempty"`

	SeveritySummary *SeveritySummary `json:"severity_summary,omitempty"`
}

func (o NonCompliantSummary) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NonCompliantSummary struct{}"
	}

	return strings.Join([]string{"NonCompliantSummary", string(data)}, " ")
}
