package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListProductsRespDetail struct {

	// **参数解释**： 消息存储空间。 **取值范围**： 不涉及。
	Storage *string `json:"storage,omitempty"`

	// **参数解释**： 产品ID。 **取值范围**： 不涉及。
	ProductId *string `json:"product_id,omitempty"`

	// **参数解释**： 规格ID。 **取值范围**： 不涉及。
	SpecCode *string `json:"spec_code,omitempty"`

	// **参数解释**： IO信息。
	Io *[]ListProductsRespIo `json:"io,omitempty"`

	// **参数解释**： 资源售罄的可用区列表。 **取值范围**： 不涉及。
	UnavailableZones *[]string `json:"unavailable_zones,omitempty"`

	// **参数解释**： 有可用资源的可用区列表。 **取值范围**： 不涉及。
	AvailableZones *[]string `json:"available_zones,omitempty"`

	// **参数解释**： 该产品规格对应的虚拟机规格。 **取值范围**： 不涉及。
	EcsFlavorId *string `json:"ecs_flavor_id,omitempty"`

	// **参数解释**： 实例规格架构类型。当前仅支持X86。 **取值范围**： 不涉及。
	ArchType *string `json:"arch_type,omitempty"`
}

func (o ListProductsRespDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProductsRespDetail struct{}"
	}

	return strings.Join([]string{"ListProductsRespDetail", string(data)}, " ")
}
