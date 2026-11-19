package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CancelColdTableWaitingTasksReq struct {

	// **参数解释**：  取消目标列表。  **约束限制**：  列表长度不超过50。
	CancelTasks []CancelColdTableWaitingTask `json:"cancel_tasks"`
}

func (o CancelColdTableWaitingTasksReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CancelColdTableWaitingTasksReq struct{}"
	}

	return strings.Join([]string{"CancelColdTableWaitingTasksReq", string(data)}, " ")
}
