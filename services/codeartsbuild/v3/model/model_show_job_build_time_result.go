package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowJobBuildTimeResult 返回结果
type ShowJobBuildTimeResult struct {

	// 任务ID
	JobId *string `json:"job_id,omitempty"`

	// 平均构建时长
	AvgBuildTime *float64 `json:"avg_build_time,omitempty"`

	// 最长时间
	MaxBuildTime *int64 `json:"max_build_time,omitempty"`

	// 最短时间
	MinBuildTime *int64 `json:"min_build_time,omitempty"`

	// 每日构建数据
	Chart *[]ShowJobBuildTimeResultChart `json:"chart,omitempty"`
}

func (o ShowJobBuildTimeResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowJobBuildTimeResult struct{}"
	}

	return strings.Join([]string{"ShowJobBuildTimeResult", string(data)}, " ")
}
