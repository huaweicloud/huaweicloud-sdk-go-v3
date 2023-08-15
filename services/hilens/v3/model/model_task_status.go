package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TaskStatus 作业在实例上的状态信息
type TaskStatus struct {

	// 作业运行失败原因
	Cause *string `json:"cause,omitempty"`

	// 实例id
	PodId *string `json:"pod_id,omitempty"`

	// 实例名称
	PodName *string `json:"pod_name,omitempty"`

	// 作业在实例上的状态
	TaskStatus *string `json:"task_status,omitempty"`
}

func (o TaskStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskStatus struct{}"
	}

	return strings.Join([]string{"TaskStatus", string(data)}, " ")
}
