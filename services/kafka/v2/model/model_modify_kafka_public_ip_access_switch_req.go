package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ModifyKafkaPublicIpAccessSwitchReq struct {

	// **参数解释**： EIP地址。  **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	EipAddress *string `json:"eip_address,omitempty"`

	// **参数解释**： 公网带宽。  **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	PublicBoundwidth *int32 `json:"public_boundwidth,omitempty"`

	// **参数解释**： 公网IP的ID。  **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	PublicIpId *string `json:"publicIpId,omitempty"`
}

func (o ModifyKafkaPublicIpAccessSwitchReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyKafkaPublicIpAccessSwitchReq struct{}"
	}

	return strings.Join([]string{"ModifyKafkaPublicIpAccessSwitchReq", string(data)}, " ")
}
