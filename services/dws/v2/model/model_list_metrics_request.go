package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListMetricsRequest Request Object
type ListMetricsRequest struct {

	// **参数解释**： 集群ID。获取方式方法请参见[获取集群ID](dws_02_00068.xml)。 **约束限制**： 必须是有效的dws集群ID。 **取值范围**： 36位UUID。 **默认取值**： 不涉及。
	ClusterId string `json:"cluster_id"`

	// **参数解释**： 分页查询，偏移量。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 0
	Offset int32 `json:"offset"`

	// **参数解释**： 分页单页大小。 **约束限制**： 不涉及。 **取值范围**： 大于0，最大1000。 **默认取值**： 不限制。
	Limit int32 `json:"limit"`

	// 排序字段。固定取值。 create_time：创建时间。
	OrderBy *string `json:"order_by,omitempty"`

	// 正序还是倒叙。固定取值。 asc：正序。 desc：倒序。
	SortBy *string `json:"sort_by,omitempty"`
}

func (o ListMetricsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMetricsRequest struct{}"
	}

	return strings.Join([]string{"ListMetricsRequest", string(data)}, " ")
}
