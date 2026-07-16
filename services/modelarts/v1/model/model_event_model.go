package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EventModel 事件描述信息。
type EventModel struct {

	// **参数描述**：API版本。 **取值范围**：可选值如下： - v1
	ApiVersion string `json:"apiVersion"`

	// **参数描述**：资源类型。 **取值范围**：可选值如下： - Event：事件
	Kind string `json:"kind"`

	// **参数描述**：事件类型。 **取值范围**：可选值如下： - Normal：正常 - Warning：异常
	Type string `json:"type"`

	// **参数描述**：事件第一次出现时间。 **取值范围**：不涉及。
	FirstTimestamp *string `json:"firstTimestamp,omitempty"`

	// **参数描述**：事件最后一次出现时间。 **取值范围**：不涉及。
	LastTimestamp *string `json:"lastTimestamp,omitempty"`

	// **参数描述**：事件连续出现次数。 **取值范围**：不涉及。
	Count *int32 `json:"count,omitempty"`

	// **参数描述**：事件产生的原因。 **取值范围**：不涉及。
	Reason *string `json:"reason,omitempty"`

	// **参数描述**：事件详细信息。 **取值范围**：不涉及。
	Message *string `json:"message,omitempty"`
}

func (o EventModel) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EventModel struct{}"
	}

	return strings.Join([]string{"EventModel", string(data)}, " ")
}
