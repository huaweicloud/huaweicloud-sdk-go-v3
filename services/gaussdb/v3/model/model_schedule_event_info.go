package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ScheduleEventInfo **参数解释**：  事件详情。  **约束限制**：  不涉及。
type ScheduleEventInfo struct {

	// **参数解释**：  事件ID。  **取值范围**：  只能由英文字母、数字组成，前面为UUID，后缀为ev07，长度为36个字符。
	Id *string `json:"id,omitempty"`

	// **参数解释**：  事件类别。  **取值范围**：  Maintenance：计划内运维事件。
	Category *string `json:"category,omitempty"`

	// **参数解释**：  事件影响。  **取值范围**：  不涉及。
	Impact *string `json:"impact,omitempty"`

	// **参数解释**：  事件状态。  **取值范围**：    - inquiring：待授权。   - scheduled：待执行。   - executing：执行中。   - completed：执行完成。   - canceled：事件关闭。   - failed：执行失败。
	Status *string `json:"status,omitempty"`

	// **参数解释**：  事件原因。  **取值范围**：  不涉及。
	Reason *string `json:"reason,omitempty"`

	// **参数解释**：  事件级别。  **取值范围**：  - critical：紧急。 - major：重要。 - minor：一般。 - info：提示。
	Level *string `json:"level,omitempty"`

	// **参数解释**：  实例ID。  **取值范围**：  只能由英文字母、数字组成，前面为UUID，后缀为in07，长度为36个字符。
	InstanceId *string `json:"instance_id,omitempty"`

	// **参数解释**：  实例名称。  **取值范围**：  最小为4个字符，最大为64个字符且不超过64个字节（注意：一个中文字符占用3个字节），必须以字母或中文开头，区分大小写，可以包含字母、数字、中划线、下划线或中文，不能包含其他特殊字符。
	InstanceName *string `json:"instance_name,omitempty"`

	// **参数解释**：  引擎名称。  **取值范围**：  taurus：TaurusDB企业版。
	DbType *string `json:"db_type,omitempty"`

	// **参数解释**：  创建时间。UTC，格式为“yyyy-mm-ddThh:mm:ssZ”。 其中，T指某个时间的开始；Z指时区偏移量，例如偏移1个小时显示为+0100。  **取值范围**：  不涉及。
	CreatedTime *string `json:"created_time,omitempty"`

	// **参数解释**：  更新时间。UTC，格式为“yyyy-mm-ddThh:mm:ssZ”。 其中，T指某个时间的开始；Z指时区偏移量，例如偏移1个小时显示为+0100。  **取值范围**：  不涉及。
	UpdatedTime *string `json:"updated_time,omitempty"`

	// **参数解释**：  事件类型。  **取值范围**：  - system.lifecycle.rebuild_node：备机重建事件。 - system.lifecycle.db_upgrade：数据库内核小版本升级事件。 - system.scheduled_event.high_cpu_memory：实例CPU或内存高负载事件，需要变更实例规格。
	Type *string `json:"type,omitempty"`

	// **参数解释**：  扩展信息。  **取值范围**：  不涉及。
	ExtendInfo *string `json:"extend_info,omitempty"`

	// **参数解释**：  事件的执行时间。UTC，格式为“yyyy-mm-ddThh:mm:ssZ”。 其中，T指某个时间的开始；Z指时区偏移量，例如偏移1个小时显示为+0100。  **取值范围**：  不涉及。
	ExecuteTime *string `json:"execute_time,omitempty"`

	// **参数解释**：  事件执行窗口。
	ExecutionTimeWindow *interface{} `json:"execution_time_window,omitempty"`

	// **参数解释**：  事件对象信息列表，包含事件对象ID和事件对象的执行状态
	EventEntities *[]EventEntity `json:"event_entities,omitempty"`
}

func (o ScheduleEventInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScheduleEventInfo struct{}"
	}

	return strings.Join([]string{"ScheduleEventInfo", string(data)}, " ")
}
