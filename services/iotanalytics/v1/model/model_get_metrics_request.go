package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 查询指标请求
type GetMetricsRequest struct {

	// 查询类型
	Type string `json:"type" xml:"type"`

	Transform *TransformMetrics `json:"transform,omitempty" xml:"transform"`

	Aggregate *AggregateMetrics `json:"aggregate,omitempty" xml:"aggregate"`
}

func (o GetMetricsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetMetricsRequest struct{}"
	}

	return strings.Join([]string{"GetMetricsRequest", string(data)}, " ")
}
