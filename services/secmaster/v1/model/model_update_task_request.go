package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateTaskRequest Request Object
type UpdateTaskRequest struct {

	// **参数解释：** 内容类型 - application/json;charset=UTF-8    普通API请求的类型  **约束限制：** 不涉及 **取值范围：** - application/json;charset=UTF-8  **默认取值：** 不涉及
	ContentType string `json:"Content-Type"`

	// **参数解释**: 工作空间ID **取值范围**: 不涉及
	WorkspaceId string `json:"workspace_id"`

	// **参数解释**: 待办任务的ID **取值范围**: 不涉及
	TaskId string `json:"task_id"`

	Body *ModifyTaskInfo `json:"body,omitempty"`
}

func (o UpdateTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTaskRequest struct{}"
	}

	return strings.Join([]string{"UpdateTaskRequest", string(data)}, " ")
}
