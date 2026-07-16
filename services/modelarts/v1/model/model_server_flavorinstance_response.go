package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ServerFlavorinstanceResponse **参数解释**：Lite Server规格详情。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及
type ServerFlavorinstanceResponse struct {

	// **参数解释**：CPU架构。 **约束限制**：不涉及。 **取值范围**： - X86：CPU架构为X86 - ARM：CPU架构为ARM  **默认取值**：不涉及。
	Arch *string `json:"arch,omitempty"`

	// **参数解释**：分区名。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	AvailabilityZone *string `json:"availability_zone,omitempty"`

	// **参数解释**：计费模式。 **约束限制**：不涉及。 **取值范围**： - PRE_PAID：计费模式为包年/包月 - POST_PAID：计费模式为按需计费 **默认取值**：不涉及。
	ChargingMode *string `json:"charging_mode,omitempty"`

	// **参数解释**：数量。 **约束限制**：不涉及。 **默认取值**：不涉及。
	Count *int32 `json:"count,omitempty"`

	// **参数解释**：规格名称。 **约束限制**：不涉及。 **默认取值**：不涉及。
	Flavor *string `json:"flavor,omitempty"`

	// **参数解释**：规格类型。 **约束限制**：不涉及。 **默认取值**：不涉及。
	FlavorType *string `json:"flavor_type,omitempty"`

	// **参数解释**：roce数量。 **约束限制**：不涉及。 **默认取值**：不涉及。
	RoceNum *int32 `json:"roce_num,omitempty"`

	// **参数解释**：服务类型。 **约束限制**：不涉及。 **取值范围**： - BMS：资源类型为裸金属服务器 - ECS：资源类型为弹性云服务器 - HPS：资源类型为超节点服务器  **默认取值**：不涉及。
	ServerType *string `json:"server_type,omitempty"`

	// **参数解释**：计费码。 **约束限制**：不涉及。 **默认取值**：不涉及。
	SkuCode *string `json:"sku_code,omitempty"`

	// **参数解释**：规格详情。 **约束限制**：不涉及。 **默认取值**：不涉及。
	Specification *string `json:"specification,omitempty"`

	// **参数解释**：状态。 **约束限制**：不涉及。 **默认取值**：不涉及。
	Status *string `json:"status,omitempty"`

	// **参数解释**：是否售罄。 **约束限制**：不涉及。 **取值范围**： - true：已售罄 - false：未售罄  **默认取值**：false。
	IsSoldOut *bool `json:"is_sold_out,omitempty"`
}

func (o ServerFlavorinstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServerFlavorinstanceResponse struct{}"
	}

	return strings.Join([]string{"ServerFlavorinstanceResponse", string(data)}, " ")
}
