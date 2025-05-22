package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListClusterNodesResponse Response Object
type ListClusterNodesResponse struct {

	// **参数解释**： 集群节点列表。 **取值范围**： 不涉及。
	NodeList *[]ClusterNodeInfo `json:"node_list,omitempty"`

	// **参数解释**： 集群节点总数。 **取值范围**： 大于0的整数。
	Count *int32 `json:"count,omitempty"`

	// **参数解释**： 逻辑集群节点失败总数。一般为0。 **取值范围**： 大于等于0的整数。
	FailedCount    *int32 `json:"failed_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListClusterNodesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListClusterNodesResponse struct{}"
	}

	return strings.Join([]string{"ListClusterNodesResponse", string(data)}, " ")
}
