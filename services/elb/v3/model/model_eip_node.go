package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EipNode
type EipNode struct {

	// **参数解释**：不涉及EIP id。  **取值范围**：不涉及
	Id string `json:"id"`

	// **参数解释**：EIP地址。  **取值范围**：不涉及
	IpAddress string `json:"ip_address"`

	// **参数解释**：EIP 地址类型。  **取值范围**：不涉及 - 4：ipv4。 - 6：ipv6
	IpVersion int32 `json:"ip_version"`
}

func (o EipNode) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EipNode struct{}"
	}

	return strings.Join([]string{"EipNode", string(data)}, " ")
}
