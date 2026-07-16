package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateWorkflowScheduleResponse Response Object
type CreateWorkflowScheduleResponse struct {

	// 类型，仅支持time（时间）。
	Type *string `json:"type,omitempty"`

	// 内容。
	Content map[string]interface{} `json:"content,omitempty"`

	// 动作，仅支持run。
	Action *string `json:"action,omitempty"`

	// Workflow工作流ID。
	WorkflowId *string `json:"workflow_id,omitempty"`

	// 用户ID。
	UserId *string `json:"user_id,omitempty"`

	// 定时调度信息，使能标记。
	Enable *bool `json:"enable,omitempty"`

	// ID标记。
	Uuid *string `json:"uuid,omitempty"`

	Policies *WorkflowSchedulePolicies `json:"policies,omitempty"`

	// 创建时间。
	CreatedAt      *string `json:"created_at,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateWorkflowScheduleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateWorkflowScheduleResponse struct{}"
	}

	return strings.Join([]string{"CreateWorkflowScheduleResponse", string(data)}, " ")
}
