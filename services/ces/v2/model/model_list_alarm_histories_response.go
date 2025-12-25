package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAlarmHistoriesResponse Response Object
type ListAlarmHistoriesResponse struct {

	// **参数解释**： 告警记录列表。
	AlarmHistories *[]AlarmHistoryItemV2 `json:"alarm_histories,omitempty"`

	// **参数解释**： 告警记录总数。 **取值范围**： [0,2147483647]
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
