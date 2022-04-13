package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListAlarmsResponse struct {
	// 告警详情

	Alarms *[]ListAlarmResponseBodyAlarms `json:"alarms,omitempty"`
	// 告警列表总数

	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListAlarmsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAlarmsResponse struct{}"
	}

	return strings.Join([]string{"ListAlarmsResponse", string(data)}, " ")
}
