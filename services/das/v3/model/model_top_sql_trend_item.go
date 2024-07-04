package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TopSqlTrendItem struct {

	// 执行时间点，毫秒时间戳。表示统计数据的时间范围为execute_at到execute_at + interval_millis。
	ExecuteAt int64 `json:"execute_at"`

	// 执行耗时小于100ms。
	QueryTimeIn100ms int64 `json:"query_time_in_100ms"`

	// 执行耗时100ms-500ms。
	QueryTimeIn500ms int64 `json:"query_time_in_500ms"`

	// 执行耗时500ms-1000ms
	QueryTimeIn1s int64 `json:"query_time_in_1s"`

	// 执行耗时大于1000ms。
	QueryTimeOver1s int64 `json:"query_time_over_1s"`
}

func (o TopSqlTrendItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TopSqlTrendItem struct{}"
	}

	return strings.Join([]string{"TopSqlTrendItem", string(data)}, " ")
}
