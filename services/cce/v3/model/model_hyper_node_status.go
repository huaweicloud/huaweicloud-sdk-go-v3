package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type HyperNodeStatus struct {

	// **参数解释**： 超节点状态 **约束限制**： 不涉及 **取值范围**： - provisioning：创建中。 - active：整体可用，代表超节点下所有节点都可用。 - partially-available：超节点下存在不可用节点时会从 active 转成此状态。 - error：错误状态。 - deleting：删除中。 - reinstalling：重置中。 - scaling：扩容或缩容中。  **默认取值**： 不涉及
	Phase *string `json:"phase,omitempty"`

	// **参数解释**： 超节点ID **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	InstanceID *string `json:"instanceID,omitempty"`

	// **参数解释**： 超节点下节点总数 **约束限制**： 不涉及 **取值范围**： 大于等于0的整数 **默认取值**： 不涉及
	CurrentNode *int32 `json:"currentNode,omitempty"`

	// **参数解释**： 超节点下处于删除中的节点数 **约束限制**： 不涉及 **取值范围**： 大于等于0的整数 **默认取值**： 不涉及
	DeletingNode *int32 `json:"deletingNode,omitempty"`

	// **参数解释**： 超节点下处于创建中的节点数 **约束限制**： 不涉及 **取值范围**： 大于等于0的整数 **默认取值**： 不涉及
	CreatingNode *int32 `json:"creatingNode,omitempty"`

	// **参数解释**： 超节点下处于可用状态的节点数 **约束限制**： 不涉及 **取值范围**： 大于等于0的整数 **默认取值**： 不涉及
	ActiveNode *int32 `json:"activeNode,omitempty"`

	// **参数解释**： 超节点是否为纳管节点。纳管节点指用户已有的存量服务器接入CCE集群，而非由CCE自动创建的ECS/BMS。 **约束限制**： 不涉及 **取值范围**： - true：纳管节点，服务器在加入集群前已存在，删除超节点时不会释放底层云服务器资源。 - false：CCE创建的节点，生命周期由CCE管理，删除时会释放底层资源。 **默认取值**： false
	IsStatic *bool `json:"isStatic,omitempty"`
}

func (o HyperNodeStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HyperNodeStatus struct{}"
	}

	return strings.Join([]string{"HyperNodeStatus", string(data)}, " ")
}
