package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NodeResource 节点资源量。
type NodeResource struct {

	// **参数解释**：节点的CPU核心数量。 **取值范围**：不涉及。
	Cpu string `json:"cpu"`

	// **参数解释**：节点的内存大小。以Gi为单位。 **取值范围**：不涉及。
	Memory string `json:"memory"`

	// **参数解释**：节点的GPU卡数。 **取值范围**：不涉及。
	NvidiaComGpu *string `json:"nvidia.com/gpu,omitempty"`

	// **参数解释**：节点的snt3型NPU卡数量。 **取值范围**：不涉及。
	HuaweiComAscendSnt3 *string `json:"huawei.com/ascend-snt3,omitempty"`

	// **参数解释**：节点的snt9型NPU卡数量。 **取值范围**：不涉及。
	HuaweiComAscendSnt9 *string `json:"huawei.com/ascend-snt9,omitempty"`
}

func (o NodeResource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodeResource struct{}"
	}

	return strings.Join([]string{"NodeResource", string(data)}, " ")
}
