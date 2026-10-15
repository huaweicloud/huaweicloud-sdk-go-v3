package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FlinkJobEvent **参数解释**：描述Flink事件信息。
type FlinkJobEvent struct {

	// **参数解释**：运行事件ID。 **取值范围**：UUIDv7格式，固定36位，由英文小写字母、数字及中划线组成。
	EventId string `json:"event_id"`

	// **参数解释**：事件创建时间戳。 **取值范围**：0~4102416000000，13位毫秒时间戳。
	CreateTime int64 `json:"create_time"`

	// **参数解释**：运行事件名称。 **取值范围**：1~128个字符。
	EventName string `json:"event_name"`

	// **参数解释**：运行事件信息。 **取值范围**：0~1024个字符。
	EventMessage string `json:"event_message"`
}

func (o FlinkJobEvent) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FlinkJobEvent struct{}"
	}

	return strings.Join([]string{"FlinkJobEvent", string(data)}, " ")
}
