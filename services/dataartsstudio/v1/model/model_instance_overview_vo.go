package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type InstanceOverviewVo struct {

	// id
	Id *int64 `json:"id,omitempty"`

	// 名称
	Name *string `json:"name,omitempty"`

	// task id
	TaskId *int64 `json:"task_id,omitempty"`

	// QUALITY_TASK:质量作业,CONSISTENCY_TASK:对账作业
	TaskType *string `json:"task_type,omitempty"`

	// RUNNING:运行中,FAILED:失败,ALARMING:报警,SUCCESS:正常
	RunStatus *string `json:"run_status,omitempty"`

	// NOT_TRIGGERED:未触发,SUCCESS:成功,FAILED:失败
	NotifyStatus *string `json:"notify_status,omitempty"`

	// 开始时间
	StartTime *int64 `json:"start_time,omitempty"`

	// 结束时间
	EndTime *int64 `json:"end_time,omitempty"`
}

func (o InstanceOverviewVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceOverviewVo struct{}"
	}

	return strings.Join([]string{"InstanceOverviewVo", string(data)}, " ")
}
