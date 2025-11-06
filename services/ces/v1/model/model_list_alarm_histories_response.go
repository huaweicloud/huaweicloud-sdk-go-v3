package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAlarmHistoriesResponse Response Object
type ListAlarmHistoriesResponse struct {

	// **参数解释**： 一条或者多条告警历史详细信息
	AlarmHistories *[]AlarmHistoryInfoResp `json:"alarm_histories,omitempty"`

	MetaData       *MetaDataForAlarmHistoryResp `json:"meta_data,omitempty"`
	HttpStatusCode int                          `json:"-"`
}

func (o ListAlarmHistoriesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAlarmHistoriesResponse struct{}"
	}

	return strings.Join([]string{"ListAlarmHistoriesResponse", string(data)}, " ")
}
