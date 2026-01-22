package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type IpsWhiteListVo struct {

	// **参数解释**： 白名单源地址 **取值范围**： 不涉及
	SourceAddress *string `json:"source_address,omitempty"`

	// **参数解释**：  目的地址  **取值范围**：  不涉及
	DestAddress *string `json:"dest_address,omitempty"`

	// **参数解释**：  白名单名称  **取值范围**： 不涉及
	Name *string `json:"name,omitempty"`

	// **参数解释**：  IPS白名单ID  **取值范围**： 不涉及
	ListId *string `json:"list_id,omitempty"`

	// **参数解释**：  生效范围：NAT EIP VPC HCS场景中还有VGW信息  **取值范围**： 不涉及
	EffectiveScopes *[]string `json:"effective_scopes,omitempty"`

	// **参数解释**：  生效范围：IPS规则ID，all代表所有IPS **取值范围**： 不涉及
	IpsId *string `json:"ips_id,omitempty"`

	// **参数解释**：  IP类型 **取值范围**：  4：表示IP类型为IPv4 6：表示IP类型为IPv6
	IpVersion *int32 `json:"ip_version,omitempty"`

	// **参数解释**：  白名单描述 **取值范围**： 不涉及
	Description *string `json:"description,omitempty"`
}

func (o IpsWhiteListVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IpsWhiteListVo struct{}"
	}

	return strings.Join([]string{"IpsWhiteListVo", string(data)}, " ")
}
