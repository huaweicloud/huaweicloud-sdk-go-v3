package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RiskTrendMetric 风险趋势指标数据
type RiskTrendMetric struct {

	// 数值
	Series *[]float64 `json:"series,omitempty"`

	// 时间戳
	Timestamps *[]int64 `json:"timestamps,omitempty"`
}

func (o RiskTrendMetric) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RiskTrendMetric struct{}"
	}

	return strings.Join([]string{"RiskTrendMetric", string(data)}, " ")
}
