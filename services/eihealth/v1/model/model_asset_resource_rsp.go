package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AssetResourceRsp struct {

	// **参数解释**： 实例ID。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	Id *string `json:"id,omitempty"`

	// **参数解释**： 资源ID。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	ResourceId *string `json:"resource_id,omitempty"`

	// **参数解释**： 规格信息。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	SpecCode *string `json:"spec_code,omitempty"`

	// **参数解释**： 计费类型。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	ChargeMode *string `json:"charge_mode,omitempty"`

	Status *AssetResourceStatusEnum `json:"status,omitempty"`

	// **参数解释**： 购买时间，UTC时间。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	CreateTime *string `json:"create_time,omitempty"`
}

func (o AssetResourceRsp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssetResourceRsp struct{}"
	}

	return strings.Join([]string{"AssetResourceRsp", string(data)}, " ")
}
