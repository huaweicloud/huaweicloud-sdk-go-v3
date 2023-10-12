package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchUpdateOneClickAlarmPoliciesEnabledStateRequest Request Object
type BatchUpdateOneClickAlarmPoliciesEnabledStateRequest struct {

	// 一键告警ID
	OneClickAlarmId string `json:"one_click_alarm_id"`

	// 告警规则ID
	AlarmId string `json:"alarm_id"`

	Body *BatchEnableAlarmPoliciesRequestBody `json:"body,omitempty"`
}

func (o BatchUpdateOneClickAlarmPoliciesEnabledStateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateOneClickAlarmPoliciesEnabledStateRequest struct{}"
	}

	return strings.Join([]string{"BatchUpdateOneClickAlarmPoliciesEnabledStateRequest", string(data)}, " ")
}
