package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EventInstances **参数解释**：  事件信息。  **约束限制**：  不涉及。
type EventInstances struct {

	// **参数解释**：  事件ID。  获取方法请参见[获取事件列表](https://support.huaweicloud.com/api-taurusdb/ShowInstanceScheduleEvents.html)。  **约束限制**：  不涉及。  **取值范围**：  只能由英文字母、数字组成，前面为UUID，后缀为ev07，长度为36个字符。  **默认取值**：  不涉及。
	EventId string `json:"event_id"`
}

func (o EventInstances) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EventInstances struct{}"
	}

	return strings.Join([]string{"EventInstances", string(data)}, " ")
}
