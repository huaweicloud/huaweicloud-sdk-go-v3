package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type WorkflowTask struct {
	Input *ObsInfo `json:"input,omitempty"`

	// 工作流任务结果列表
	TaskResult *[]ObjectTaskResult `json:"task_result,omitempty"`
}

func (o WorkflowTask) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkflowTask struct{}"
	}

	return strings.Join([]string{"WorkflowTask", string(data)}, " ")
}
