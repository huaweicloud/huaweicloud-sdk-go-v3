package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AlertRuleSchedule 告警规则时间表
type AlertRuleSchedule struct {

	// 整型时长
	DelayInterval *int32 `json:"delay_interval,omitempty"`

	// 整型时长
	FrequencyInterval *int32 `json:"frequency_interval,omitempty"`

	FrequencyUnit *FrequencyUnit `json:"frequency_unit,omitempty"`

	// 整型时长
	OvertimeInterval *int32 `json:"overtime_interval,omitempty"`

	// 整型时长
	PeriodInterval *int32 `json:"period_interval,omitempty"`

	PeriodUnit *FrequencyUnit `json:"period_unit,omitempty"`
}

func (o AlertRuleSchedule) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlertRuleSchedule struct{}"
	}

	return strings.Join([]string{"AlertRuleSchedule", string(data)}, " ")
}
