package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListMetricsDataRequest Request Object
type ListMetricsDataRequest struct {

	// 集群ID。
	ClusterId string `json:"cluster_id"`

	// 指标名称。
	MetricName string `json:"metric_name"`

	// 偏移量，表示从此偏移量开始查询，offset>=0。
	Offset int32 `json:"offset"`

	// 每页显示的条目数量，最大1000。
	Limit int32 `json:"limit"`

	// 采集开始时间，13位时间戳。
	From int64 `json:"from"`

	// 采集结束时间，13位时间戳。开始时间到结束时间最多不超过一天。
	To int64 `json:"to"`

	// 排序字段。固定取值。 ctime：采集时间。
	OrderBy *string `json:"order_by,omitempty"`

	// 正序还是倒叙。固定取值。 asc：正序。 desc：倒序。
	SortBy *string `json:"sort_by,omitempty"`
}

func (o ListMetricsDataRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMetricsDataRequest struct{}"
	}

	return strings.Join([]string{"ListMetricsDataRequest", string(data)}, " ")
}
