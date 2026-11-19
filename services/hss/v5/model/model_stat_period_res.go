package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StatPeriodRes 统计时间段
type StatPeriodRes struct {

	// **参数解释**： 开始时间 **取值范围**： 最小值0，最大值9223372036854775807，时间格式：毫秒级时间戳（UTC时区，从1970-01-01 00:00:00开始计算），单位：ms
	BeginTime *int64 `json:"begin_time,omitempty"`

	// **参数解释**： 结束时间 **取值范围**： 最小值0，最大值9223372036854775807，时间格式：毫秒级时间戳（UTC时区，从1970-01-01 00:00:00开始计算），单位：ms
	EndTime *int64 `json:"end_time,omitempty"`
}

func (o StatPeriodRes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StatPeriodRes struct{}"
	}

	return strings.Join([]string{"StatPeriodRes", string(data)}, " ")
}
