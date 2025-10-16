package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// JobScheduleDurationInfo 策略有效期。
type JobScheduleDurationInfo struct {

	// 第一次执行日期 yyyy-MM-dd。取值范围1990-01-01至2099-12-31
	ActiveStartDate *string `json:"active_start_date,omitempty"`

	// 最后一次执行日期。默认不会结束 yyyy-MM-dd
	ActiveEndDate *string `json:"active_end_date,omitempty"`
}

func (o JobScheduleDurationInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobScheduleDurationInfo struct{}"
	}

	return strings.Join([]string{"JobScheduleDurationInfo", string(data)}, " ")
}
