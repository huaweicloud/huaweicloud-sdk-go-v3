package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PublicIp **参数解释**： 弹性公网IP信息对象。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： null
type PublicIp struct {

	// **参数解释**： 弹性IP绑定类型。 **约束限制**： 不涉及。 **取值范围**： auto_assign：自动绑定。 not_use：暂未使用。 bind_existing ：使用已有。 **默认取值**： null
	PublicBindType *string `json:"public_bind_type,omitempty"`

	// **参数解释**： 弹性公网IP的id。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： null
	EipId *string `json:"eip_id,omitempty"`

	// **参数解释**： 弹性公网IP的类型。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： null
	IpType *string `json:"ip_type,omitempty"`

	// **参数解释**： 弹性公网IP的地址。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： null
	EipAddress *string `json:"eip_address,omitempty"`

	// **参数解释**： 弹性公网IP的带宽。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： null
	BandWidth *int32 `json:"band_width,omitempty"`

	// **参数解释**： 弹性公网IP的状态。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： null
	Status *string `json:"status,omitempty"`

	// **参数解释**： 弹性公网IP的错误码。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： null
	ErrorMessage *string `json:"error_message,omitempty"`
}

func (o PublicIp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PublicIp struct{}"
	}

	return strings.Join([]string{"PublicIp", string(data)}, " ")
}
