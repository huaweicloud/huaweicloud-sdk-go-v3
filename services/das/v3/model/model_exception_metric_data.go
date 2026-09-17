package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExceptionMetricData ExceptionMetricData对象
type ExceptionMetricData struct {

	// 实例ID
	InstanceId *string `json:"instance_id,omitempty"`

	// 指标名
	Name *string `json:"name,omitempty"`

	// 指标值列表
	Series *[]float64 `json:"series,omitempty"`

	// 时间戳列表
	Timestamps *[]int64 `json:"timestamps,omitempty"`
}

func (o ExceptionMetricData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExceptionMetricData struct{}"
	}

	return strings.Join([]string{"ExceptionMetricData", string(data)}, " ")
}
