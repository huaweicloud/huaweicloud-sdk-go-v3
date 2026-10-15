package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FlinkSavepointSummary **参数解释**：Flink Savepoint摘要信息（列表展示）。
type FlinkSavepointSummary struct {

	// **参数解释**：Savepoint唯一标识。 **取值范围**：UUIDv7格式，固定36位，由英文小写字母、数字及中划线组成。
	SavepointId *string `json:"savepoint_id,omitempty"`

	// **参数解释**：关联的Flink作业ID。 **取值范围**：UUIDv7格式，固定36位，由英文小写字母、数字及中划线组成。
	JobId *string `json:"job_id,omitempty"`

	TriggerType *FlinkSavepointTriggerType `json:"trigger_type,omitempty"`

	Status *FlinkSavepointStatus `json:"status,omitempty"`

	// **参数解释**：创建时间戳。 **取值范围**：0~4102416000000，13位毫秒时间戳。
	CreatedTime *int64 `json:"created_time,omitempty"`

	// **参数解释**：快照存储路径。 **取值范围**：1~512个字符，需以obs://开头。
	Location *string `json:"location,omitempty"`

	// **参数解释**：作业快照描述。 **取值范围**：0~512个字符。
	Description *string `json:"description,omitempty"`
}

func (o FlinkSavepointSummary) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FlinkSavepointSummary struct{}"
	}

	return strings.Join([]string{"FlinkSavepointSummary", string(data)}, " ")
}
