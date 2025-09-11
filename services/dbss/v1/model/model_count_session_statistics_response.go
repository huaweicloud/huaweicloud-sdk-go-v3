package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CountSessionStatisticsResponse Response Object
type CountSessionStatisticsResponse struct {

	// 生成时间
	GenerateTime *string `json:"generate_time,omitempty"`

	// 统计数据列表
	SessionStatistics *[]SessionStatisticsBean `json:"session_statistics,omitempty"`

	// 状态  - FINISHED：已完成  - RUNNING：运行中
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CountSessionStatisticsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CountSessionStatisticsResponse struct{}"
	}

	return strings.Join([]string{"CountSessionStatisticsResponse", string(data)}, " ")
}
