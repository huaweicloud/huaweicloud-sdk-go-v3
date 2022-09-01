package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type GetMetricsValue struct {

	// 查询类型，经过转换计算的序列值（transform）、经过聚合计算的序列值（aggregate）
	Type string `json:"type" xml:"type"`

	Transform *TransformMetricsRequest `json:"transform,omitempty" xml:"transform"`

	Aggregate *AggregateMetricsRequest `json:"aggregate,omitempty" xml:"aggregate"`
}

func (o GetMetricsValue) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetMetricsValue struct{}"
	}

	return strings.Join([]string{"GetMetricsValue", string(data)}, " ")
}
