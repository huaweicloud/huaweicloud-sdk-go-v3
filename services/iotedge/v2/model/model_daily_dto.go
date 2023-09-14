package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DailyDto 每天的时间描述结构体
type DailyDto struct {
	ExceptionalDates *ExceptionalDates `json:"exceptional_dates,omitempty"`

	// 时间段描述
	TimeSpans []TimeSpans `json:"time_spans"`

	// 描述此任务有效的星期几，为数组。如果为null，则表示start_time-endtime之间，每天都执行；如果为空数组，则表示start_time-endtime之间，每天都不执行；如果为[1, 2]，则表示每周一、二执行。
	Weekdays *interface{} `json:"weekdays,omitempty"`
}

func (o DailyDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DailyDto struct{}"
	}

	return strings.Join([]string{"DailyDto", string(data)}, " ")
}
