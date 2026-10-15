package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// QuotaResource 配额信息
type QuotaResource struct {

	// **参数解释**：  配额类型，用于标识不同维度的资源配额。  **约束限制**：  不涉及。  **取值范围**：  - instances：实例的数量配额  **默认取值**：  不涉及
	Type *string `json:"type,omitempty"`

	// **参数解释**：  配额允许的最小值，表示该资源配额的下限。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	Min *int32 `json:"min,omitempty"`

	// **参数解释**：  配额允许的最大值，表示该资源配额的上限。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	Max *int32 `json:"max,omitempty"`

	// **参数解释**：  用户当前配额的实际值，即系统为该用户分配的配额数量。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	Quota *int32 `json:"quota,omitempty"`

	// **参数解释**：  用户已使用的配额值，表示当前已占用的资源数量。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	Used *int32 `json:"used,omitempty"`
}

func (o QuotaResource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QuotaResource struct{}"
	}

	return strings.Join([]string{"QuotaResource", string(data)}, " ")
}
