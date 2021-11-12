package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type SessionStatsTotal struct {
	// 会话总数。

	SessionCount int64 `json:"session_count"`
	// 独立用户个数。

	UserCount int64 `json:"user_count"`
	// 平均会话轮数，保留小数点后三位。

	AvgRequestCount float64 `json:"avg_request_count"`
	// 平均会话时长，保留小数点后三位。

	AvgSessionTime float64 `json:"avg_session_time"`
}

func (o SessionStatsTotal) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SessionStatsTotal struct{}"
	}

	return strings.Join([]string{"SessionStatsTotal", string(data)}, " ")
}
