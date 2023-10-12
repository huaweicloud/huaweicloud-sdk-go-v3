package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateFactorySupplementDataInstanceRequestBody struct {

	// 补数据实例名称
	Name string `json:"name"`

	// 作业名称
	JobName string `json:"job_name"`

	// 补数据开始时间
	StartDate string `json:"start_date"`

	// 补数据结束时间
	EndDate string `json:"end_date"`

	// 并行周期数
	Parallel int32 `json:"parallel"`

	// 依赖作业信息
	DependJobs *[]CreateFactorySupplementDataInstanceRequestBodyDependJobs `json:"depend_jobs,omitempty"`

	// 是否按天粒度补数据
	IsDayGranularity *bool `json:"is_day_granularity,omitempty"`

	// 优先级
	Priority *int32 `json:"priority,omitempty"`

	// 失败时作业是否停止
	IsStopWhenFail *bool `json:"is_stop_when_fail,omitempty"`

	// 按照时间倒序补跑
	ReverseOrder *int32 `json:"reverse_order,omitempty"`

	SupplementDataRunTime *CreateFactorySupplementDataInstanceRequestBodySupplementDataRunTime `json:"supplement_data_run_time,omitempty"`

	SupplementDataInstanceTime *CreateFactorySupplementDataInstanceRequestBodySupplementDataInstanceTime `json:"supplement_data_instance_time,omitempty"`

	// 当前有补数据实例在运行时，是否强制补数据
	Force *string `json:"force,omitempty"`
}

func (o CreateFactorySupplementDataInstanceRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateFactorySupplementDataInstanceRequestBody struct{}"
	}

	return strings.Join([]string{"CreateFactorySupplementDataInstanceRequestBody", string(data)}, " ")
}
