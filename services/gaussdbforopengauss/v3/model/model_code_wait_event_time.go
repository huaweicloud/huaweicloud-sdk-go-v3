package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CodeWaitEventTime struct {

	// **参数解释**: 总耗时（单位：微秒）。 **取值范围**: 不涉及。
	AllTime int64 `json:"all_time"`

	CodeWaitEventTimeDetails *EventTimeInfo `json:"code_wait_event_time_details"`
}

func (o CodeWaitEventTime) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CodeWaitEventTime struct{}"
	}

	return strings.Join([]string{"CodeWaitEventTime", string(data)}, " ")
}
