package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EventEntity **参数解释**：  事件对象信息。
type EventEntity struct {

	// **参数解释**：  事件对象ID。  **取值范围**：  实例ID或者节点ID。只能由英文字母、数字组成，后缀为in07或no07，长度为36个字符。
	EventEntityId *string `json:"event_entity_id,omitempty"`

	// **参数解释**：  事件对象的执行状态。  **取值范围**：    - inquiring：待授权。   - scheduled：待执行。   - executing：执行中。   - completed：执行完成。   - canceled：事件关闭。   - failed：执行失败。
	EventEntityStatus *string `json:"event_entity_status,omitempty"`
}

func (o EventEntity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EventEntity struct{}"
	}

	return strings.Join([]string{"EventEntity", string(data)}, " ")
}
