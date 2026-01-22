package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateIpsWhitelistDto **参数解释**：  更新ips白名单请求体 **约束限制**：  不涉及  **取值范围**： 不涉及  **默认取值**：  不涉及
type UpdateIpsWhitelistDto struct {

	// **参数解释**： 白名单源地址 **约束限制**： 不涉及  **取值范围**： 不涉及  **默认取值**： 不涉及
	SourceAddress string `json:"source_address"`

	// **参数解释**：  目的地址  **约束限制**：  不涉及  **取值范围**：  不涉及  **默认取值**：  不涉及
	DestAddress string `json:"dest_address"`

	// **参数解释**：  IP类型 **约束限制**：  不涉及  **取值范围**：  4：表示IP类型为IPv4 6：表示IP类型为IPv6 **默认取值**：  不涉及
	IpVersion int32 `json:"ip_version"`

	// **参数解释**：  白名单名称  **约束限制**：  不涉及  **取值范围**： 不涉及  **默认取值**：  不涉及
	Name string `json:"name"`

	// **参数解释**：  ips规则id，不填写时默认为全部ips规则  **约束限制**：  不涉及  **取值范围**： 不涉及  **默认取值**：  不涉及
	IpsId *string `json:"ips_id,omitempty"`

	// **参数解释**：  生效范围：NAT EIP VPC HCS场景中还有VGW信息  **约束限制**：  不涉及  **取值范围**： 不涉及  **默认取值**：  不涉及
	EffectiveScopes []string `json:"effective_scopes"`

	// **参数解释**：  白名单描述  **约束限制**：  不涉及  **取值范围**： 不涉及  **默认取值**：  不涉及
	Description *string `json:"description,omitempty"`
}

func (o UpdateIpsWhitelistDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateIpsWhitelistDto struct{}"
	}

	return strings.Join([]string{"UpdateIpsWhitelistDto", string(data)}, " ")
}
