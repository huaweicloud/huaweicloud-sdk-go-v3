package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TaskInfo 获取作业详情
type TaskInfo struct {

	// 作业名称
	Name *string `json:"name,omitempty"`

	// 作业创建时间
	Timestamp *string `json:"timestamp,omitempty"`

	// 作业描述
	Description *string `json:"description,omitempty"`

	// 作业流详情
	Streams *[]TaskStream `json:"streams,omitempty"`

	// 在实例上运行成功的作业数
	OkPodNumber *int32 `json:"ok_pod_number,omitempty"`

	// 在实例上正在运行的作业数
	CurPodNumber *int32 `json:"cur_pod_number,omitempty"`

	// 在实例上运行过的作业总数
	SumPodNumber *int32 `json:"sum_pod_number,omitempty"`

	// 在实例上运行失败的作业数
	FailPodNumber *int32 `json:"fail_pod_number,omitempty"`

	// 在实例上等待运行的作业数
	PendingPodNumber *int32 `json:"pending_pod_number,omitempty"`

	// 作业状态信息
	TaskStatus *[]TaskStatus `json:"task_status,omitempty"`

	// 作业id
	TaskId *string `json:"task_id,omitempty"`

	// 用户作业id
	UserTaskId *string `json:"user_task_id,omitempty"`

	StartTimeInfo *StartTimeInfo `json:"start_time_info,omitempty"`

	SourceUsageEstimate *TaskSourceUsageEstimate `json:"source_usage_estimate,omitempty"`
}

func (o TaskInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskInfo struct{}"
	}

	return strings.Join([]string{"TaskInfo", string(data)}, " ")
}
