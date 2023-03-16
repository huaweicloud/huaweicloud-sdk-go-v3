package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 云堡垒机产品信息。
type ProductInfos struct {

	// 产品标识，通过订购询价接口获得。
	ProductId string `json:"product_id"`

	// 云服务类型，填写“hws.service.type.cbh”。
	CloudServiceType string `json:"cloud_service_type"`

	// 云堡垒机资源类型，填写“hws.resource.type.cbh.ins”。
	ResourceType string `json:"resource_type"`

	// 待创建云堡垒机规格ID，例如： - cbh.basic.50 - cbh.enhance.50 已上线的规格请参见《云堡垒机产品介绍》的“服务版本差异(https://support.huaweicloud.com/productdesc-cbh/cbh_01_0010.html)”章节。
	ResourceSpecCode string `json:"resource_spec_code"`

	// 资源容量度量标识。 - 15：Mbps（购买带宽时使用） - 17：GB（购买云硬盘时使用） - 14：个/次（购买堡垒机使用）
	ResourceSizeMeasureId string `json:"resource_size_measure_id"`

	// 资源容量大小。默认为1
	ResourceSize string `json:"resource_size"`
}

func (o ProductInfos) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProductInfos struct{}"
	}

	return strings.Join([]string{"ProductInfos", string(data)}, " ")
}
