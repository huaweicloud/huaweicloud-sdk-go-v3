package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PrivateEndpointResponse **参数解释**： 内网域名信息。 **取值范围**： 不涉及。
type PrivateEndpointResponse struct {

	// **参数解释**： 内网IP信息，多个IP逗号分割。 **取值范围**： 不涉及。
	Ip *string `json:"ip,omitempty"`

	// **参数解释**： 端口信息。 **取值范围**： 8000~30000
	Port *int32 `json:"port,omitempty"`

	// **参数解释**： 子域名前缀。 **取值范围**： 不涉及。
	DomainName *string `json:"domain_name,omitempty"`

	// **参数解释**： 子域名后缀。 **取值范围**： 不涉及。
	DomainNameSuffix *string `json:"domain_name_suffix,omitempty"`

	// **参数解释**： 子域名信息。 **取值范围**： 不涉及。
	ZoneName *string `json:"zone_name,omitempty"`

	// **参数解释**： 内网域名TTL。 **取值范围**： 不涉及。
	DomainNameTtl *int32 `json:"domain_name_ttl,omitempty"`

	// **参数解释**： 内网域名状态。 **取值范围**： 不涉及。
	DomainNameStatus *string `json:"domain_name_status,omitempty"`

	// **参数解释**： ELB的内网IP信息。 **取值范围**： 不涉及。
	ElbIp *string `json:"elb_ip,omitempty"`

	// **参数解释**： IP绑定状态。 **取值范围**： 不涉及。
	BindManageIpStatus *int32 `json:"bind_manage_ip_status,omitempty"`
}

func (o PrivateEndpointResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PrivateEndpointResponse struct{}"
	}

	return strings.Join([]string{"PrivateEndpointResponse", string(data)}, " ")
}
