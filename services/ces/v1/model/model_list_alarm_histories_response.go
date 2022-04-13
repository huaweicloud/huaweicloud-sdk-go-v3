package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListAlarmHistoriesResponse struct {
	// 一条或者多条告警历史详细信息

	AlarmHistories *[]AlarmHistoryInfo `json:"alarm_histories,omitempty"`

	MetaData       *MetaDataForAlarmHistory `json:"meta_data,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o ListAlarmHistoriesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAlarmHistoriesResponse struct{}"
	}

	return strings.Join([]string{"ListAlarmHistoriesResponse", string(data)}, " ")
}
