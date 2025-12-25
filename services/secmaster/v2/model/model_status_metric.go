package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StatusMetric 状态指标
type StatusMetric struct {

	// 已开启的数量
	Enabled *int64 `json:"enabled,omitempty"`

	// 未开启的数量
	Disabled *int64 `json:"disabled,omitempty"`

	// 错误的数量
	Error *int64 `json:"error,omitempty"`

	// 是否百分比
	IsPercentage *bool `json:"is_percentage,omitempty"`
}

func (o StatusMetric) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StatusMetric struct{}"
	}

	return strings.Join([]string{"StatusMetric", string(data)}, " ")
}
