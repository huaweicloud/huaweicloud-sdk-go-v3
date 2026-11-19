package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type StatPeriod struct {

	// **参数解释**： 结束时间，毫秒 **取值范围**： 最小值0，最大值9223372036854775807
	BeginTime *int64 `json:"begin_time,omitempty"`

	// **参数解释**： 结束时间，毫秒 **取值范围**： 最小值0，最大值9223372036854775807
	EndTime *int64 `json:"end_time,omitempty"`
}

func (o StatPeriod) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StatPeriod struct{}"
	}

	return strings.Join([]string{"StatPeriod", string(data)}, " ")
}
