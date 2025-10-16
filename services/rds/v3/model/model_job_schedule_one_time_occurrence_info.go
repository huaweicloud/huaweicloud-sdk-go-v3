package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// JobScheduleOneTimeOccurrenceInfo 执行一次执行时间。
type JobScheduleOneTimeOccurrenceInfo struct {

	// 执行日期 yyyy-MM-dd。取值范围1990-01-01至2099-12-31
	ActiveStartDate *string `json:"active_start_date,omitempty"`

	// 执行时间。HH:mm:ss
	ActiveStartTime *string `json:"active_start_time,omitempty"`
}

func (o JobScheduleOneTimeOccurrenceInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobScheduleOneTimeOccurrenceInfo struct{}"
	}

	return strings.Join([]string{"JobScheduleOneTimeOccurrenceInfo", string(data)}, " ")
}
