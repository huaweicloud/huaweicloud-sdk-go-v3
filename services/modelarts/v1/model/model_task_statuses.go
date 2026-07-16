package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TaskStatuses 训练在子任务状态信息。
type TaskStatuses struct {

	// 训练作业子任务名称。
	Task *string `json:"task,omitempty"`

	// 训练作业子任务退出码。
	ExitCode *int32 `json:"exit_code,omitempty"`

	// 训练作业子任务错误消息。
	Message *string `json:"message,omitempty"`
}

func (o TaskStatuses) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskStatuses struct{}"
	}

	return strings.Join([]string{"TaskStatuses", string(data)}, " ")
}
