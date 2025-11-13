package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowZoneNameServerResponse Response Object
type ShowZoneNameServerResponse struct {

	// **参数解释：** 是否全部为华为云DNS服务器地址。 **取值范围：** 不涉及。
	AllHwDns *bool `json:"all_hw_dns,omitempty"`

	// **参数解释：** 是否包含华为云DNS服务器地址。  **取值范围：** 不涉及。
	IncludeHwDns *bool `json:"include_hw_dns,omitempty"`

	// **参数解释：** DNS服务器地址。 **取值范围：** 不涉及。
	DnsServers *[]string `json:"dns_servers,omitempty"`

	// **参数解释：** 期望的DNS服务器地址。 **取值范围：** 不涉及。
	ExpectedDnsServers *[]string `json:"expected_dns_servers,omitempty"`

	// **参数解释：** 公网域名。 **取值范围：** 不涉及。
	DomainName     *string `json:"domain_name,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowZoneNameServerResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowZoneNameServerResponse struct{}"
	}

	return strings.Join([]string{"ShowZoneNameServerResponse", string(data)}, " ")
}
