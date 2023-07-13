package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NeutronListSubnetsRequest Request Object
type NeutronListSubnetsRequest struct {

	// 每页返回的个数
	Limit *int32 `json:"limit,omitempty"`

	// 分页查询起始的资源ID，为空时查询第一页
	Marker *string `json:"marker,omitempty"`

	// 按照子网对应的ID过滤查询
	Id *string `json:"id,omitempty"`

	// 按照子网的cidr过滤查询
	Cidr *string `json:"cidr,omitempty"`

	// 按照子网的名称过滤查询
	Name *string `json:"name,omitempty"`

	// 按照子网是否开启dhcp过滤查询，取值范围：true or false
	EnableDhcp *bool `json:"enable_dhcp,omitempty"`

	// 按照子网所属network_id过滤查询
	NetworkId *string `json:"network_id,omitempty"`

	// 按照子网的IP协议版本过滤查询
	IpVersion *int32 `json:"ip_version,omitempty"`

	// 按照子网的网关IP过滤查询
	GatewayIp *string `json:"gateway_ip,omitempty"`

	// 按照子网所属的项目ID过滤查询
	TenantId *string `json:"tenant_id,omitempty"`
}

func (o NeutronListSubnetsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronListSubnetsRequest struct{}"
	}

	return strings.Join([]string{"NeutronListSubnetsRequest", string(data)}, " ")
}
