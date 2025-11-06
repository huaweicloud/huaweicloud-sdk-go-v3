package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateOneClickAlarmNotificationsRequest Request Object
type UpdateOneClickAlarmNotificationsRequest struct {

	// **参数解释**： 一键告警ID **约束限制**： 不涉及。 **取值范围**： 长度为[1,64]个字符。 **默认取值**： 不涉及。
	OneClickAlarmId string `json:"one_click_alarm_id"`

	Body *UpdateOneClickAlarmNotificationsRequestBody `json:"body,omitempty"`
}

func (o UpdateOneClickAlarmNotificationsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateOneClickAlarmNotificationsRequest struct{}"
	}

	return strings.Join([]string{"UpdateOneClickAlarmNotificationsRequest", string(data)}, " ")
}
