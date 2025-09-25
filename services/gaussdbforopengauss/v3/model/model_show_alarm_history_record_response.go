package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAlarmHistoryRecordResponse Response Object
type ShowAlarmHistoryRecordResponse struct {

	// **参数解释**： 总数量。 **默认取值**： 不涉及。
	TotalCount *int32 `json:"total_count,omitempty"`

	// **参数解释**： 告警记录列表。
	HistoryRecords *[]AlarmHistoryRecordResult `json:"history_records,omitempty"`
	HttpStatusCode int                         `json:"-"`
}

func (o ShowAlarmHistoryRecordResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAlarmHistoryRecordResponse struct{}"
	}

	return strings.Join([]string{"ShowAlarmHistoryRecordResponse", string(data)}, " ")
}
