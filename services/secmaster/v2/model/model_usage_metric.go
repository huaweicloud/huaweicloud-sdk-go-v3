package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UsageMetric 使用量指标
type UsageMetric struct {

	// 已使用
	Used *string `json:"used,omitempty"`

	// 未使用
	UnUsed *string `json:"un_used,omitempty"`

	// 是否百分比
	IsPercentage *bool `json:"is_percentage,omitempty"`
}

func (o UsageMetric) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UsageMetric struct{}"
	}

	return strings.Join([]string{"UsageMetric", string(data)}, " ")
}
