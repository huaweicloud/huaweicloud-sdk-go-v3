package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CompliantSummary struct {

	// 合规补丁数量
	CompliantCount *int32 `json:"compliant_count,omitempty"`

	SeveritySummary *SeveritySummary `json:"severity_summary,omitempty"`
}

func (o CompliantSummary) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CompliantSummary struct{}"
	}

	return strings.Join([]string{"CompliantSummary", string(data)}, " ")
}
