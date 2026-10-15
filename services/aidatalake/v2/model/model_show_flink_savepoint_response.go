package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowFlinkSavepointResponse Response Object
type ShowFlinkSavepointResponse struct {

	// **参数解释**：Savepoint唯一标识。 **取值范围**：UUIDv7格式，固定36位，由英文小写字母、数字及中划线组成。
	SavepointId *string `json:"savepoint_id,omitempty"`

	// **参数解释**：关联的Flink作业ID。 **取值范围**：UUIDv7格式，固定36位，由英文小写字母、数字及中划线组成。
	JobId *string `json:"job_id,omitempty"`

	TriggerType *FlinkSavepointTriggerType `json:"trigger_type,omitempty"`

	Format *FlinkSavepointFormat `json:"format,omitempty"`

	Status *FlinkSavepointStatus `json:"status,omitempty"`

	// **参数解释**：是否启用了Drain模式。启用后，触发Savepoint前将先发送Watermark到所有算子以触发窗口计算完成，然后再保存状态快照。 **取值范围**：   - true：已启用Drain模式。   - false：未启用Drain模式。
	DrainModeEnabled *bool `json:"drain_mode_enabled,omitempty"`

	// **参数解释**：Savepoint存储路径。 **取值范围**：1~512个字符，需以obs://开头。
	Location *string `json:"location,omitempty"`

	// **参数解释**：创建时间戳。 **取值范围**：0~4102416000000，13位毫秒时间戳。
	CreatedTime *int64 `json:"created_time,omitempty"`

	// **参数解释**：完成时间戳。 **取值范围**：0~4102416000000，13位毫秒时间戳。
	CompletedTime *int64 `json:"completed_time,omitempty"`

	CreateUser *FlinkCreateUser `json:"create_user,omitempty"`

	// **参数解释**：作业快照描述。 **取值范围**：0~512个字符。
	Description    *string `json:"description,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowFlinkSavepointResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFlinkSavepointResponse struct{}"
	}

	return strings.Join([]string{"ShowFlinkSavepointResponse", string(data)}, " ")
}
