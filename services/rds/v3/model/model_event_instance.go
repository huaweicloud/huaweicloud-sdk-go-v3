package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type EventInstance struct {

	// **参数解释**：  事件id  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	EventId string `json:"event_id"`

	// **参数解释**：  实例id  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	InstanceId string `json:"instance_id"`
}

func (o EventInstance) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EventInstance struct{}"
	}

	return strings.Join([]string{"EventInstance", string(data)}, " ")
}
