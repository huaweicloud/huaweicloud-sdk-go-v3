package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResourceRequirementVo 作业使用的资源量信息。
type ResourceRequirementVo struct {

	// **参数解释**：作业的运行使用的CPU资源量。 **取值范围**：不涉及。
	Cpu *string `json:"cpu,omitempty"`

	// **参数解释**：作业的运行使用的内存资源量。 **取值范围**：不涉及。
	Memory *string `json:"memory,omitempty"`

	// **参数解释**：作业的运行使用的GPU资源量。 **取值范围**：不涉及。
	NvidiaComGpu *string `json:"nvidia.com/gpu,omitempty"`

	// **参数解释**：作业的运行使用的snt3类型NPU资源量。 **取值范围**：不涉及。
	HuaweiComAscend310 *string `json:"huawei.com/ascend-310,omitempty"`

	// **参数解释**：作业的运行使用的snt9类型NPU资源量。 **取值范围**：不涉及。
	HuaweiComAscend1980 *string `json:"huawei.com/ascend-1980,omitempty"`
}

func (o ResourceRequirementVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceRequirementVo struct{}"
	}

	return strings.Join([]string{"ResourceRequirementVo", string(data)}, " ")
}
