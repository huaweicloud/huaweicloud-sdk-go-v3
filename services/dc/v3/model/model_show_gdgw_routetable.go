package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowGdgwRoutetable 路由表详细
type ShowGdgwRoutetable struct {

	// 唯一ID
	Id *string `json:"id,omitempty"`

	// 租户ID
	TenantId *string `json:"tenant_id,omitempty"`

	// 网关ID
	GatewayId *string `json:"gateway_id,omitempty"`

	// 描述信息
	Destination *string `json:"destination,omitempty"`

	// 下一跳ID
	Nexthop *string `json:"nexthop,omitempty"`

	// 类型
	Type *string `json:"type,omitempty"`

	// 获得模式
	ObtainMode *string `json:"obtain_mode,omitempty"`

	// 状态：ACTIVE-正常，ERROR-异常
	Status *string `json:"status,omitempty"`

	// 地址簇：ipv4 | ipv6
	AddressFamily *string `json:"address_family,omitempty"`

	// 描述信息
	Description *string `json:"description,omitempty"`
}

func (o ShowGdgwRoutetable) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowGdgwRoutetable struct{}"
	}

	return strings.Join([]string{"ShowGdgwRoutetable", string(data)}, " ")
}
