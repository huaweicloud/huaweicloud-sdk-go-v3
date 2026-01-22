package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ClientData struct {

	// **参数解释**： 客户端语言。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Language *string `json:"language,omitempty"`

	// **参数解释**： 客户端版本。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Version *string `json:"version,omitempty"`

	// **参数解释**： 客户端ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	ClientId *string `json:"client_id,omitempty"`

	// **参数解释**： 客户端地址。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	ClientAddr *string `json:"client_addr,omitempty"`

	// **参数解释**： 订阅关系列表。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Subscriptions *[]Subscription `json:"subscriptions,omitempty"`
}

func (o ClientData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClientData struct{}"
	}

	return strings.Join([]string{"ClientData", string(data)}, " ")
}
