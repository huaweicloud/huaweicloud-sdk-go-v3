package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRiskTrendResponse Response Object
type ListRiskTrendResponse struct {

	// 指标码
	MetricCode *string `json:"metric_code,omitempty"`

	Metric         *QueryRiskTrendMetric `json:"metric,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o ListRiskTrendResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRiskTrendResponse struct{}"
	}

	return strings.Join([]string{"ListRiskTrendResponse", string(data)}, " ")
}
