package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HpaEventResponse 自动扩缩容事件返回体
type HpaEventResponse struct {

	// **参数解释：** 自动扩缩容策略事件ID **取值范围：** 事件ID
	Id *string `json:"id,omitempty"`

	// **参数解释：** 自动扩缩容策略ID **取值范围：** 策略ID
	HpaId *string `json:"hpa_id,omitempty"`

	// **参数解释：** 自动扩缩容事件状态。 **取值范围：** - SUCCESS: 成功 - FAILED: 失败
	Status *string `json:"status,omitempty"`

	// **参数解释：** 自动扩缩容规则执行信息。 **取值范围：** 不涉及
	Message *string `json:"message,omitempty"`

	// **参数解释：** 扩缩容前实例数。 **取值范围：** 不涉及。
	CurrentReplicas *int32 `json:"current_replicas,omitempty"`

	// **参数解释：** 预设目标实例数。 **取值范围：** 不涉及。
	TargetReplicas *int32 `json:"target_replicas,omitempty"`

	// **参数解释：** 扩缩容后实例数。 **取值范围：** 不涉及。
	FinalReplicas *int32 `json:"final_replicas,omitempty"`

	// **参数解释：** 执行记录时间。 **取值范围：** 2025-05-20 10:05:55
	RecordTime *string `json:"record_time,omitempty"`
}

func (o HpaEventResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HpaEventResponse struct{}"
	}

	return strings.Join([]string{"HpaEventResponse", string(data)}, " ")
}
