package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTaskDetailsResponse Response Object
type ShowTaskDetailsResponse struct {

	// **参数解释**：  实例ID。  **取值范围**：  不涉及。
	InstanceId *string `json:"instance_id,omitempty"`

	// **参数解释**：  实例名称。  **取值范围**：  不涉及。
	InstanceName *string `json:"instance_name,omitempty"`

	// **参数解释**：  任务ID。  **取值范围**：  不涉及。
	JobId *string `json:"job_id,omitempty"`

	// **参数解释**：  任务名称。  **取值范围**：  不涉及。
	JobName *string `json:"job_name,omitempty"`

	// **参数解释**：  任务状态。  **取值范围**：    - Pending：表示待执行。   - Running：表示运行中。   - Completed：表示已完成。
	Status *string `json:"status,omitempty"`

	// **参数解释**：  任务详情列表。  **取值范围**：  不涉及。
	SubTaskList    *[]SubTaskInfo `json:"sub_task_list,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ShowTaskDetailsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTaskDetailsResponse struct{}"
	}

	return strings.Join([]string{"ShowTaskDetailsResponse", string(data)}, " ")
}
