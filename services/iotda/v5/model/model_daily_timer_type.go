package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DailyTimerType 条件中每日定时类型的信息，自定义结构。
type DailyTimerType struct {

	// **参数说明**：规则触发的时间，格式：HH:MM。
	Time string `json:"time"`

	// **参数说明**：星期列表，以逗号分隔。1代表周日，2代表周一，依次类推，默认为每天。
	DaysOfWeek *string `json:"days_of_week,omitempty"`
}

func (o DailyTimerType) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DailyTimerType struct{}"
	}

	return strings.Join([]string{"DailyTimerType", string(data)}, " ")
}
