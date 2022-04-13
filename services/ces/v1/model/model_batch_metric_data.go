package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type BatchMetricData struct {
	// 指标单位。

	Unit *string `json:"unit,omitempty"`
	// 指标数据列表。由于查询数据时，云监控会根据所选择的聚合粒度向前取整from参数，所以datapoints中包含的数据点有可能会多于预期。

	Datapoints []DatapointForBatchMetric `json:"datapoints"`
	// 指标命名空间，格式为service.item；service和item必须是字符串，必须以字母开头，只能包含0-9/a-z/A-Z/_，总长度最短为3，最大为32；各服务的命名空间可查看：“[服务命名空间](https://support.huaweicloud.com/usermanual-ces/zh-cn_topic_0202622212.html)”。

	Namespace *string `json:"namespace,omitempty"`
	// 指标名称，例如弹性云服务器监控指标中的cpu_util；各服务的指标名称可查看：“[服务指标名称](https://support.huaweicloud.com/usermanual-ces/zh-cn_topic_0202622212.html)”。

	MetricName string `json:"metric_name"`
	// 指标维度列表。

	Dimensions *[]MetricsDimension `json:"dimensions,omitempty"`
}

func (o BatchMetricData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchMetricData struct{}"
	}

	return strings.Join([]string{"BatchMetricData", string(data)}, " ")
}
