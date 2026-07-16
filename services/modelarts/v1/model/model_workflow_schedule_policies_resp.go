package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WorkflowSchedulePoliciesResp 工作流调度策略。
type WorkflowSchedulePoliciesResp struct {

	// **参数解释**：定时调度策略中的标记，失败时触发。 **取值范围**：不涉及。
	OnFailure *string `json:"on_failure,omitempty"`

	// **参数解释**：定时调度策略中的标记，running时触发。 **取值范围**：不涉及。
	OnRunning *string `json:"on_running,omitempty"`
}

func (o WorkflowSchedulePoliciesResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkflowSchedulePoliciesResp struct{}"
	}

	return strings.Join([]string{"WorkflowSchedulePoliciesResp", string(data)}, " ")
}
