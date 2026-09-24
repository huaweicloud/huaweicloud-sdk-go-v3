package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SwitchMySqlProxyEipRequestBody Proxy绑定解绑弹性公网IP请求体。
type SwitchMySqlProxyEipRequestBody struct {

	// **参数解释**：  待绑定的弹性公网IP地址。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	PublicIp string `json:"public_ip"`

	// **参数解释**：  弹性公网IP地址对应的ID。请求为绑定弹性公网IP时需传入该参数，请求为解绑弹性公网IP时无需传入。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	PublicIpId *string `json:"public_ip_id,omitempty"`

	// **参数解释**：  请求是否为绑定弹性公网IP。  **约束限制**：  不涉及。  **取值范围**：  - true：表示请求为绑定弹性公网IP。 - false：表示请求为解绑弹性公网IP。  **默认取值**：  不涉及。
	Bind string `json:"bind"`
}

func (o SwitchMySqlProxyEipRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchMySqlProxyEipRequestBody struct{}"
	}

	return strings.Join([]string{"SwitchMySqlProxyEipRequestBody", string(data)}, " ")
}
