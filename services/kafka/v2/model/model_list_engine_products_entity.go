package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 产品规格的详细信息。
type ListEngineProductsEntity struct {

	// 产品类型。当前产品类型有单机和集群。
	Type *string `json:"type,omitempty" xml:"type"`

	// 产品ID。
	ProductId *string `json:"product_id,omitempty" xml:"product_id"`

	// 底层资源类型。
	EcsFlavorId *string `json:"ecs_flavor_id,omitempty" xml:"ecs_flavor_id"`

	// 账单计费类型。
	BillingCode *string `json:"billing_code,omitempty" xml:"billing_code"`

	// CPU架构。
	ArchTypes *[]string `json:"arch_types,omitempty" xml:"arch_types"`

	// 计费模式。   - monthly：包年/包月类型。   - hourly：按需类型。
	ChargingMode *[]string `json:"charging_mode,omitempty" xml:"charging_mode"`

	// 支持的磁盘IO类型列表。
	Ios *[]ListEngineIosEntity `json:"ios,omitempty" xml:"ios"`

	// 当前规格实例支持的功能特性列表。
	SupportFeatures *[]ListEngineSupportFeaturesEntity `json:"support_features,omitempty" xml:"support_features"`

	Properties *ListEnginePropertiesEntity `json:"properties,omitempty" xml:"properties"`
}

func (o ListEngineProductsEntity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEngineProductsEntity struct{}"
	}

	return strings.Join([]string{"ListEngineProductsEntity", string(data)}, " ")
}
