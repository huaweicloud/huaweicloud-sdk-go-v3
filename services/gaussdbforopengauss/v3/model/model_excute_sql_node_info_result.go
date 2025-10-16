package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ExcuteSqlNodeInfoResult struct {

	// **参数解释**: 节点ID。 **取值范围**: 不涉及。
	NodeId *string `json:"node_id,omitempty"`

	// **参数解释**: 节点名称。 **取值范围**: 不涉及。
	NodeName *string `json:"node_name,omitempty"`

	// **参数解释**: 节点角色。 **取值范围**: - master：主节点。 - slave：备节点。 - secondary：日志节点。 - readreplica：只读节点。
	Role *string `json:"role,omitempty"`

	// **参数解释**: 实例ID。 **取值范围**: 不涉及。
	InstanceId *string `json:"instance_id,omitempty"`

	// **参数解释**: 组件类型。 **取值范围**: - CN：CN组件。 - DN：DN组件。
	ComponentType *string `json:"component_type,omitempty"`
}

func (o ExcuteSqlNodeInfoResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExcuteSqlNodeInfoResult struct{}"
	}

	return strings.Join([]string{"ExcuteSqlNodeInfoResult", string(data)}, " ")
}
