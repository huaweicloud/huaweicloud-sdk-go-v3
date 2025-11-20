package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type HyperNodeStatus struct {

	// **参数解释** 超节点状态 **取值范围** - provisioning: 创建中。 - active: 整体可用，代表超节点下所有节点都可用。 - partially-available: 超节点下存在不可用节点时会从 active 转成此状态。 - error: 错误状态。 - deleting: 删除中。 - reinstalling: 重置中。 - scaling: 扩容或缩容中。
	Phase *string `json:"phase,omitempty"`

	// **参数解释** 超节点实例 ID
	InstanceID *string `json:"instanceID,omitempty"`

	// **参数解释** 超节点下节点总数
	CurrentNode *int32 `json:"currentNode,omitempty"`

	// **参数解释** 超节点下处于删除中的节点数
	DeletingNode *int32 `json:"deletingNode,omitempty"`

	// **参数解释** 超节点下处于创建中的节点数
	CreatingNode *int32 `json:"creatingNode,omitempty"`

	// **参数解释** 超节点下处于可用状态的节点数
	ActiveNode *int32 `json:"activeNode,omitempty"`
}

func (o HyperNodeStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HyperNodeStatus struct{}"
	}

	return strings.Join([]string{"HyperNodeStatus", string(data)}, " ")
}
