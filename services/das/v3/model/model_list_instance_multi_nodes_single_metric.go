package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListInstanceMultiNodesSingleMetric struct {

	// 指标名称
	MetricName string `json:"metric_name"`

	// 开始时间
	StartAt int64 `json:"start_at"`

	// 结束时间
	EndAt int64 `json:"end_at"`

	// 实例信息列表
	InstanceInfos []ListInstanceMultiNodesSingleMetricInstanceInfos `json:"instance_infos"`
}

func (o ListInstanceMultiNodesSingleMetric) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceMultiNodesSingleMetric struct{}"
	}

	return strings.Join([]string{"ListInstanceMultiNodesSingleMetric", string(data)}, " ")
}
