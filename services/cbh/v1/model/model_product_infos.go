package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 产品信息
type ProductInfos struct {

	// 产品标识，通过订购询价接口获得
	ProductId string `json:"product_id"`

	// CBH云服务类型
	CloudServiceType string `json:"cloud_service_type"`

	// CBH资源类型
	ResourceType string `json:"resource_type"`

	// CBH资源规格
	ResourceSpecCode string `json:"resource_spec_code"`

	// 可用区id
	AvailableZoneId string `json:"available_zone_id"`

	// 资源容量度量标识
	ResourceSizeMeasureId string `json:"resource_size_measure_id"`

	// 资源容量大小
	ResourceSize string `json:"resource_size"`
}

func (o ProductInfos) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProductInfos struct{}"
	}

	return strings.Join([]string{"ProductInfos", string(data)}, " ")
}
