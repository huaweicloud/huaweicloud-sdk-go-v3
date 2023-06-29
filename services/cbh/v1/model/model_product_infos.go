package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ProductInfos 云堡垒机产品信息。
type ProductInfos struct {

	// 产品标识，通过订购询价接口获得。
	ProductId string `json:"product_id"`

	// 云服务类型，填写“hws.service.type.cbh”。
	CloudServiceType string `json:"cloud_service_type"`

	// 云堡垒机资源类型，填写“hws.resource.type.cbh.ins”。
	ResourceType string `json:"resource_type"`

	// 待创建云堡垒机规格ID，例如： - cbh.basic.50 - cbh.enhance.50 已上线的规格请参见《云堡垒机产品介绍》的[服务版本差异](https://support.huaweicloud.com/productdesc-cbh/cbh_01_0010.html)章节。
	ResourceSpecCode string `json:"resource_spec_code"`
}

func (o ProductInfos) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProductInfos struct{}"
	}

	return strings.Join([]string{"ProductInfos", string(data)}, " ")
}
