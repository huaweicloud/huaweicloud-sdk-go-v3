package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TaskLogsVo struct {

	// **参数解释**：  分片变更任务ID。  **参数范围**：  不涉及。
	TaskId *string `json:"task_id,omitempty"`

	// **参数解释**：  DDM实例ID。  **参数范围**：  不涉及。
	DdmInstanceId *string `json:"ddm_instance_id,omitempty"`

	// **参数解释**：  等级。  **参数范围**：  不涉及。
	Level *string `json:"level,omitempty"`

	// **参数解释**：  创建时间。  **参数范围**：  不涉及。
	CreatedTime float32 `json:"created_time,omitempty"`

	// **参数解释**：  消息。  **参数范围**：  不涉及。
	Message *string `json:"message,omitempty"`
}

func (o TaskLogsVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskLogsVo struct{}"
	}

	return strings.Join([]string{"TaskLogsVo", string(data)}, " ")
}
