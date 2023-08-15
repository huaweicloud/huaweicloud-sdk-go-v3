package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SessionStatsIntervals
type SessionStatsIntervals struct {

	// 间隔周期开始时间。
	Start string `json:"start"`

	// 间隔周期会话总数。
	SessionCount int64 `json:"session_count"`

	// 间隔周期独立用户个数。
	UserCount int64 `json:"user_count"`

	// 间隔周期平均会话轮数，保留小数点后三位。
	AvgRequestCount float64 `json:"avg_request_count"`

	// 间隔周期平均会话时长，保留小数点后三位。
	AvgSessionTime float64 `json:"avg_session_time"`
}

func (o SessionStatsIntervals) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SessionStatsIntervals struct{}"
	}

	return strings.Join([]string{"SessionStatsIntervals", string(data)}, " ")
}
