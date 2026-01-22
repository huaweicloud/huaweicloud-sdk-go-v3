package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BindingsDetails struct {

	// **参数解释**： Exchange名称。 **取值范围**： 不涉及。
	Source *string `json:"source,omitempty"`

	// **参数解释**： 绑定目标的类型。 **取值范围**： - exchange：交换机。 - queue：队列。
	DestinationType *string `json:"destination_type,omitempty"`

	// **参数解释**： 绑定目标。 **取值范围**： 不涉及。
	Destination *string `json:"destination,omitempty"`

	// **参数解释**： 绑定键值。 **取值范围**： 不涉及。
	RoutingKey *string `json:"routing_key,omitempty"`

	// **参数解释**： 经过URL转译后routing_key。 **取值范围**： 不涉及。
	PropertiesKey *string `json:"properties_key,omitempty"`
}

func (o BindingsDetails) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BindingsDetails struct{}"
	}

	return strings.Join([]string{"BindingsDetails", string(data)}, " ")
}
