package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Extension struct {

	// **参数解释**：EP ID.  **约束限制**：不涉及  **取值范围**：不涉及  **默认取值**：不涉及
	EpId *string `json:"ep_id,omitempty"`

	// **参数解释**：EP Service ID.  **约束限制**：不涉及  **取值范围**：不涉及  **默认取值**：不涉及
	EpServiceId *string `json:"ep_service_id,omitempty"`
}

func (o Extension) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Extension struct{}"
	}

	return strings.Join([]string{"Extension", string(data)}, " ")
}
