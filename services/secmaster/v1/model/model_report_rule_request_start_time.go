package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ReportRuleRequestStartTime 报告统计周期的开始时间
type ReportRuleRequestStartTime struct {

	// 报告统计周期的日期，前一天为-1，当天为0
	Day *int32 `json:"day,omitempty"`

	// 报告统计周期的日期，前一周为-1，本周为0
	Week *int32 `json:"week,omitempty"`

	// 报告统计周期的日期，上一月为-1，本月为0
	Month *int32 `json:"month,omitempty"`

	// 报告统计周期的开始，具体时间，时分秒格式
	Time *string `json:"time,omitempty"`
}

func (o ReportRuleRequestStartTime) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReportRuleRequestStartTime struct{}"
	}

	return strings.Join([]string{"ReportRuleRequestStartTime", string(data)}, " ")
}
