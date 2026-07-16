package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MainContainerAllocatedResources 训练作业的训练容器实际到手的资源规格。
type MainContainerAllocatedResources struct {

	// **参数解释**： cpu架构。 **取值范围**： 不涉及。
	CpuArch *string `json:"cpu_arch,omitempty"`

	// **参数解释**： 核数。 **取值范围**： 不涉及。
	CpuCoreNum *float32 `json:"cpu_core_num,omitempty"`

	// **参数解释**： 内存信息。 **取值范围**： 不涉及。
	MemSize *float32 `json:"mem_size,omitempty"`

	// **参数解释**： 加速卡卡数。 **取值范围**： 不涉及。
	AcceleratorNum *float32 `json:"accelerator_num,omitempty"`

	// **参数解释**： 加速卡类型。如：ascend-Snt9b，ascend-snt9c等 **取值范围**： 不涉及。
	AcceleratorType *string `json:"accelerator_type,omitempty"`
}

func (o MainContainerAllocatedResources) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MainContainerAllocatedResources struct{}"
	}

	return strings.Join([]string{"MainContainerAllocatedResources", string(data)}, " ")
}
