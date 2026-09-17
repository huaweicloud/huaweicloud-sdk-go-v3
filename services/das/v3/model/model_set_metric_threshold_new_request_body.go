package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetMetricThresholdNewRequestBody 设置指标阈值请求体
type SetMetricThresholdNewRequestBody struct {

	// 指标码
	MetricCode string `json:"metric_code"`

	// 数据库类型
	EngineType string `json:"engine_type"`

	// 新阈值
	NewThreshold float64 `json:"new_threshold"`
}

func (o SetMetricThresholdNewRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetMetricThresholdNewRequestBody struct{}"
	}

	return strings.Join([]string{"SetMetricThresholdNewRequestBody", string(data)}, " ")
}
