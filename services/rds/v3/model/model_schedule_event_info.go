package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ScheduleEventInfo 事件详情
type ScheduleEventInfo struct {

	// **参数解释**：  事件ID。  **约束限制**：  不涉及。
	Id string `json:"id"`

	// **参数解释**：  实例ID。  **约束限制**：  不涉及。
	InstanceId string `json:"instance_id"`

	// **参数解释**：  实例名称。  **约束限制**：  不涉及。
	InstanceName string `json:"instance_name"`

	// **参数解释**：  数据库类型。  **约束限制**：  不涉及。
	DbType string `json:"db_type"`

	// **参数解释**：  事件创建的时间。  **约束限制**：  不涉及。
	CreatedTime string `json:"created_time"`

	// **参数解释**：  事件更新的时间。  **约束限制**：  不涉及。
	UpdateTime string `json:"update_time"`

	// **参数解释**：  事件类型。  **约束限制**：  不涉及。
	Type string `json:"type"`

	// **参数解释**：  事件对系统的影响。  **约束限制**：  不涉及。
	Impact string `json:"impact"`

	// **参数解释**：  事件状态。  **约束限制**：  不涉及。
	Status string `json:"status"`

	// **参数解释**：  事件发生的原因。  **约束限制**：  不涉及。
	Reason string `json:"reason"`

	// **参数解释**：  事件的严重级别。  **约束限制**：  不涉及。
	Level string `json:"level"`

	// **参数解释**：  事件执行的时间。  **约束限制**：  不涉及。
	ExecuteTime *string `json:"execute_time,omitempty"`

	// **参数解释**：  事件最近一次执行的时间。  **约束限制**：  不涉及。
	LatestExecutionTime *string `json:"latest_execution_time,omitempty"`
}

func (o ScheduleEventInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScheduleEventInfo struct{}"
	}

	return strings.Join([]string{"ScheduleEventInfo", string(data)}, " ")
}
