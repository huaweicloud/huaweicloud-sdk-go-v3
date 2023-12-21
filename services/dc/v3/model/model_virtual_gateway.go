package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// VirtualGateway 虚拟网关对象
type VirtualGateway struct {

	// 虚拟网关的ID
	Id *string `json:"id,omitempty"`

	// 虚拟网关接入的VPC的ID
	VpcId *string `json:"vpc_id,omitempty"`

	// 实例所属项目ID。
	TenantId *string `json:"tenant_id,omitempty"`

	// 虚拟网关的名字
	Name *string `json:"name,omitempty"`

	// 虚拟网关的描述
	Description *string `json:"description,omitempty"`

	// 虚拟网关类型：default
	Type *string `json:"type,omitempty"`

	// 虚拟网关到访问云上服务IPv4子网列表，通常是vpc的cidrs
	LocalEpGroup *[]string `json:"local_ep_group,omitempty"`

	// 预留字段用于虚拟网关到访问云上服务IPv6子网列表，通常是vpc的cidrs
	LocalEpGroupIpv6 *[]string `json:"local_ep_group_ipv6,omitempty"`

	// 管理状态：true或false
	AdminStateUp *bool `json:"admin_state_up,omitempty"`

	// 操作状态，合法值是：ACTIVE，DOWN，BUILD，ERROR，PENDING_CREATE，PENDING_UPDATE，PENDING_DELETE
	Status *string `json:"status,omitempty"`

	// 虚拟网关本地的BGP自冶域号(asn)
	BgpAsn *int32 `json:"bgp_asn,omitempty"`

	// 实例所属企业项目ID
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 归属的设备ID
	DeviceId *string `json:"device_id,omitempty"`

	// 归属的冗余设备ID
	RedundantDeviceId *string `json:"redundant_device_id,omitempty"`

	// 归属的可用区对应的边界组(public border group)，标识是否homezone局点。
	PublicBorderGroup *string `json:"public_border_group,omitempty"`

	// 标签信息
	Tags *[]Tag `json:"tags,omitempty"`
}

func (o VirtualGateway) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VirtualGateway struct{}"
	}

	return strings.Join([]string{"VirtualGateway", string(data)}, " ")
}
