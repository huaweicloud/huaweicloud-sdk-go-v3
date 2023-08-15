package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TaskDetailResult struct {
	InstanceInfo *InstanceInfoResult `json:"instance_info,omitempty"`

	// 任务ID。
	JobId *string `json:"job_id,omitempty"`

	// 任务名称。
	Name *string `json:"name,omitempty"`

	// 任务状态。
	Status *string `json:"status,omitempty"`

	// 任务进度，单位：%。
	Process *string `json:"process,omitempty"`

	// 失败原因。
	FailReason *string `json:"fail_reason,omitempty"`
}

func (o TaskDetailResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskDetailResult struct{}"
	}

	return strings.Join([]string{"TaskDetailResult", string(data)}, " ")
}
