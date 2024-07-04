package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportSlowSqlTrendDetailsResponse Response Object
type ExportSlowSqlTrendDetailsResponse struct {

	// 慢SQL数量趋势。
	SlowSqlTrendItems *[]SlowSqlTrendItem `json:"slow_sql_trend_items,omitempty"`

	// 返回列表两个时间点之间的时间间隔。总查询时长3小时之内间隔1分钟，3小时到6小时范围内间隔5分钟，6小时到12小时范围内间隔30分钟，12小时以上间隔1小时。单位为毫秒。
	IntervalMillis *int64 `json:"interval_millis,omitempty"`

	// 趋势数据总数。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ExportSlowSqlTrendDetailsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportSlowSqlTrendDetailsResponse struct{}"
	}

	return strings.Join([]string{"ExportSlowSqlTrendDetailsResponse", string(data)}, " ")
}
