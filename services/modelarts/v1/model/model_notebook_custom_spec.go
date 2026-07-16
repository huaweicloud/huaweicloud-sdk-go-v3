package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NotebookCustomSpec **参数描述**：自定义规格参数信息。 **约束限制**：CPU、GPU专属池下，支持的自定义规格参数，此字段与请求体中的flavor字段不可同时填写。
type NotebookCustomSpec struct {

	// **参数描述**：实例申请的GPU卡数。 **约束限制**：CPU专属池场景下无此字段，GPU专属池场景必填。 **取值范围**：不涉及。 **默认取值**：不涉及。
	Gpu *float32 `json:"gpu,omitempty"`

	// **参数描述**：实例申请的CPU核数大小。 **约束限制**：CPU/GPU专属池场景下自定义规格必填。 **取值范围**：整数部分最多10位，小数部分最多2位，且数值不得小于0.4。 **默认取值**：不涉及。
	Cpu float32 `json:"cpu"`

	// **参数描述**：实例申请的内存大小。 **约束限制**：CPU/GPU专属池场景下自定义规格必填。 **取值范围**：必须是整数，整数部分最多10位，且数值不得小于513。 **默认取值**：不涉及。
	Memory float32 `json:"memory"`

	// **参数描述**：实例申请的GPU加速卡类型。 **约束限制**：CPU专属池场景下无此字段，GPU专属池场景下必填。 **取值范围**：不涉及。 **默认取值**：不涉及。
	GpuType *string `json:"gpu_type,omitempty"`

	// **参数描述**：实例申请的CPU架构。 **约束限制**：CPU/GPU专属池场景下自定义规格必填。 **取值范围**：枚举类型，取值如下： - X86_64：x86架构 - AARCH64：ARM架构  **默认取值**：不涉及。
	Arch *string `json:"arch,omitempty"`

	// **参数描述**：实例申请的规格类型。 **约束限制**：CPU/GPU专属池场景下自定义规格必填。 **取值范围**：枚举类型，取值如下： - CPU：CPU规格。 - GPU：GPU规格。  **默认取值**：不涉及。
	Category *string `json:"category,omitempty"`

	// **参数解释**：实例选择的目标资源池节点实例规格。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	ResourceFlavor *string `json:"resource_flavor,omitempty"`
}

func (o NotebookCustomSpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NotebookCustomSpec struct{}"
	}

	return strings.Join([]string{"NotebookCustomSpec", string(data)}, " ")
}
