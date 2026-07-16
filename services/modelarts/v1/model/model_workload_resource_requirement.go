package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WorkloadResourceRequirement **参数解释**：作业运行请求的资源量。
type WorkloadResourceRequirement struct {

	// **参数解释**：作业的运行使用的CPU资源量。 **取值范围**：不涉及。
	Cpu *string `json:"cpu,omitempty"`

	// **参数解释**：作业的运行使用的内存资源量。 **取值范围**：不涉及。
	Memory *string `json:"memory,omitempty"`

	// **参数解释**：作业的运行使用的GPU资源量。 **取值范围**：不涉及。
	NvidiaComGpu *string `json:"nvidia.com/gpu,omitempty"`

	// **参数解释**：作业的运行使用的snt3类型NPU资源量。 **取值范围**：不涉及
	HuaweiComAscend310 *string `json:"huawei.com/ascend-310,omitempty"`

	// **参数解释**：作业的运行使用的snt9类型NPU资源量。 **取值范围**：不涉及。
	HuaweiComAscend1980 *string `json:"huawei.com/ascend-1980,omitempty"`
}

func (o WorkloadResourceRequirement) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkloadResourceRequirement struct{}"
	}

	return strings.Join([]string{"WorkloadResourceRequirement", string(data)}, " ")
}
