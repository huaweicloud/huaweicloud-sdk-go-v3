package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowInstanceMetricRequestBody 查询实例指标请求体
type ShowInstanceMetricRequestBody struct {

	// 数据库引擎类型
	EngineType string `json:"engine_type"`

	// 实例信息列表
	Infos []InstanceInfoForMetric `json:"infos"`

	// 指标名称
	MetricNames []string `json:"metric_names"`
}

func (o ShowInstanceMetricRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceMetricRequestBody struct{}"
	}

	return strings.Join([]string{"ShowInstanceMetricRequestBody", string(data)}, " ")
}
