package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTrainingJobMetricsResponse Response Object
type ShowTrainingJobMetricsResponse struct {

	// 运行指标。
	Metrics        *[]MetricObject `json:"metrics,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ShowTrainingJobMetricsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTrainingJobMetricsResponse struct{}"
	}

	return strings.Join([]string{"ShowTrainingJobMetricsResponse", string(data)}, " ")
}
