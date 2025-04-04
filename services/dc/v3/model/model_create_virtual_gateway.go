package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateVirtualGateway 创建虚拟网关对象参数
type CreateVirtualGateway struct {

	// 虚拟网关接入的VPC的ID[，当选择创建接入VPC的虚拟网关时必选。](tag:dt)
	VpcId string `json:"vpc_id"`

	// 虚拟网关接入的ER的ID，当选择创建接入ER的虚拟网关时必选。
	EnterpriseRouterId *string `json:"enterprise_router_id,omitempty"`

	// 虚拟网关名字
	Name *string `json:"name,omitempty"`

	// 虚拟网关的描述信息
	Description *string `json:"description,omitempty"`

	// 虚拟网关到访问云上服务IPv4子网列表，通常是vpc的cidrs[，当选择创建接入VPC的虚拟网关时必选。](tag:dt)
	LocalEpGroup []string `json:"local_ep_group"`

	// 预留字段用于虚拟网关到访问云上服务IPv6子网列表，通常是vpc的cidrs
	LocalEpGroupIpv6 *[]string `json:"local_ep_group_ipv6,omitempty"`

	// 虚拟网关本地的BGP自治域号(asn)
	BgpAsn *int32 `json:"bgp_asn,omitempty"`

	// 实例所属企业项目ID
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 标签信息
	Tags *[]Tag `json:"tags,omitempty"`
}

func (o CreateVirtualGateway) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVirtualGateway struct{}"
	}

	return strings.Join([]string{"CreateVirtualGateway", string(data)}, " ")
}
