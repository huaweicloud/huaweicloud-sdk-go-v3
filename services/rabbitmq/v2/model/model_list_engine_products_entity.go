package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEngineProductsEntity **参数解释**： 产品规格的详细信息。 **取值范围**： 不涉及。
type ListEngineProductsEntity struct {

	// **参数解释**： 产品类型。 **取值范围**： - single：单机。    - cluster：集群。 [- single.professional：单机专业版，AMQP版本产品类型。](tag:hws,hws_hk,hws_eu) [- cluster.professional：集群专业版，AMQP版本产品类型。](tag:hws,hws_hk,hws_eu)
	Type *string `json:"type,omitempty"`

	// **参数解释**： 产品ID。 **取值范围**： 不涉及。
	ProductId *string `json:"product_id,omitempty"`

	// **参数解释**： 底层资源类型。 **取值范围**： 不涉及。
	EcsFlavorId *string `json:"ecs_flavor_id,omitempty"`

	// **参数解释**： 账单计费类型。 **取值范围**： 不涉及。
	BillingCode *string `json:"billing_code,omitempty"`

	// **参数解释**： CPU架构。
	ArchTypes *[]string `json:"arch_types,omitempty"`

	// **参数解释**： 计费模式。
	ChargingMode *[]string `json:"charging_mode,omitempty"`

	// **参数解释**： 支持的磁盘IO类型列表。
	Ios *[]ListEngineIosEntity `json:"ios,omitempty"`

	// **参数解释**： 当前规格实例支持的功能特性列表。
	SupportFeatures *[]interface{} `json:"support_features,omitempty"`

	Properties *ListEnginePropertiesEntity `json:"properties,omitempty"`
}

func (o ListEngineProductsEntity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEngineProductsEntity struct{}"
	}

	return strings.Join([]string{"ListEngineProductsEntity", string(data)}, " ")
}
