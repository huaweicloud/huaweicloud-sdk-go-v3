package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListOneClickAlarmRulesResponse Response Object
type ListOneClickAlarmRulesResponse struct {

	// 告警规则列表
	Alarms         *[]ListAlarmsRespAlarms `json:"alarms,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o ListOneClickAlarmRulesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListOneClickAlarmRulesResponse struct{}"
	}

	return strings.Join([]string{"ListOneClickAlarmRulesResponse", string(data)}, " ")
}
