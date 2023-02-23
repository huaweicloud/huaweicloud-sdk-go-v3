package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ValiditySchedule struct {

	// end时间
	EndTime *string `json:"end_time,omitempty"`

	// 策略递归
	Recurrences *[]ValidityRecurrence `json:"recurrences,omitempty"`

	// 开始时间
	StartTime *string `json:"start_time,omitempty"`

	// 时间域
	TimeZone *string `json:"time_zone,omitempty"`
}

func (o ValiditySchedule) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ValiditySchedule struct{}"
	}

	return strings.Join([]string{"ValiditySchedule", string(data)}, " ")
}
