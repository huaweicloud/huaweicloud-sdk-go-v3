package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExecutionBrief 工作流执行简要信息。
type ExecutionBrief struct {

	// 工作流执行ID。
	ExecutionId *string `json:"execution_id,omitempty"`

	// 工作流执行的创建时间。
	CreatedAt *string `json:"created_at,omitempty"`

	// 工作流状态。
	Status *string `json:"status,omitempty"`

	// 运行的节点。
	RunningSteps *[]string `json:"running_steps,omitempty"`

	// 当前节点。
	CurrentSteps *[]string `json:"current_steps,omitempty"`

	// 运行时长。
	Duration *int64 `json:"duration,omitempty"`
}

func (o ExecutionBrief) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecutionBrief struct{}"
	}

	return strings.Join([]string{"ExecutionBrief", string(data)}, " ")
}
