package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type GpUsInfo struct {

	// **参数解释**：GPU类型。 **取值范围**：不涉及。
	Type *string `json:"type,omitempty"`

	// **参数解释**：GPU卡数。 **取值范围**：不涉及。
	Gpu *float64 `json:"gpu,omitempty"`

	// **参数解释**：GPU内存。 **取值范围**：不涉及。
	GpuMemory *string `json:"gpu_memory,omitempty"`
}

func (o GpUsInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GpUsInfo struct{}"
	}

	return strings.Join([]string{"GpUsInfo", string(data)}, " ")
}
