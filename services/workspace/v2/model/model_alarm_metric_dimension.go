package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AlarmMetricDimension 指标维度，目前最大可添加4个维度。
type AlarmMetricDimension struct {

	// 资源维度。
	Name *string `json:"name,omitempty"`

	// 资源维度值。
	Value *string `json:"value,omitempty"`
}

func (o AlarmMetricDimension) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlarmMetricDimension struct{}"
	}

	return strings.Join([]string{"AlarmMetricDimension", string(data)}, " ")
}
