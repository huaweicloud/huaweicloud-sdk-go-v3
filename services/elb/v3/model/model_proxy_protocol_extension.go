package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ProxyProtocolExtension struct {

	// **参数解释**：ipv4 vip地址。  **约束限制**：不涉及  **取值范围**：不涉及  **默认取值**：不涉及
	VipAddress *string `json:"vip_address,omitempty"`

	// **参数解释**：ipv6 vip地址。  **约束限制**：不涉及  **取值范围**：不涉及  **默认取值**：不涉及
	Ipv6VipAddress *string `json:"ipv6_vip_address,omitempty"`

	Extension *Extension `json:"extension"`
}

func (o ProxyProtocolExtension) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProxyProtocolExtension struct{}"
	}

	return strings.Join([]string{"ProxyProtocolExtension", string(data)}, " ")
}
