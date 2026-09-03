package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTrainingFlavorMaxAvailableResourceResponse Response Object
type ShowTrainingFlavorMaxAvailableResourceResponse struct {

	// **参数解释**：最大可用CPU核数。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	CpuCoreNum *int32 `json:"cpu_core_num,omitempty"`

	// **参数解释**：最大可用内存大小，单位为GB。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	MemSize        *int32 `json:"mem_size,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowTrainingFlavorMaxAvailableResourceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTrainingFlavorMaxAvailableResourceResponse struct{}"
	}

	return strings.Join([]string{"ShowTrainingFlavorMaxAvailableResourceResponse", string(data)}, " ")
}
