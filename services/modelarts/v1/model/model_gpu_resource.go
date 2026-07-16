package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type GpuResource struct {

	// **参数解释：** GPU卡数。 **取值范围：** 不涉及。
	Gpu *int32 `json:"gpu,omitempty"`

	// **参数解释：** GPU内存。 **取值范围：** 不涉及。
	GpuMemory *string `json:"gpu_memory,omitempty"`

	// **参数解释：** GPU类型。 **取值范围：** 不涉及。
	Type *string `json:"type,omitempty"`
}

func (o GpuResource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GpuResource struct{}"
	}

	return strings.Join([]string{"GpuResource", string(data)}, " ")
}
