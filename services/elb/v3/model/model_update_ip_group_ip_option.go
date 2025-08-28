package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateIpGroupIpOption **参数解释**：更新IP地址组IP请求参数。  **约束限制**：不涉及
type UpdateIpGroupIpOption struct {

	// **参数解释**：IP地址组中的IP信息。  **约束限制**：不涉及  **取值范围**：支持IPv4和IPv6地址类型，格式为IP地址、IP地址段或IP地址范围，IP地址范围格式为ip-ip，例如192.168.1.2-192.168.2.253或者2001:0DB8:02de::0e12-2001:0DB8:02de::0e13，终止IP需要大于起始IP。  **默认取值**：不涉及  [不支持IPv6，请勿设置为IPv6地址。](tag:dt)
	Ip string `json:"ip"`

	// **参数解释**：备注信息。  **约束限制**：不涉及  **取值范围**：长度为0-255个字符。  **默认取值**：不涉及
	Description *string `json:"description,omitempty"`
}

func (o UpdateIpGroupIpOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateIpGroupIpOption struct{}"
	}

	return strings.Join([]string{"UpdateIpGroupIpOption", string(data)}, " ")
}
