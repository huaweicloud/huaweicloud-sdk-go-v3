package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IssuableTimeStatsDto 合并请求的时间统计信息
type IssuableTimeStatsDto struct {

	// 合并请求时间估计
	TimeEstimate *int32 `json:"time_estimate,omitempty"`

	// 合并请求总时长
	TotalTimeSpent *int32 `json:"total_time_spent,omitempty"`

	// 合并请求人类时间估计
	HumanTimeEstimate *string `json:"human_time_estimate,omitempty"`

	// 合并请求人类总时长
	HumanTotalTimeSpent *string `json:"human_total_time_spent,omitempty"`
}

func (o IssuableTimeStatsDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IssuableTimeStatsDto struct{}"
	}

	return strings.Join([]string{"IssuableTimeStatsDto", string(data)}, " ")
}
