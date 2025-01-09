package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAlarmsResponse Response Object
type ListAlarmsResponse struct {

	// alarmHistories列表
	AlarmHistories *[]AlarmHistoryItem `json:"alarm_histories,omitempty"`

	// 告警记录列表总数
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
