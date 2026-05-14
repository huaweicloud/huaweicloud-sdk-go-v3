package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MetricsDimension
type MetricsDimension struct {

	// 维度名称。
	Name *string `json:"name,omitempty"`

	// 维度值。
	Value *string `json:"value,omitempty"`
}

func (o MetricsDimension) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MetricsDimension struct{}"
	}

	return strings.Join([]string{"MetricsDimension", string(data)}, " ")
}
