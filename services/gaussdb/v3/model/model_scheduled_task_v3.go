package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ScheduledTaskV3 定时任务详情信息
type ScheduledTaskV3 struct {

	// **参数解释**：  任务创建时间。  **取值范围**： 不涉及。
	CreateTime *int64 `json:"create_time,omitempty"`

	// **参数解释**：  任务结束时间。  **取值范围**： 不涉及。
	EndTime *int64 `json:"end_time,omitempty"`

	// **参数解释**：  任务绑定的实例ID。  **取值范围**： 不涉及。
	InstanceId *string `json:"instance_id,omitempty"`

	// **参数解释**：  租户项目ID。  **取值范围**： 不涉及。
	ProjectId *string `json:"project_id,omitempty"`

	// **参数解释**：  任务使用的数据库代理ID。  **取值范围**： 不涉及。
	ProxyId *string `json:"proxy_id,omitempty"`

	// **参数解释**：  任务使用的数据库代理名称。  **取值范围**： 不涉及。
	ProxyName *string `json:"proxy_name,omitempty"`

	// **参数解释**：  任务开始时间。  **取值范围**： 不涉及。
	StartTime *int64 `json:"start_time,omitempty"`

	// **参数解释**：  任务ID，此参数是任务的唯一标识。  **取值范围**： 不涉及。
	TaskId *string `json:"task_id,omitempty"`

	// **参数解释**：  任务名称。  **取值范围**： 不涉及。
	TaskName *string `json:"task_name,omitempty"`

	// **参数解释**：  任务状态。  **取值范围**： - RUNNING：任务正在执行。 - SUCCESS：任务执行成功。 - FAIL：任务执行失败。 - CANCELED：任务被取消。 - WAITING：任务等待执行。
	TaskStatus *string `json:"task_status,omitempty"`
}

func (o ScheduledTaskV3) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScheduledTaskV3 struct{}"
	}

	return strings.Join([]string{"ScheduledTaskV3", string(data)}, " ")
}
