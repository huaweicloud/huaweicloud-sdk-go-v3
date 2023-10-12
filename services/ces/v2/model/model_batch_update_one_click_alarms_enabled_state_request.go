package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchUpdateOneClickAlarmsEnabledStateRequest Request Object
type BatchUpdateOneClickAlarmsEnabledStateRequest struct {

	// 一键告警ID
	OneClickAlarmId string `json:"one_click_alarm_id"`

	Body *BatchEnableAlarmsRequestBody `json:"body,omitempty"`
}

func (o BatchUpdateOneClickAlarmsEnabledStateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateOneClickAlarmsEnabledStateRequest struct{}"
	}

	return strings.Join([]string{"BatchUpdateOneClickAlarmsEnabledStateRequest", string(data)}, " ")
}
