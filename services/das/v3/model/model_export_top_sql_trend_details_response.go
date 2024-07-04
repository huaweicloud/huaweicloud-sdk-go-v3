package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportTopSqlTrendDetailsResponse Response Object
type ExportTopSqlTrendDetailsResponse struct {

	// 返回列表两个时间点之间的时间间隔。总查询时长一小时之内间隔10s，一小时到六小时范围内间隔60s。单位为毫秒。
	IntervalMillis *int64 `json:"interval_millis,omitempty"`

	// SQL执行耗时区间数据。
	TopSqlTrendItems *[]TopSqlTrendItem `json:"top_sql_trend_items,omitempty"`

	// 耗时区间数据总数。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ExportTopSqlTrendDetailsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportTopSqlTrendDetailsResponse struct{}"
	}

	return strings.Join([]string{"ExportTopSqlTrendDetailsResponse", string(data)}, " ")
}
