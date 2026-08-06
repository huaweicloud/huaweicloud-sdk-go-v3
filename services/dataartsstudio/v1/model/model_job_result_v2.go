package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// JobResultV2 作业详情信息
type JobResultV2 struct {

	// 作业名称。
	Name *string `json:"name,omitempty"`

	// 作业类型： - BATCH: 批处理作业 - REAL_TIME: 实时作业
	JobType *string `json:"job_type,omitempty"`

	// 作业状态。
	Status *string `json:"status,omitempty"`

	// 作业创建者。
	CreateUser *string `json:"create_user,omitempty"`

	// 作业创建时间，13位时间戳。
	CreateTime *int64 `json:"create_time,omitempty"`

	// 作业开始时间，13位时间戳。
	StartTime *int64 `json:"start_time,omitempty"`

	// 作业结束时间，13位时间戳。
	EndTime *int64 `json:"end_time,omitempty"`

	// 上次实例运行状态。
	LastInstanceStatus *string `json:"last_instance_status,omitempty"`

	// 上次实例结束时间，13位时间戳。
	LastInstanceEndTime *int64 `json:"last_instance_end_time,omitempty"`

	// 作业负责人。
	Owner *string `json:"owner,omitempty"`

	// 最后更新人。
	LastUpdateUser *string `json:"last_update_user,omitempty"`

	// 作业优先级。
	Priority *int32 `json:"priority,omitempty"`

	// Flink作业信息。
	FlinkJobInfo *string `json:"flink_job_info,omitempty"`

	// 作业路径。
	Path *string `json:"path,omitempty"`

	// 是否为单节点作业。
	SingleNodeJobFlag *bool `json:"single_node_job_flag,omitempty"`

	// 告警信息列表。
	Alarms *[]JobAlarm `json:"alarms,omitempty"`

	// 最后更新时间，13位时间戳。
	LastUpdateTime *int64 `json:"last_update_time,omitempty"`

	// 单节点作业类型。
	SingleNodeJobType *string `json:"single_node_job_type,omitempty"`

	// 空跑作业标识。
	EmptyRunningJob *string `json:"empty_running_job,omitempty"`

	// 下次计划执行时间。
	NextPlanTime *string `json:"next_plan_time,omitempty"`
}

func (o JobResultV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobResultV2 struct{}"
	}

	return strings.Join([]string{"JobResultV2", string(data)}, " ")
}
