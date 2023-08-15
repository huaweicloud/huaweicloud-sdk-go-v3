package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateVirtualGateway 更新虚拟网关参数
type UpdateVirtualGateway struct {

	// 更新虚拟网关的名字
	Name *string `json:"name,omitempty"`

	// 虚拟网关的描述信息
	Description *string `json:"description,omitempty"`

	// 虚拟网关到访问云上服务IPv4子网列表，通常是vpc的cidrs
	LocalEpGroup *[]string `json:"local_ep_group,omitempty"`

	// 虚拟网关到访问云上服务IPv6子网列表，通常是vpc的cidrs
	LocalEpGroupIpv6 *[]string `json:"local_ep_group_ipv6,omitempty"`
}

func (o UpdateVirtualGateway) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateVirtualGateway struct{}"
	}

	return strings.Join([]string{"UpdateVirtualGateway", string(data)}, " ")
}
