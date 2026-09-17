package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UserSeekTrend 用户最后查找趋势
type UserSeekTrend struct {

	// 近1天用户访问条数
	LastDayUserSeekCount *int64 `json:"last_day_user_seek_count,omitempty"`

	// 近1周用户访问条数
	LastWeekUserSeekCount *int64 `json:"last_week_user_seek_count,omitempty"`

	// 近2周用户访问条数
	LastTwoWeekUserSeekCount *int64 `json:"last_two_week_user_seek_count,omitempty"`

	// 近1月用户访问条数
	LastMonthUserSeekCount *int64 `json:"last_month_user_seek_count,omitempty"`
}

func (o UserSeekTrend) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UserSeekTrend struct{}"
	}

	return strings.Join([]string{"UserSeekTrend", string(data)}, " ")
}
