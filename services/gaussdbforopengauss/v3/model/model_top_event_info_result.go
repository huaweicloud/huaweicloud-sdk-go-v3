package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TopEventInfoResult struct {

	// **参数解释**: 事件名。 **取值范围**: 不涉及。
	EventName string `json:"event_name"`

	// **参数解释**: 事件耗时（单位：微秒）。 **取值范围**: 不涉及。
	EventTime int64 `json:"event_time"`
}

func (o TopEventInfoResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TopEventInfoResult struct{}"
	}

	return strings.Join([]string{"TopEventInfoResult", string(data)}, " ")
}
