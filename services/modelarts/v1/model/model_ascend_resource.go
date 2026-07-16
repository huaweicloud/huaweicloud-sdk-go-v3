package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AscendResource struct {

	// **参数解释：** NPU数量。 **取值范围：** 不涉及。
	Npu *int32 `json:"npu,omitempty"`

	// **参数解释：** NPU内存。 **取值范围：** 不涉及。
	NpuMemory *string `json:"npu_memory,omitempty"`

	// **参数解释：** NPU类型。 **取值范围：** 不涉及。
	Type *string `json:"type,omitempty"`

	// **参数解释：** 切分规格中的ai_core。 **取值范围：** 不涉及。
	AiCore *string `json:"ai_core,omitempty"`

	// **参数解释：** 切分规格中的ai_cpu。 **取值范围：** 不涉及。
	AiCpu *string `json:"ai_cpu,omitempty"`
}

func (o AscendResource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AscendResource struct{}"
	}

	return strings.Join([]string{"AscendResource", string(data)}, " ")
}
