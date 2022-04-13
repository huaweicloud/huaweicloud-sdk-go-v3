package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowMetricDataResponse struct {
	// 指标数据列表。由于查询数据时，云监控会根据所选择的聚合粒度向前取整from参数，所以datapoints中包含的数据点有可能会多于预期。

	Datapoints *[]Datapoint `json:"datapoints,omitempty"`
	// 指标名称，例如弹性云服务器监控指标中的cpu_util。

	MetricName     *string `json:"metric_name,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowMetricDataResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMetricDataResponse struct{}"
	}

	return strings.Join([]string{"ShowMetricDataResponse", string(data)}, " ")
}
