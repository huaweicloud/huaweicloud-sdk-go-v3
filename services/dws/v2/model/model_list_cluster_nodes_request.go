package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListClusterNodesRequest Request Object
type ListClusterNodesRequest struct {

	// **参数解释**： 集群ID。获取方法请参见[获取集群ID](dws_02_00068.xml)。 **约束限制**： 必须是有效的dws集群ID。 **取值范围**： 36位UUID。 **默认取值**： 不涉及。
	ClusterId string `json:"cluster_id"`

	// **参数解释**： 节点ID列表。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： null
	NodeIds *[]string `json:"node_ids,omitempty"`

	// **参数解释**： 分页偏移量，从0开始，页数减1。 **约束限制**： 不涉及。 **取值范围**： 大于等于0。 **默认取值**： 0
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**： 分页查询，每页显示的条目数量。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 100
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**： 过滤字段。 **约束限制**： 不涉及。 **取值范围**： instCreateType：根据资源状态过滤 status：根据节点状态过滤 **默认取值**： null
	FilterBy *string `json:"filter_by,omitempty"`

	// **参数解释**： 过滤字段内容。 **约束限制**： 不涉及。 **取值范围**： 当根据资源状态过滤时，可选如下值： - ALL：全部 - INST：已使用 - NODE：空虚 当根据节点状态过滤时，可选如下值： - ALL：全部 - CREATING：创建中 - FREE：空闲 - ACTIVE：可用 - FAILED：不可用 - UNKNOWN：未知 - CREATE_FAILED：创建失败 - DELETING：删除中 - DELETE_FAILED：删除失败 **默认取值**： null
	Filter *string `json:"filter,omitempty"`

	// **参数解释**： 排序字段。默认无序返回。 **约束限制**： 不涉及。 **取值范围**： name：根据名称过滤 **默认取值**： null
	OrderBy *string `json:"order_by,omitempty"`

	// **参数解释**： 排序：升序/降序。 **约束限制**： 不涉及。 **取值范围**： asc：升序 desc：降序 **默认取值**： null
	Order *string `json:"order,omitempty"`

	// **参数解释**： 是否被删除，字段已废弃。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： null
	Deleted *string `json:"deleted,omitempty"`
}

func (o ListClusterNodesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListClusterNodesRequest struct{}"
	}

	return strings.Join([]string{"ListClusterNodesRequest", string(data)}, " ")
}
