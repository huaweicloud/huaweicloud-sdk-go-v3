package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type IpGroupItem struct {

	// **参数解释**：  弹性负载均衡ip。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	Ip *string `json:"ip,omitempty"`

	// **参数解释**：  弹性负载均衡描述。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	Description *string `json:"description,omitempty"`
}

func (o IpGroupItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IpGroupItem struct{}"
	}

	return strings.Join([]string{"IpGroupItem", string(data)}, " ")
}
