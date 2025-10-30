package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ServerOffering struct {

	// 服务器销售商品id
	Id *string `json:"id,omitempty"`

	// 地区编码，表示允许在这个地区购买部署
	ZoneCode *string `json:"zone_code,omitempty"`

	// 销售场景编码
	SceneCode *string `json:"scene_code,omitempty"`

	Status *OfferingStatus `json:"status,omitempty"`

	Spec *ServerSpec `json:"spec,omitempty"`

	ProductInfo *ProductInfo `json:"product_info,omitempty"`

	SaleCycles *[]SaleCycle `json:"sale_cycles,omitempty"`
}

func (o ServerOffering) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServerOffering struct{}"
	}

	return strings.Join([]string{"ServerOffering", string(data)}, " ")
}
