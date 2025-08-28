package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PortRange struct {

	// **参数解释**：起始端口。  **约束限制**：不涉及  **取值范围**：1-65535  **默认取值：不涉及
	StartPort *int32 `json:"start_port,omitempty"`

	// **参数解释**：结束端口，需大于等于起始端口  **约束限制**：不涉及  **取值范围**：1-65535  **默认取值：不涉及
	EndPort *int32 `json:"end_port,omitempty"`
}

func (o PortRange) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PortRange struct{}"
	}

	return strings.Join([]string{"PortRange", string(data)}, " ")
}
