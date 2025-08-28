package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EipInfo **参数解释**：弹性IP，同publicips。
type EipInfo struct {

	// **参数解释**：弹性IP的ID。  **取值范围**：不涉及
	EipId *string `json:"eip_id,omitempty"`

	// **参数解释**：弹性IP的IP地址。  **取值范围**：不涉及
	EipAddress *string `json:"eip_address,omitempty"`

	// **参数解释**：IP版本号。  **取值范围**： - 4：表示IPv4地址。 - 6：表示IPv6地址。  [不支持IPv6，请勿设置为6。](tag:dt)
	IpVersion *int32 `json:"ip_version,omitempty"`
}

func (o EipInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EipInfo struct{}"
	}

	return strings.Join([]string{"EipInfo", string(data)}, " ")
}
