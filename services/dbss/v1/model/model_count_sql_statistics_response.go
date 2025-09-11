package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CountSqlStatisticsResponse Response Object
type CountSqlStatisticsResponse struct {

	// 生成时间
	GenerateTime *string `json:"generate_time,omitempty"`

	// SQL统计数据
	SqlStatistics *[]SqlStatisticsBean `json:"sql_statistics,omitempty"`

	// 状态  - FINISHED：已完成  - RUNNING：运行中
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CountSqlStatisticsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CountSqlStatisticsResponse struct{}"
	}

	return strings.Join([]string{"CountSqlStatisticsResponse", string(data)}, " ")
}
