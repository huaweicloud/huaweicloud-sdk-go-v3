package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CountRiskTrendStatisticsResponse Response Object
type CountRiskTrendStatisticsResponse struct {

	// 生成时间
	GenerateTime *string `json:"generate_time,omitempty"`

	// 数据列表
	RiskStatistics *[]RiskStatisticsBean `json:"risk_statistics,omitempty"`

	// 状态  - FINISHED：已完成  - RUNNING：运行中
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CountRiskTrendStatisticsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CountRiskTrendStatisticsResponse struct{}"
	}

	return strings.Join([]string{"CountRiskTrendStatisticsResponse", string(data)}, " ")
}
