package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DnsIpResponse 负载均衡器dns ip信息。
type DnsIpResponse struct {

	// dns ip信息的ID。
	Id *string `json:"id,omitempty"`

	// **参数解释**：ip是否加入了域名解析。  **取值范围**： true：已加入域名解析。 false：未加入域名解析。
	Enable *bool `json:"enable,omitempty"`

	// **参数解释**：ip地址。可以是ipv4地址也可以是ipv6地址。  **约束限制**：必须是负载均衡器的私网地址或者公网地址。
	IpAddress *string `json:"ip_address,omitempty"`

	// **参数解释**：地址类型。  **取值范围**： vip：私网ip。 eip：公网ip。
	Type *string `json:"type,omitempty"`

	// **参数解释**：ip对应的域名。  **约束限制**： - 如果ip为私网类型，则这里为负载均衡实例的私网域名 - 如果ip为公网类型，则这里为负载均衡实例的公网域名
	DomainName *string `json:"domain_name,omitempty"`

	// 创建时间。格式：yyyy-MM-dd'T'HH:mm:ss'Z'，UTC时区。
	CreatedAt *string `json:"created_at,omitempty"`

	// 更新时间。格式：yyyy-MM-dd'T'HH:mm:ss'Z'，UTC时区。
	UpdatedAt *string `json:"updated_at,omitempty"`
}

func (o DnsIpResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DnsIpResponse struct{}"
	}

	return strings.Join([]string{"DnsIpResponse", string(data)}, " ")
}
