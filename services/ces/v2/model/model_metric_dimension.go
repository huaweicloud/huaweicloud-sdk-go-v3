package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 指标维度
type MetricDimension struct {

	// 指标维度名称
	Name string `json:"name" xml:"name"`

	// 指标维度值
	Value *string `json:"value,omitempty" xml:"value"`
}

func (o MetricDimension) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MetricDimension struct{}"
	}

	return strings.Join([]string{"MetricDimension", string(data)}, " ")
}
