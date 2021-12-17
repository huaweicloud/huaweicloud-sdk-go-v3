package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 查询经过聚合计算的序列值定义
type AggregateMetricsRequest struct {
	TimeSpan *TimeSpanDt `json:"time_span,omitempty"`
	// 聚合时间间隔，正则：\"^[1-9][0-9]*[dhms]$\"，示例：\"1d|1h|10m|10s\"

	Interval string `json:"interval"`
	// 聚合时间偏移量，需要小于interval，正则： \"^[1-9][0-9]*[hms]$\"，示例： \"1h|10m|10s\"

	Offset *string `json:"offset,omitempty"`
	// 对property按指定tags标签进行过滤查询，填入资产标签属性的属性名与属性值，不可为空，例如 {\"tagPropertyA\": \"id0001\"}；注意，标签过滤只对打上标签时刻之后的数据生效，打标签之前的数据不能通过标签过滤

	Tags map[string]interface{} `json:"tags,omitempty"`
	// 属性过滤器，最多5个

	PropertyFilter *[]PropertyFilter `json:"property_filter,omitempty"`
	// 聚合查询指标列表，对资产属性进行聚合查询得到指标

	Metrics []DtAggregateMetrics `json:"metrics"`
	// 返回值个数限制，最多2000个

	Limit *int32 `json:"limit,omitempty"`
}

func (o AggregateMetricsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AggregateMetricsRequest struct{}"
	}

	return strings.Join([]string{"AggregateMetricsRequest", string(data)}, " ")
}
