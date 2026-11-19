package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StatPeriodReq 统计时间段
type StatPeriodReq struct {

	// **参数解释**： 开始时间 **约束限制**： 不涉及 **取值范围**： 最小值0，最大值9223372036854775807，时间格式：毫秒级时间戳（UTC时区，从1970-01-01 00:00:00开始计算），单位：ms **默认取值**： 不涉及
	BeginTime *int64 `json:"begin_time,omitempty"`

	// **参数解释**： 结束时间 **约束限制**： 不涉及 **取值范围**： 最小值0，最大值9223372036854775807，时间格式：毫秒级时间戳（UTC时区，从1970-01-01 00:00:00开始计算），单位：ms **默认取值**： 不涉及
	EndTime *int64 `json:"end_time,omitempty"`
}

func (o StatPeriodReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StatPeriodReq struct{}"
	}

	return strings.Join([]string{"StatPeriodReq", string(data)}, " ")
}
