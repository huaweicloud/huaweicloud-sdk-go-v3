package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type JobSpec struct {

	// 任务进度。
	Progress float32 `json:"progress,omitempty"`

	// 任务状态。
	Status *string `json:"status,omitempty"`

	// 子任务。
	Tasks *[]Task `json:"tasks,omitempty"`
}

func (o JobSpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobSpec struct{}"
	}

	return strings.Join([]string{"JobSpec", string(data)}, " ")
}
