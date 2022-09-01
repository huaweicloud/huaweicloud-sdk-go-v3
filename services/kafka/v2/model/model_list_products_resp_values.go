package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListProductsRespValues struct {

	// 规格详情。
	Detail *[]ListProductsRespDetail `json:"detail,omitempty" xml:"detail"`

	// 实例类型。
	Name *string `json:"name,omitempty" xml:"name"`

	// 资源售罄的可用区列表。
	UnavailableZones *[]string `json:"unavailable_zones,omitempty" xml:"unavailable_zones"`

	// 有可用资源的可用区列表。
	AvailableZones *[]string `json:"available_zones,omitempty" xml:"available_zones"`
}

func (o ListProductsRespValues) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProductsRespValues struct{}"
	}

	return strings.Join([]string{"ListProductsRespValues", string(data)}, " ")
}
