package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListMetricsDataRequest Request Object
type ListMetricsDataRequest struct {

	// **参数解释**： 集群ID。获取方式方法请参见[获取集群ID](dws_02_00068.xml)。 **约束限制**： 必须是有效的dws集群ID。 **取值范围**： 36位UUID。 **默认取值**： 不涉及。
	ClusterId string `json:"cluster_id"`

	// 指标名称。
	MetricName string `json:"metric_name"`

	// **参数解释**： 分页查询，偏移量。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 0
	Offset int32 `json:"offset"`

	// **参数解释**： 分页单页大小。 **约束限制**： 不涉及。 **取值范围**： 大于0，最大1000。 **默认取值**： 不限制。
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
