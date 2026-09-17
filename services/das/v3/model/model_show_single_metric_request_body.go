package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSingleMetricRequestBody Query Single Metric New请求体
type ShowSingleMetricRequestBody struct {

	// 指标名称
	MetricName string `json:"metric_name"`

	// 开始时间（Unix timestamp，毫秒）
	StartAt int64 `json:"start_at"`

	// 结束时间（Unix timestamp，毫秒）
	EndAt int64 `json:"end_at"`

	// 实例信息列表
	InstanceInfos []InstanceInfoDtoForMetric `json:"instance_infos"`
}

func (o ShowSingleMetricRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSingleMetricRequestBody struct{}"
	}

	return strings.Join([]string{"ShowSingleMetricRequestBody", string(data)}, " ")
}
