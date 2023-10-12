package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateFactorySupplementDataInstanceRequestBodySupplementDataRunTime 运行时间段
type CreateFactorySupplementDataInstanceRequestBodySupplementDataRunTime struct {

	// 每天的可补数据时间段，如：每天的10点15分到23点30分，表示：10:15-23:30
	TimeOfDay *string `json:"time_of_day,omitempty"`

	// 每周的星期几可以补数据，如：每周一，周三的每天10点15分到23点30分。
	DayOfWeek *string `json:"day_of_week,omitempty"`

	// 每个月的哪几天可以补数据，如每月1号，3号，表示：1,3
	DayOfMonth *string `json:"day_of_month,omitempty"`
}

func (o CreateFactorySupplementDataInstanceRequestBodySupplementDataRunTime) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateFactorySupplementDataInstanceRequestBodySupplementDataRunTime struct{}"
	}

	return strings.Join([]string{"CreateFactorySupplementDataInstanceRequestBodySupplementDataRunTime", string(data)}, " ")
}
