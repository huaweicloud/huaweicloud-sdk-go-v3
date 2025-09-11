package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CountSqlTrendStatisticsResponse Response Object
type CountSqlTrendStatisticsResponse struct {

	// 列表数据
	CountStatistics *[]CountStatisticsBean `json:"count_statistics,omitempty"`

	// 生成时间
	GenerateTime *string `json:"generate_time,omitempty"`

	// 状态  - FINISHED：已完成  - RUNNING：运行中
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CountSqlTrendStatisticsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CountSqlTrendStatisticsResponse struct{}"
	}

	return strings.Join([]string{"CountSqlTrendStatisticsResponse", string(data)}, " ")
}
