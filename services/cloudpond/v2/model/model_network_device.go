package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NetworkDevice 网络设备详情
type NetworkDevice struct {

	// 网络设备资源ID
	Id *string `json:"id,omitempty"`

	// 账号ID
	DomainId *string `json:"domain_id,omitempty"`

	// 站点ID
	EdgeSiteId *string `json:"edge_site_id,omitempty"`

	// 名称
	Name *string `json:"name,omitempty"`

	Status *NetworkDeviceStatus `json:"status,omitempty"`

	// 商品ID
	OfferingId *string `json:"offering_id,omitempty"`

	Spec *NetworkDeviceSpec `json:"spec,omitempty"`

	MarketOptions *MarketOptions `json:"market_options,omitempty"`

	ProductInfo *ProductInfo `json:"product_info,omitempty"`

	Location *LayoutLocation `json:"location,omitempty"`

	// 创建时间
	CreatedAt *sdktime.SdkTime `json:"created_at,omitempty"`

	// 更新时间
	UpdatedAt *sdktime.SdkTime `json:"updated_at,omitempty"`
}

func (o NetworkDevice) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NetworkDevice struct{}"
	}

	return strings.Join([]string{"NetworkDevice", string(data)}, " ")
}
