package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 添加一条或多条自定义指标监控数据，请求参数。
type MetricDataItem struct {
	Metric *MetricInfo `json:"metric"`
	// 数据的有效期，超出该有效期则自动删除该数据，单位秒，最大值604800。

	Ttl int32 `json:"ttl"`
	// 数据收集时间  UNIX时间戳，单位毫秒。  说明： 因为客户端到服务器端有延时，因此插入数据的时间戳应该在[当前时间-3天+20秒，当前时间+10分钟-20秒]区间内，保证到达服务器时不会因为传输时延造成数据不能插入数据库。

	CollectTime int64 `json:"collect_time"`
	// 指标数据的值。

	Value float64 `json:"value"`
	// 数据的单位。

	Unit *string `json:"unit,omitempty"`
	// 数据的类型，只能是\"int\"或\"float\"

	Type *string `json:"type,omitempty"`
}

func (o MetricDataItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MetricDataItem struct{}"
	}

	return strings.Join([]string{"MetricDataItem", string(data)}, " ")
}
