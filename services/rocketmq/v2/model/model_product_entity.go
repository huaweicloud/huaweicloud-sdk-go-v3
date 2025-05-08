package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ProductEntity struct {

	// **参数解释**： 产品类型。 **取值范围**： - single：单机。 - cluster：集群。
	Type *string `json:"type,omitempty"`

	// **参数解释**： 产品ID。 **取值范围**： 不涉及。
	ProductId *string `json:"product_id,omitempty"`

	// **参数解释**： ECS的Flavor ID。 **取值范围**： 不涉及。
	EcsFlavorId *string `json:"ecs_flavor_id,omitempty"`

	// **参数解释**： CBC的规格码。 **取值范围**： 不涉及。
	BillingCode *string `json:"billing_code,omitempty"`

	// **参数解释**： 架构类型。
	ArchTypes *[]string `json:"arch_types,omitempty"`

	// **参数解释**： 计费模式。 **取值范围**： 不涉及。
	ChargingMode *interface{} `json:"charging_mode,omitempty"`

	// **参数解释**： 支持的io类型。 **取值范围**： 不涉及。
	Ios *interface{} `json:"ios,omitempty"`

	// **参数解释**： 支持的特性。 **取值范围**： 不涉及。
	SupportFeatures *interface{} `json:"support_features,omitempty"`

	// **参数解释**： 产品特性。 **取值范围**： 不涉及。
	Properties *interface{} `json:"properties,omitempty"`

	// **参数解释**： 可用分区。
	AvailableZones *[]string `json:"available_zones,omitempty"`

	// **参数解释**： 不可用分区。
	UnavailableZones *[]string `json:"unavailable_zones,omitempty"`

	// **参数解释**： 是否为擎天实例。 **取值范围**： - true：是 - false：否
	QingtianIncompatible *bool `json:"qingtian_incompatible,omitempty"`
}

func (o ProductEntity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProductEntity struct{}"
	}

	return strings.Join([]string{"ProductEntity", string(data)}, " ")
}
