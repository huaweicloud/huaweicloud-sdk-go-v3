package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAlarmSubResponse Response Object
type DeleteAlarmSubResponse struct {

	// **参数解释**： 告警订阅ID。 **取值范围**： 不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释**： 告警订阅名称。 **取值范围**： 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**： 是否开启订阅。 **取值范围**： 不涉及。
	Enable *int32 `json:"enable,omitempty"`

	// **参数解释**： 告警级别。 **取值范围**： 不涉及。
	AlarmLevel *string `json:"alarm_level,omitempty"`

	// **参数解释**： 项目ID。 **取值范围**： 不涉及。
	ProjectId *string `json:"project_id,omitempty"`

	// **参数解释**： 所属服务。 **取值范围**： 不涉及。
	NameSpace *string `json:"name_space,omitempty"`

	// **参数解释**： 消息主题地址。 **取值范围**： 不涉及。
	NotificationTarget *string `json:"notification_target,omitempty"`

	// **参数解释**： 消息主题名称。 **取值范围**： 不涉及。
	NotificationTargetName *string `json:"notification_target_name,omitempty"`

	// **参数解释**： 消息主题类型。 **取值范围**： - SMN：SMN类型
	NotificationTargetType *string `json:"notification_target_type,omitempty"`

	// **参数解释**： 语言。 **取值范围**： 不涉及。
	Language *string `json:"language,omitempty"`

	// **参数解释**： 时区。 **取值范围**： 不涉及。
	TimeZone       *string `json:"time_zone,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteAlarmSubResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAlarmSubResponse struct{}"
	}

	return strings.Join([]string{"DeleteAlarmSubResponse", string(data)}, " ")
}
