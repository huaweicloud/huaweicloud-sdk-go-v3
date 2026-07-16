package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WorkflowScheduleResp 工作流调度信息。
type WorkflowScheduleResp struct {

	// **参数解释**：类型，仅支持time（时间）。 **取值范围**：不涉及。
	Type *string `json:"type,omitempty"`

	// **参数解释**：内容。
	Content map[string]string `json:"content,omitempty"`

	// **参数解释**：动作，仅支持run。 **取值范围**：不涉及。
	Action *string `json:"action,omitempty"`

	// **参数解释**：Workflow工作流ID。 **取值范围**：不涉及。
	WorkflowId *string `json:"workflow_id,omitempty"`

	// **参数解释**：用户ID。 **取值范围**：不涉及。
	UserId *string `json:"user_id,omitempty"`

	// **参数解释**：定时调度信息，使能标记。 **取值范围**： - true：生效 - false：不生效
	Enable *bool `json:"enable,omitempty"`

	// **参数解释**：ID标记。 **取值范围**：不涉及。
	Uuid *string `json:"uuid,omitempty"`

	Policies *WorkflowSchedulePoliciesResp `json:"policies,omitempty"`

	// **参数解释**：创建时间。 **取值范围**：不涉及。
	CreatedAt *string `json:"created_at,omitempty"`
}

func (o WorkflowScheduleResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkflowScheduleResp struct{}"
	}

	return strings.Join([]string{"WorkflowScheduleResp", string(data)}, " ")
}
