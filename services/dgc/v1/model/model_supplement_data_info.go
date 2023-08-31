package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SupplementDataInfo struct {

	// 补数据实例名称
	Name *string `json:"name,omitempty"`

	// 补数据依赖的作业名称
	JobName *string `json:"jobName,omitempty"`

	// 开始补数据时间
	StartDate *string `json:"startDate,omitempty"`

	// 截止补数据时间
	EndDate *string `json:"endDate,omitempty"`

	// 优先级
	Parallel *int32 `json:"parallel,omitempty"`

	// 依赖的作业信息
	DependJobs *[]JobInfo `json:"dependJobs,omitempty"`

	IsDayGranularity *bool `json:"isDayGranularity,omitempty"`

	Priority *int32 `json:"priority,omitempty"`

	IsStopWhenFail *bool `json:"is_stop_when_fail,omitempty"`

	ReverseOrder *int32 `json:"reverseOrder,omitempty"`

	Force *string `json:"force,omitempty"`
}

func (o SupplementDataInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SupplementDataInfo struct{}"
	}

	return strings.Join([]string{"SupplementDataInfo", string(data)}, " ")
}
