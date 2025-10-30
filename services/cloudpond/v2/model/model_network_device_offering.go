package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type NetworkDeviceOffering struct {

	// 销售策略ID
	Id *string `json:"id,omitempty"`

	// 地区编码，表示允许在这个地区购买部署
	ZoneCode *string `json:"zone_code,omitempty"`

	// 销售场景编码
	SceneCode *string `json:"scene_code,omitempty"`

	Status *OfferingStatus `json:"status,omitempty"`

	Spec *NetworkDeviceSpec `json:"spec,omitempty"`

	ProductInfo *ProductInfo `json:"product_info,omitempty"`

	SaleCycles *[]SaleCycle `json:"sale_cycles,omitempty"`
}

func (o NetworkDeviceOffering) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NetworkDeviceOffering struct{}"
	}

	return strings.Join([]string{"NetworkDeviceOffering", string(data)}, " ")
}
