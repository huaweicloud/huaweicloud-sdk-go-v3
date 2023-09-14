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

	// 补数据开始时间
	StartDate *string `json:"startDate,omitempty"`

	// 补数据结束时间
	EndDate *string `json:"endDate,omitempty"`

	// 并行周期数
	Parallel *int32 `json:"parallel,omitempty"`

	// 依赖的作业信息
	DependJobs *[]JobInfo `json:"dependJobs,omitempty"`

	// 是否按天粒度补数据
	IsDayGranularity *bool `json:"isDayGranularity,omitempty"`

	// 优先级
	Priority *int32 `json:"priority,omitempty"`

	// 失败时作业是否停止
	IsStopWhenFail *bool `json:"is_stop_when_fail,omitempty"`

	// 按照时间倒序补跑
	ReverseOrder *int32 `json:"reverseOrder,omitempty"`

	// 当前有补数据实例在运行时，是否强制补数据
	Force *string `json:"force,omitempty"`

	SupplementDataRunTime *SupplementDataInfoSupplementDataRunTime `json:"supplement_data_run_time,omitempty"`

	SupplementDataInstanceTime *SupplementDataInfoSupplementDataInstanceTime `json:"supplement_data_instance_time,omitempty"`
}

func (o SupplementDataInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SupplementDataInfo struct{}"
	}

	return strings.Join([]string{"SupplementDataInfo", string(data)}, " ")
}
