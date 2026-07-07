package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ScheduleEventInfoResult **参数解释**: 事件详情。
type ScheduleEventInfoResult struct {

	// **参数解释**: 事件ID。 **取值范围**: 不涉及。
	Id string `json:"id"`

	// **参数解释**: 实例ID。 **取值范围**: 不涉及。
	InstanceId string `json:"instance_id"`

	// **参数解释**: 实例名称。 **取值范围**: 不涉及。
	InstanceName string `json:"instance_name"`

	// **参数解释**: 数据库类型。 **取值范围**: 不涉及。
	DbType string `json:"db_type"`

	// **参数解释**: 创建时间，格式为\"yyyy-mm-ddThh:mm:ssZ\"。其中，T指某个时间的开始；Z指时区偏移量，例如北京时间偏移显示为+0800。 **取值范围**: 不涉及。
	CreatedTime string `json:"created_time"`

	// **参数解释**: 更新时间，格式为\"yyyy-mm-ddThh:mm:ssZ\"。其中，T指某个时间的开始；Z指时区偏移量，例如北京时间偏移显示为+0800。 **取值范围**: 不涉及。
	UpdateTime string `json:"update_time"`

	// **参数解释**: 事件类型。 **取值范围**: - RESTAT_NODE：重启实例节点
	Type string `json:"type"`

	// **参数解释**: 事件对业务的影响。 **取值范围**: 不涉及。
	Impact string `json:"impact"`

	// **参数解释**: 事件状态。 **取值范围**: - WAITING：等待中 - INQUIRING：待授权 - SCHEDULED：待执行 - EXECUTING：执行中 - COMPLETED：已完成 - FAILED：失败 - CANCELED：已取消
	Status string `json:"status"`

	// **参数解释**: 事件发生的原因。 **取值范围**: 不涉及。
	Reason string `json:"reason"`

	ExecutionTimeWindow *ExecuteWindowResult `json:"execution_time_window"`

	// **参数解释**: 事件级别。 **取值范围**: - CRITICAL：紧急 - MAJOR：重要 - MINOR：一般 - INFO：提示
	Level string `json:"level"`

	// **参数解释**: 事件执行时间，格式为\"yyyy-mm-ddThh:mm:ssZ\"。其中，T指某个时间的开始；Z指时区偏移量，例如北京时间偏移显示为+0800。 **取值范围**: 不涉及。
	ExecuteTime *string `json:"execute_time,omitempty"`

	// **参数解释**: 最晚执行时间，事件将在该时间之前执行。格式为\"yyyy-mm-ddThh:mm:ssZ\"。其中，T指某个时间的开始；Z指时区偏移量，例如北京时间偏移显示为+0800。 **取值范围**: 不涉及。
	LatestExecutionTime *string `json:"latest_execution_time,omitempty"`
}

func (o ScheduleEventInfoResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScheduleEventInfoResult struct{}"
	}

	return strings.Join([]string{"ScheduleEventInfoResult", string(data)}, " ")
}
