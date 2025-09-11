package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAlarmNotificationsRequest Request Object
type UpdateAlarmNotificationsRequest struct {

	// **参数解释**： 告警规则ID **约束限制**： 不涉及。 **取值范围**： 以al开头，后跟22位字母或数字。 **默认取值**： 不涉及。
	AlarmId string `json:"alarm_id"`

	Body *PutAlarmNotificationReq `json:"body,omitempty"`
}

func (o UpdateAlarmNotificationsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAlarmNotificationsRequest struct{}"
	}

	return strings.Join([]string{"UpdateAlarmNotificationsRequest", string(data)}, " ")
}
