package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MetricDimension 指标维度
type MetricDimension struct {

	// 指标维度名称
	Name string `json:"name"`

	// 指标维度值
	Value *string `json:"value,omitempty"`
}

func (o MetricDimension) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MetricDimension struct{}"
	}

	return strings.Join([]string{"MetricDimension", string(data)}, " ")
}
