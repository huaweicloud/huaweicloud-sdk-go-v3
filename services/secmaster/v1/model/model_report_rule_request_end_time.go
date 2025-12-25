package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ReportRuleRequestEndTime 报告统计周期的结束时间
type ReportRuleRequestEndTime struct {

	// 报告统计周期的日期，前一天为-1，当天为0
	Day *int32 `json:"day,omitempty"`

	// 报告统计周期的日期，前一周为-1，本周为0
	Week *int32 `json:"week,omitempty"`

	// 报告统计周期的日期，上一月为-1，本月为0
	Month *int32 `json:"month,omitempty"`

	// 报告统计周期的结束，具体时间，时分秒格式
	Time *string `json:"time,omitempty"`
}

func (o ReportRuleRequestEndTime) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReportRuleRequestEndTime struct{}"
	}

	return strings.Join([]string{"ReportRuleRequestEndTime", string(data)}, " ")
}
