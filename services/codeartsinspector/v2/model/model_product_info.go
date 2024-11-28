package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ProductInfo struct {

	// 用户购买的云服务的主云服务类型
	CloudServiceType string `json:"cloud_service_type"`

	// 产品标识，通过订购询价接口获得
	ProductId string `json:"product_id"`

	// 用户购买云服务产品的资源规格
	ResourceSpeccode string `json:"resource_speccode"`

	// 用户购买云服务产品的资源类型
	ResourceType string `json:"resource_type"`

	// 资源容量度量标识，购买vss服务时使用14： 15：Mbps（购买带宽时使用） 17：GB（购买云硬盘时使用） 14：个/次
	ResouceSizeMeasureId *int32 `json:"resouce_size_measure_id,omitempty"`

	// 资源容量大小
	ResourceSize *int32 `json:"resource_size,omitempty"`
}

func (o ProductInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProductInfo struct{}"
	}

	return strings.Join([]string{"ProductInfo", string(data)}, " ")
}
