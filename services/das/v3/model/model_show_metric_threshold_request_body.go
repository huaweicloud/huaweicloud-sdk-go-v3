package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowMetricThresholdRequestBody 获取指标阈值请求体
type ShowMetricThresholdRequestBody struct {

	// 数据库引擎类型
	EngineType string `json:"engine_type"`

	// 指标名称
	MetricNames []string `json:"metric_names"`
}

func (o ShowMetricThresholdRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMetricThresholdRequestBody struct{}"
	}

	return strings.Join([]string{"ShowMetricThresholdRequestBody", string(data)}, " ")
}
