package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Event 训练作业事件
type Event struct {

	// 事件信息。
	Message *string `json:"message,omitempty"`

	// 事件级别。
	Level *string `json:"level,omitempty"`

	// 事件发生的时间。
	Time *string `json:"time,omitempty"`

	// **参数解释**：事件来源。 **取值范围**：不涉及。
	Source *string `json:"source,omitempty"`
}

func (o Event) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Event struct{}"
	}

	return strings.Join([]string{"Event", string(data)}, " ")
}
