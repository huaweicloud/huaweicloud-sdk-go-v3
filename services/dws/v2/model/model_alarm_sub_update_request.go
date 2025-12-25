package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AlarmSubUpdateRequest struct {

	// **参数解释**： 告警订阅名称。 **取值范围**： 不涉及。
	Name string `json:"name"`

	// **参数解释**： 是否开启订阅。 **取值范围**： 不涉及。
	Enable *int32 `json:"enable,omitempty"`

	// **参数解释**： 告警级别。 **取值范围**： 不涉及。
	AlarmLevel *string `json:"alarm_level,omitempty"`

	// **参数解释**： 消息主题地址。 **取值范围**： 不涉及。
	NotificationTarget string `json:"notification_target"`

	// **参数解释**： 消息主题名称。 **取值范围**： 不涉及。
	NotificationTargetName string `json:"notification_target_name"`

	// **参数解释**： 消息主题类型。 **取值范围**： - SMN：SMN类型
	NotificationTargetType string `json:"notification_target_type"`
}

func (o AlarmSubUpdateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlarmSubUpdateRequest struct{}"
	}

	return strings.Join([]string{"AlarmSubUpdateRequest", string(data)}, " ")
}
