package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MetricDimensionExpandParam 指标维度扩充参数
type MetricDimensionExpandParam struct {

	// 维度扩充标签
	Labels []string `json:"labels"`

	// 维度扩充方法，填写指标数据面内置方法， 参数index从1开始
	Functions []string `json:"functions"`
}

func (o MetricDimensionExpandParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MetricDimensionExpandParam struct{}"
	}

	return strings.Join([]string{"MetricDimensionExpandParam", string(data)}, " ")
}
