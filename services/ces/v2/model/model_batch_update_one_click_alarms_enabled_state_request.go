package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchUpdateOneClickAlarmsEnabledStateRequest Request Object
type BatchUpdateOneClickAlarmsEnabledStateRequest struct {

	// **参数解释**： 一键告警ID。 **约束限制**： 不涉及。 **取值范围**： 只能为字母或者数字，字符长度为[1,64] **默认取值**： 不涉及。
	OneClickAlarmId string `json:"one_click_alarm_id"`

	Body *BatchUpdateOneClickAlarmsEnabledStateRequestBody `json:"body,omitempty"`
}

func (o BatchUpdateOneClickAlarmsEnabledStateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateOneClickAlarmsEnabledStateRequest struct{}"
	}

	return strings.Join([]string{"BatchUpdateOneClickAlarmsEnabledStateRequest", string(data)}, " ")
}
