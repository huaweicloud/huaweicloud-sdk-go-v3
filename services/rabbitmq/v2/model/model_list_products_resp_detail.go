package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListProductsRespDetail struct {

	// 消息存储空间。
	Storage *string `json:"storage,omitempty" xml:"storage"`

	// 产品ID。
	ProductId *string `json:"product_id,omitempty" xml:"product_id"`

	// 规格ID。
	SpecCode *string `json:"spec_code,omitempty" xml:"spec_code"`

	// IO信息。
	Io *[]ListProductsRespIo `json:"io,omitempty" xml:"io"`

	// 资源售罄的可用区列表。
	UnavailableZones *[]string `json:"unavailable_zones,omitempty" xml:"unavailable_zones"`

	// 有可用资源的可用区列表。
	AvailableZones *[]string `json:"available_zones,omitempty" xml:"available_zones"`

	// 该产品规格对应的虚拟机规格。
	EcsFlavorId *string `json:"ecs_flavor_id,omitempty" xml:"ecs_flavor_id"`

	// 实例规格架构类型。当前仅支持X86。
	ArchType *string `json:"arch_type,omitempty" xml:"arch_type"`
}

func (o ListProductsRespDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProductsRespDetail struct{}"
	}

	return strings.Join([]string{"ListProductsRespDetail", string(data)}, " ")
}
