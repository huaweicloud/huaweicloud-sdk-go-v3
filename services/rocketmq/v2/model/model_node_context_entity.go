package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type NodeContextEntity struct {

	// **参数解释**： ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释**： broker名称。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	BrokerName *string `json:"broker_name,omitempty"`

	// **参数解释**： broker的ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	BrokerId *int32 `json:"broker_id,omitempty"`

	// **参数解释**： 地址。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Address *string `json:"address,omitempty"`

	// **参数解释**： 公网地址。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	PublicAddress *string `json:"public_address,omitempty"`
}

func (o NodeContextEntity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodeContextEntity struct{}"
	}

	return strings.Join([]string{"NodeContextEntity", string(data)}, " ")
}
