package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAlarmNotifyResponse Response Object
type ListAlarmNotifyResponse struct {

	// 告警通知列表。
	FrontAlarmNotifyResults *[]FrontAlarmNotifyResult `json:"front_alarm_notify_results,omitempty"`

	// 消息总数。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListAlarmNotifyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAlarmNotifyResponse struct{}"
	}

	return strings.Join([]string{"ListAlarmNotifyResponse", string(data)}, " ")
}
