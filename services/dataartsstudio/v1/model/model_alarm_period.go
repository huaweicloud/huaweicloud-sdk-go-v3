package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AlarmPeriod struct {

	// 任务完成时间。
	CompleteTime *string `json:"complete_time,omitempty"`

	// 周期。
	Period *int32 `json:"period,omitempty"`
}

func (o AlarmPeriod) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlarmPeriod struct{}"
	}

	return strings.Join([]string{"AlarmPeriod", string(data)}, " ")
}
