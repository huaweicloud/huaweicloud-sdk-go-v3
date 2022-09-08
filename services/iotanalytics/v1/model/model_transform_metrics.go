package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 转换计算定义
type TransformMetrics struct {
	TimeSpan *TimeSpan `json:"time_span"`

	// 对property按指定tags标签进行过滤查询，填入设备标签与标签值，不可为空，例如 {\"deviceId\": \"id0001\"}
	Tags map[string]string `json:"tags"`

	// 查询的测量指标列表
	Metrics []TransformMetric `json:"metrics"`

	// 返回值个数限制
	Limit *int32 `json:"limit,omitempty"`
}

func (o TransformMetrics) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TransformMetrics struct{}"
	}

	return strings.Join([]string{"TransformMetrics", string(data)}, " ")
}
