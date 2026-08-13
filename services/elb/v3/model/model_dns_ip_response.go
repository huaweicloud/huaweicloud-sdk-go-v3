package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DnsIpResponse 负载均衡器IP地址的域名解析配置。
type DnsIpResponse struct {

	// **参数解释**：IP地址是否已加入到域名解析。  **取值范围**： - true：已加入域名解析。 - false：未加入域名解析。
	Enable *bool `json:"enable,omitempty"`

	// **参数解释**：IPv4或IPv6地址。  **约束限制**：必须是当前负载均衡器绑定的私网地址或者公网地址。
	IpAddress *string `json:"ip_address,omitempty"`

	// **参数解释**：IP地址类型。  **取值范围**： - vip：私网IP。 - eip：公网IP。
	Type *string `json:"type,omitempty"`

	// **参数解释**：当前IP地址关联的负载均衡实例域名。  **约束限制**： - 如果IP为私网类型，则这里为负载均衡实例的私网域名。 - 如果IP为公网类型，则这里为负载均衡实例的公网域名。
	DomainName *string `json:"domain_name,omitempty"`

	// **参数解释**：创建时间。  **取值范围**：不涉及
	CreatedAt *string `json:"created_at,omitempty"`

	// **参数解释**：更新时间。  **取值范围**：不涉及
	UpdatedAt *string `json:"updated_at,omitempty"`
}

func (o DnsIpResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DnsIpResponse struct{}"
	}

	return strings.Join([]string{"DnsIpResponse", string(data)}, " ")
}
