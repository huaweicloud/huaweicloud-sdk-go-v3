package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SubTaskInfo 子任务信息。
type SubTaskInfo struct {

	// **参数解释**：  子任务名。  **取值范围**：  不涉及。
	SubTaskName *string `json:"sub_task_name,omitempty"`

	// **参数解释**：  子任务进度百分比。  **取值范围**：  0-100。
	Percent *string `json:"percent,omitempty"`

	// **参数解释**：  子任务状态。  **取值范围**：  - Pending：表示待执行。 - Running：表示运行中。 - Completed：表示已完成。
	Status *string `json:"status,omitempty"`

	// **参数解释**：  剩余时间，单位为秒。  **取值范围**：  不涉及。
	RemainingTime *string `json:"remaining_time,omitempty"`
}

func (o SubTaskInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubTaskInfo struct{}"
	}

	return strings.Join([]string{"SubTaskInfo", string(data)}, " ")
}
