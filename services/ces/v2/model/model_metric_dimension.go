package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 指标维度
type MetricDimension struct {
	Name string `json:"name"`

	Value *string `json:"value,omitempty"`
}

func (o MetricDimension) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MetricDimension struct{}"
	}

	return strings.Join([]string{"MetricDimension", string(data)}, " ")
}
