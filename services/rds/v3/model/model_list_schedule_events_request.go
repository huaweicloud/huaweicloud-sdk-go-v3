package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListScheduleEventsRequest Request Object
type ListScheduleEventsRequest struct {

	// **参数解释**：  事件ID。  **约束限制**：  不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释**：  实例ID。  **约束限制**：  不涉及。
	InstanceId *string `json:"instance_id,omitempty"`

	// **参数解释**：  事件状态  **约束限制**：  不涉及。  **取值范围**：  枚举值：WAITING 等待中；INQUIRING 待授权; SCHEDULED 待执行; EXECUTING 执行中; COMPLETED 已完成; FAILED 失败; CANCELED 已取消。
	Status *string `json:"status,omitempty"`

	// **参数解释**：  事件类型  **约束限制**：  不涉及。  **取值范围**：  枚举值：RESTAT_NODE 重启实例节点。
	Type *string `json:"type,omitempty"`

	// **参数解释**：  事件级别  **约束限制**：  不涉及。  **取值范围**：  CRITICAL 紧急；MAJOR 重要；MINOR 一般；INFO 提示。
	Level *string `json:"level,omitempty"`

	// **参数解释**：  排序字段，支持planned_execution_time、created_time、latest_execution_time。  **约束限制**：  不涉及。
	SortField *string `json:"sort_field,omitempty"`

	// **参数解释**：  排序顺序  **约束限制**：  不涉及。  **取值范围**：  DESC 降序；ASC升序   **默认取值**：  DESC 降序
	Order *string `json:"order,omitempty"`

	// **参数解释**：  索引位置，偏移量。从第一条数据偏移offset条数据后开始查询，默认为0（偏移0条数据，表示从第一条数据开始查询），必须为数字，不能为负数。  **约束限制**：  不涉及。  **取值范围**：  不涉及  **默认取值**：  0
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**：  查询记录数。默认为10，不能为负数，最小值为1，最大值为100。  **约束限制**：  不涉及。  **取值范围**：  不涉及  **默认取值**：  10
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListScheduleEventsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListScheduleEventsRequest struct{}"
	}

	return strings.Join([]string{"ListScheduleEventsRequest", string(data)}, " ")
}
