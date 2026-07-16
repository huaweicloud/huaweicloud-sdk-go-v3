package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ServerPublicIp struct {

	// **参数解释**：EIP的ID。 **约束限制**：必填。 **取值范围**：1 - 64字符。 **默认取值**：不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释**：EIP的IP地址。 **约束限制**：必填。 **取值范围**：1 - 64字符，标准IPv4地址。 **默认取值**：不涉及。
	Address *string `json:"address,omitempty"`
}

func (o ServerPublicIp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServerPublicIp struct{}"
	}

	return strings.Join([]string{"ServerPublicIp", string(data)}, " ")
}
