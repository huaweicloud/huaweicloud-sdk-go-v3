package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAlarmResponse Response Object
type ShowAlarmResponse struct {

	// **参数解释**： 告警对象列表。
	MetricAlarms   *[]MetricAlarmsResp `json:"metric_alarms,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o ShowAlarmResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAlarmResponse struct{}"
	}

	return strings.Join([]string{"ShowAlarmResponse", string(data)}, " ")
}
