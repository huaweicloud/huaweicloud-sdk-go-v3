package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDeadLockStatisticsResponse Response Object
type ShowDeadLockStatisticsResponse struct {

	// 近1天死锁/锁阻塞总数
	LastDayCount *int64 `json:"last_day_count,omitempty"`

	// 近1周死锁/锁阻塞总数
	LastWeekCount *int64 `json:"last_week_count,omitempty"`

	// 近2周死锁/锁阻塞总数
	LastTwoWeekCount *int64 `json:"last_two_week_count,omitempty"`

	// 近1月死锁/锁阻塞总数
	LastMonthCount *int64 `json:"last_month_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowDeadLockStatisticsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDeadLockStatisticsResponse struct{}"
	}

	return strings.Join([]string{"ShowDeadLockStatisticsResponse", string(data)}, " ")
}
