package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListQueriesRequestBody struct {

	// **参数解释**： 集群ID。获取方法请参见[获取集群ID](dws_02_00068.xml)。 **约束限制**： 必须是有效的dws集群ID。 **取值范围**： 36位UUID。 **默认取值**： 不涉及。
	ClusterId string `json:"cluster_id"`

	// **参数解释**： 项目ID。获取方法请参见[获取项目ID](dws_02_0011.xml)。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	ProjectId string `json:"project_id"`

	// **参数解释**： 偏移量，表示从此偏移量开始查询，一般为页数减1。 **取值范围**： 大于等于0。
	Offset string `json:"offset"`

	// **参数解释**： 每页显示的条目数量。 **取值范围**： 大于0。
	Limit string `json:"limit"`

	// **参数解释**： 查询条件数组。 **取值范围**： 不涉及。
	Conditions []ListQueriesCondition `json:"conditions"`

	// **参数解释**： 排序字段。 **取值范围**： 不涉及。
	OrderBy *string `json:"order_by,omitempty"`

	// **参数解释**： 固定值db_queries。 **取值范围**： 不涉及。
	Target string `json:"target"`
}

func (o ListQueriesRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListQueriesRequestBody struct{}"
	}

	return strings.Join([]string{"ListQueriesRequestBody", string(data)}, " ")
}
