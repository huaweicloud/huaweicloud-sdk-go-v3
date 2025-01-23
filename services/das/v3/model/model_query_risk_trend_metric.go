package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// QueryRiskTrendMetric 指标值
type QueryRiskTrendMetric struct {

	// 数值
	Series *[]float64 `json:"series,omitempty"`

	// 时间戳
	Timestamps *[]int64 `json:"timestamps,omitempty"`
}

func (o QueryRiskTrendMetric) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryRiskTrendMetric struct{}"
	}

	return strings.Join([]string{"QueryRiskTrendMetric", string(data)}, " ")
}
