package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportSlowSqlStatisticsResponse Response Object
type ExportSlowSqlStatisticsResponse struct {

	// 慢SQL统计列表。
	StatisticsList *[]SlowSqlStatistics `json:"statistics_list,omitempty"`

	// 总数。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ExportSlowSqlStatisticsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportSlowSqlStatisticsResponse struct{}"
	}

	return strings.Join([]string{"ExportSlowSqlStatisticsResponse", string(data)}, " ")
}
