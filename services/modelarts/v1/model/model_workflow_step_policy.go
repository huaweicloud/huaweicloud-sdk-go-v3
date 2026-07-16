package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WorkflowStepPolicy 节点执行策略。
type WorkflowStepPolicy struct {

	// 节点执行间隔。
	PollIntervalSeconds *string `json:"poll_interval_seconds,omitempty"`

	// 最大执行时间。
	MaxExecutionMinutes *string `json:"max_execution_minutes,omitempty"`
}

func (o WorkflowStepPolicy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkflowStepPolicy struct{}"
	}

	return strings.Join([]string{"WorkflowStepPolicy", string(data)}, " ")
}
