package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SlowLogTrendPoint 慢日志趋势点
type SlowLogTrendPoint struct {

	// 毫秒时间戳
	Timestamp *int64 `json:"timestamp,omitempty"`

	// 慢SQL数量
	SlowLogCount *int64 `json:"slow_log_count,omitempty"`
}

func (o SlowLogTrendPoint) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SlowLogTrendPoint struct{}"
	}

	return strings.Join([]string{"SlowLogTrendPoint", string(data)}, " ")
}
