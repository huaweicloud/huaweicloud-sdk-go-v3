package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListLogicalClusterTasksRequest Request Object
type ListLogicalClusterTasksRequest struct {

	// **参数解释**： 集群ID。获取方法请参见[获取集群ID](dws_02_00068.xml)。 **约束限制**： 必须是有效的dws集群ID。 **取值范围**： 36位UUID。 **默认取值**： 不涉及。
	ClusterId string `json:"cluster_id"`

	// **参数解释**： 分页偏移量，从0开始，页数减1。 **约束限制**： 不涉及。 **取值范围**： 大于等于0。 **默认取值**： 0
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**： 分页单页大小。 **约束限制**： 不涉及。 **取值范围**： 大于0。 **默认取值**： 10
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**： 集群名称。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	LogicalClusterName *string `json:"logical_cluster_name,omitempty"`

	// **参数解释**： 类型。 **取值范围**： - Switch：物理集群转逻辑集群。 - Create：创建。 - Expand：从弹性池中扩容。 - Restart：重启。 - Delete：删除。 - Shrink：缩容到弹性池中。 - Grow：外部扩容 - Start：开机 - Stop：停机 - ShrinkElasticGroup：从弹性池中缩容。 - elasticExpand：自动弹性扩容。
	Type *string `json:"type,omitempty"`

	// **参数解释**： 排序字段。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	OrderBy *string `json:"order_by,omitempty"`

	// **参数解释**： 排序：升序/降序。 **约束限制**： 不涉及。 **取值范围**： ASC：表示按升序排序。 DESC：表示按降序排序。 **默认取值**： 不涉及。
	Order *string `json:"order,omitempty"`
}

func (o ListLogicalClusterTasksRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLogicalClusterTasksRequest struct{}"
	}

	return strings.Join([]string{"ListLogicalClusterTasksRequest", string(data)}, " ")
}
