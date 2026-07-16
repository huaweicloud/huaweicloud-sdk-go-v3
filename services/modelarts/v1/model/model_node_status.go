package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NodeStatus 节点状态数据模型。
type NodeStatus struct {

	// **参数解释**：节点当前状态。 **取值范围**：可选值如下： - Available：节点可用。 - Creating：节点创建中。 - Deleting：节点删除中。 - Abnormal：节点异常。 - Checking: 节点自检中。
	Phase string `json:"phase"`

	// **参数解释**：节点所在的az。 **取值范围**：不涉及。
	Az *string `json:"az,omitempty"`

	// **参数解释**：节点的IP地址。 **取值范围**：不涉及。
	PrivateIp *string `json:"privateIp,omitempty"`

	Resources *NodeResource `json:"resources"`

	AvailableResources *NodeResource `json:"availableResources"`
}

func (o NodeStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodeStatus struct{}"
	}

	return strings.Join([]string{"NodeStatus", string(data)}, " ")
}
