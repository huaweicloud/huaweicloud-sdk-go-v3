package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResourceLimit **参数解释**：资源限制。
type ResourceLimit struct {

	// **参数解释**：资源类型。 **取值范围**：   - CPU：中央处理器。   - GPU：图形处理器。   - NPU：神经网络处理器。
	Type string `json:"type"`

	// **参数解释**：资源规格编码，标识资源池中具体的资源规格，由系统生成，例如\"aidatalake.cpu.x86.general.1c2g\"。 **取值范围**：由英文字母、数字、小数点和中划线组成，长度为1~64个字符。
	SpecCode string `json:"spec_code"`

	ComputeAllocatable *ComputeAllocatable `json:"compute_allocatable,omitempty"`

	ComputeSize *ComputeSize `json:"compute_size,omitempty"`

	// **参数解释**：可选的资源配置枚举列表。
	WorkerSizes *[]WorkerSizeInfo `json:"worker_sizes,omitempty"`
}

func (o ResourceLimit) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceLimit struct{}"
	}

	return strings.Join([]string{"ResourceLimit", string(data)}, " ")
}
