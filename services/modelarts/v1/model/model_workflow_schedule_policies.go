package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WorkflowSchedulePolicies 工作流调度策略。
type WorkflowSchedulePolicies struct {

	// 定时调度策略中的标记，失败时触发。
	OnFailure *string `json:"on_failure,omitempty"`

	// 定时调度策略中的标记，running时触发。
	OnRunning *string `json:"on_running,omitempty"`
}

func (o WorkflowSchedulePolicies) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkflowSchedulePolicies struct{}"
	}

	return strings.Join([]string{"WorkflowSchedulePolicies", string(data)}, " ")
}
