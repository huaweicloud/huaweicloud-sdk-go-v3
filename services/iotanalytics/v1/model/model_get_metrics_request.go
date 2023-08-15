package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GetMetricsRequest 查询指标请求
type GetMetricsRequest struct {

	// 查询类型
	Type string `json:"type"`

	Transform *TransformMetrics `json:"transform,omitempty"`

	Aggregate *AggregateMetrics `json:"aggregate,omitempty"`
}

func (o GetMetricsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetMetricsRequest struct{}"
	}

	return strings.Join([]string{"GetMetricsRequest", string(data)}, " ")
}
