package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAlarmHistoriesResponse Response Object
type ListAlarmHistoriesResponse struct {

	// alarmHistories列表
	AlarmHistories *[]AlarmHistoryItemV2 `json:"alarm_histories,omitempty"`

	// 告警记录列表总数
	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListAlarmHistoriesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAlarmHistoriesResponse struct{}"
	}

	return strings.Join([]string{"ListAlarmHistoriesResponse", string(data)}, " ")
}
