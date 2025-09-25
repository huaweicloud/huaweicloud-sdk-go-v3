package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type EventTimeInfo struct {

	// **参数解释**: TOP5事件耗时信息列表。
	Events []TopEventInfoResult `json:"events"`

	// **参数解释**: 其余事件耗时（单位：微秒）。 **取值范围**: 不涉及。
	LeftTime int64 `json:"left_time"`

	// **参数解释**: 事件外耗时（单位：微秒）。 **取值范围**: 不涉及。
	OtherTime int64 `json:"other_time"`
}

func (o EventTimeInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EventTimeInfo struct{}"
	}

	return strings.Join([]string{"EventTimeInfo", string(data)}, " ")
}
