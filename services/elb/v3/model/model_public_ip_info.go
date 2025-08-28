package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PublicIpInfo struct {

	// **参数解释**：弹性公网EIP的ID。  **取值范围**：不涉及
	PublicipId string `json:"publicip_id"`

	// **参数解释**：EIP的IP地址。  **取值范围**：不涉及
	PublicipAddress string `json:"publicip_address"`

	// **参数解释**：IP版本信息。  **取值范围**： - 4：IPv4地址。 - 6：IPv6地址。  [不支持IPv6，请勿设置为6。](tag:dt)
	IpVersion int32 `json:"ip_version"`
}

func (o PublicIpInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PublicIpInfo struct{}"
	}

	return strings.Join([]string{"PublicIpInfo", string(data)}, " ")
}
