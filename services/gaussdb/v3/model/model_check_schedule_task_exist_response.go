package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckScheduleTaskExistResponse Response Object
type CheckScheduleTaskExistResponse struct {

	// **参数解释**：  定时任务类型是否存在。 **取值范围**： - true：指定的定时任务类型已存在。 - false：指定的定时任务类型不存在。
	Exist *bool `json:"exist,omitempty"`

	// **参数解释**：  定时任务详情列表。当 `exist` 为 true 时，此列表包含已存在的任务信息。  **取值范围**： 不涉及。
	ScheduledTasks *[]ScheduledTaskV3 `json:"scheduled_tasks,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o CheckScheduleTaskExistResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckScheduleTaskExistResponse struct{}"
	}

	return strings.Join([]string{"CheckScheduleTaskExistResponse", string(data)}, " ")
}
