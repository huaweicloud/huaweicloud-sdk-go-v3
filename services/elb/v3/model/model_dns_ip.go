package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DnsIp 负载均衡器dns ip地址。
type DnsIp struct {

	// **参数解释**：ip地址。可以是ipv4地址也可以是ipv6地址。  **约束限制**：必须是负载均衡器的私网地址或者公网地址。  **取值范围**：不涉及  **默认取值**：不涉及
	IpAddress *string `json:"ip_address,omitempty"`
}

func (o DnsIp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DnsIp struct{}"
	}

	return strings.Join([]string{"DnsIp", string(data)}, " ")
}
