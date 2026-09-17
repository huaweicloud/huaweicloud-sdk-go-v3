package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRiskTrendResponse Response Object
type ShowRiskTrendResponse struct {

	// 指标码
	MetricCode *string `json:"metric_code,omitempty"`

	Metric         *RiskTrendMetric `json:"metric,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ShowRiskTrendResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRiskTrendResponse struct{}"
	}

	return strings.Join([]string{"ShowRiskTrendResponse", string(data)}, " ")
}
