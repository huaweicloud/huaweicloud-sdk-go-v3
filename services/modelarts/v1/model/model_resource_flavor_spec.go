package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResourceFlavorSpec 资源规格描述信息。
type ResourceFlavorSpec struct {

	// **参数解释**：资源规格类型。 **取值范围**：可选值如下： - Dedicate：物理资源规格。物理资源规格可以创建节点资源。 [- Logical：逻辑资源规格。](tag:hcso)
	Type *string `json:"type,omitempty"`

	// **参数解释**：资源规格实例的计算架构。 **取值范围**：可选值如下： - x86：x86架构。 - arm64：ARM架构。
	CpuArch *string `json:"cpuArch,omitempty"`

	// **参数解释**：资源规格实例的CPU核心数量。 **取值范围**：不涉及。
	Cpu *string `json:"cpu,omitempty"`

	// **参数解释**：资源规格实例的内存大小。以Gi为单位。 **取值范围**：不涉及。
	Memory *string `json:"memory,omitempty"`

	Gpu *ResourceFlavorXpu `json:"gpu,omitempty"`

	Npu *ResourceFlavorXpu `json:"npu,omitempty"`

	// **参数解释**：资源规格实例的存储资源信息。
	DataVolume *[]ResourceFlavorSpecDataVolume `json:"dataVolume,omitempty"`

	// **参数解释**：资源规格支持的计费模式。
	BillingModes *[]int32 `json:"billingModes,omitempty"`

	// **参数解释**：资源规格计费码。 **取值范围**：不涉及。
	BillingCode *string `json:"billingCode,omitempty"`

	// **参数解释**：资源规格支持的作业类型列表。
	JobFlavors *[]string `json:"jobFlavors,omitempty"`
}

func (o ResourceFlavorSpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceFlavorSpec struct{}"
	}

	return strings.Join([]string{"ResourceFlavorSpec", string(data)}, " ")
}
