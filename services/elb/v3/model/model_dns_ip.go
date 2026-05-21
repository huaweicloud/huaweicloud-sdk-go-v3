package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DnsIp 负载均衡器域名解析配置中的IP地址。
type DnsIp struct {

	// **参数解释**：IPv4或IPv6地址。  **约束限制**：必须是当前负载均衡器绑定的私网地址或者公网地址。  **取值范围**：不涉及  **默认取值**：不涉及
	IpAddress *string `json:"ip_address,omitempty"`
}

func (o DnsIp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DnsIp struct{}"
	}

	return strings.Join([]string{"DnsIp", string(data)}, " ")
}
