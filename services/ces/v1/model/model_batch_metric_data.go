package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchMetricData
type BatchMetricData struct {

	// **参数解释** 指标单位 **取值范围** 不涉及
	Unit *string `json:"unit,omitempty"`

	// **参数解释** 指标数据列表。由于查询数据时，云监控会根据所选择的聚合粒度向前取整from参数，所以datapoints中包含的数据点有可能会多于预期，最多返回3000个数据点（响应参数metrics属性对应对象datapoints属性累加最多返回3000个数据点）
	Datapoints []DatapointForBatchMetric `json:"datapoints"`

	// **参数解释** 服务命名空间 **取值范围** 不涉及
	Namespace *string `json:"namespace,omitempty"`

	// **参数解释** 指标名称，例如弹性云服务器监控指标中的cpu_util。 **取值范围** 不涉及
	MetricName string `json:"metric_name"`

	// **参数解释** 服务维度列表
	Dimensions *[]MetricsDimensionResp `json:"dimensions,omitempty"`
}

func (o BatchMetricData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchMetricData struct{}"
	}

	return strings.Join([]string{"BatchMetricData", string(data)}, " ")
}
