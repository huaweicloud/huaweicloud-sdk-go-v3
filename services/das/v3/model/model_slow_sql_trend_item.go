package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SlowSqlTrendItem struct {

	// 毫秒时间戳。表示统计数据的时间范围为timestamp到timestamp + interval_millis。
	Timestamp int64 `json:"timestamp"`

	// 慢SQL数量。
	SlowLogCount int64 `json:"slow_log_count"`
}

func (o SlowSqlTrendItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SlowSqlTrendItem struct{}"
	}

	return strings.Join([]string{"SlowSqlTrendItem", string(data)}, " ")
}
