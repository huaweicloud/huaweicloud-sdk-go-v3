package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MetricDataDetailDimension 维度信息
type MetricDataDetailDimension struct {

	// 维度名称
	Name *string `json:"name,omitempty"`

	// 维度值
	Value *string `json:"value,omitempty"`
}

func (o MetricDataDetailDimension) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MetricDataDetailDimension struct{}"
	}

	return strings.Join([]string{"MetricDataDetailDimension", string(data)}, " ")
}
