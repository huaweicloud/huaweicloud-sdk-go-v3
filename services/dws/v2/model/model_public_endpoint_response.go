package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PublicEndpointResponse **参数解释**： 公网域名、公网IP信息对象。 **取值范围**： 不涉及。
type PublicEndpointResponse struct {

	// **参数解释**： 公网IP信息。 **取值范围**： 不涉及。
	Ip *string `json:"ip,omitempty"`

	// **参数解释**： 端口信息，创建集群时如果未指定端口则默认8000。 **取值范围**： 不涉及。
	Port *int32 `json:"port,omitempty"`

	// **参数解释**： 当前局点是否支持公网域名。 **取值范围**： 不涉及。
	Enabled *bool `json:"enabled,omitempty"`

	// **参数解释**： 公网IP的ID。 **取值范围**： 不涉及。
	IpId *string `json:"ip_id,omitempty"`

	// **参数解释**： 公网IP的带宽信息。 **取值范围**： 不涉及。
	IpBandwidth *string `json:"ip_bandwidth,omitempty"`

	// **参数解释**： 公网域名子域名信息。 **取值范围**： 不涉及。
	DomainName *string `json:"domain_name,omitempty"`

	// **参数解释**： 公网域名后缀信息。 **取值范围**： 不涉及。
	DomainNameSuffix *string `json:"domain_name_suffix,omitempty"`

	// **参数解释**： 公网域名后缀信息。 **取值范围**： 不涉及。
	ZoneName *string `json:"zone_name,omitempty"`

	// **参数解释**： 公网域名TTL。 **取值范围**： 不涉及。
	DomainNameTtl *int32 `json:"domain_name_ttl,omitempty"`

	// **参数解释**： 公网域名状态。 **取值范围**： 不涉及。
	DomainNameStatus *string `json:"domain_name_status,omitempty"`

	// **参数解释**： 公网IP状态。 **取值范围**： 不涉及。
	IpStatus *string `json:"ip_status,omitempty"`
}

func (o PublicEndpointResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PublicEndpointResponse struct{}"
	}

	return strings.Join([]string{"PublicEndpointResponse", string(data)}, " ")
}
