package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAlarmHandleHistoriesResponse Response Object
type ListAlarmHandleHistoriesResponse struct {

	// 总数量
	Count *int64 `json:"count,omitempty"`

	// 告警工单执行结果
	AlarmHandleHistories *[]AlarmHandleHistory `json:"alarm_handle_histories,omitempty"`

	ResolvedRecord *ResolvedRecordDto `json:"resolved_record,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ListAlarmHandleHistoriesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAlarmHandleHistoriesResponse struct{}"
	}

	return strings.Join([]string{"ListAlarmHandleHistoriesResponse", string(data)}, " ")
}
