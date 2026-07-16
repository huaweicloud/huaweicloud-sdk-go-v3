package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PoolDriver 资源池驱动数据模型。
type PoolDriver struct {

	// **参数解释**：资源池默认的GPU驱动版本。物理资源池中包含GPU规格时有效。 **取值范围**：不涉及。
	GpuVersion *string `json:"gpuVersion,omitempty"`

	// **参数解释**：资源池默认的NPU驱动版本。物理资源池中包含NPU规格时有效。 **取值范围**：不涉及。
	NpuVersion *string `json:"npuVersion,omitempty"`

	// **参数解释**：资源池驱动升级策略。 **取值范围**：可选值如下： - force：强制升级，立即升级节点驱动，可能影响节点上正在运行的作业。 - idle：安全升级，待节点上没有作业运行时进行驱动升级。
	UpdateStrategy *string `json:"updateStrategy,omitempty"`
}

func (o PoolDriver) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PoolDriver struct{}"
	}

	return strings.Join([]string{"PoolDriver", string(data)}, " ")
}
