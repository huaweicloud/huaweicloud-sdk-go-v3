package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListColdTableWaitingTasksResponse Response Object
type ListColdTableWaitingTasksResponse struct {

	// **参数解释**：  任务总数。  **取值范围**：  非负整数。
	TotalCount *int32 `json:"total_count,omitempty"`

	// **参数解释**：  冷热分离待下发任务列表，没有任务时默认为空数组。
	WaitingTasks   *[]ColdTableWaitingTask `json:"waiting_tasks,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o ListColdTableWaitingTasksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListColdTableWaitingTasksResponse struct{}"
	}

	return strings.Join([]string{"ListColdTableWaitingTasksResponse", string(data)}, " ")
}
