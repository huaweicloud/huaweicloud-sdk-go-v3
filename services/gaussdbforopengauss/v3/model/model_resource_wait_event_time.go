package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResourceWaitEventTime struct {

	// **参数解释**: 总耗时（单位：微秒）。 **取值范围**: 不涉及。
	AllTime int64 `json:"all_time"`

	ResourceWaitEventTimeDetails *ResourceWaitEvenTimeDetails `json:"resource_wait_event_time_details"`

	// **参数解释**: 资源类等待事件外耗时（单位：微秒）。 **取值范围**: 不涉及。
	OtherTime *int64 `json:"other_time,omitempty"`
}

func (o ResourceWaitEventTime) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceWaitEventTime struct{}"
	}

	return strings.Join([]string{"ResourceWaitEventTime", string(data)}, " ")
}
