package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SupplementDataRespSupplementDataRunTime 补数据时间段，当前仅支持每天，如果没有补数据时间，则默认为\"00:00-00:00”
type SupplementDataRespSupplementDataRunTime struct {

	// 每天的可补数据时间段，如：每天的10点15分到23点30分，表示：10:15-23:30
	TimeOfDay *string `json:"time_of_day,omitempty"`

	// 每周的星期几可以补数据，如：每周一，周三的每天10点15分到23点30分
	DayOfWeek *string `json:"day_of_week,omitempty"`

	// 每个月的哪几天可以补数据，如每月1号，3号，表示：1,3
	DayOfMonth *string `json:"day_of_month,omitempty"`
}

func (o SupplementDataRespSupplementDataRunTime) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SupplementDataRespSupplementDataRunTime struct{}"
	}

	return strings.Join([]string{"SupplementDataRespSupplementDataRunTime", string(data)}, " ")
}
