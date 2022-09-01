package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type NeutronListFloatingIpsRequest struct {

	// 每页显示的条目数量。
	Limit *int32 `json:"limit,omitempty" xml:"limit"`

	// 取值为上一页数据的最后一条记录的id，当marker参数为无效id时，response将响应错误码400
	Marker *string `json:"marker,omitempty" xml:"marker"`

	// False/True，是否设置分页的顺序。
	PageReverse *bool `json:"page_reverse,omitempty" xml:"page_reverse"`

	// 浮动IP的id。
	Id *string `json:"id,omitempty" xml:"id"`

	// 浮动IP地址。
	FloatingIpAddress *string `json:"floating_ip_address,omitempty" xml:"floating_ip_address"`

	// 所属路由器id。
	RouterId *string `json:"router_id,omitempty" xml:"router_id"`

	// 端口id。
	PortId *string `json:"port_id,omitempty" xml:"port_id"`

	// 关联端口的私有IP地址。
	FixedIpAddress *string `json:"fixed_ip_address,omitempty" xml:"fixed_ip_address"`

	// 项目ID。
	TenantId *string `json:"tenant_id,omitempty" xml:"tenant_id"`

	// 外部网络的id。只能使用固定的外网，外部网络的信息请通过GET /v2.0/networks?router:external=True或GET /v2.0/networks?name={floating_network}或neutron net-external-list方式查询
	FloatingNetworkId *string `json:"floating_network_id,omitempty" xml:"floating_network_id"`
}

func (o NeutronListFloatingIpsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronListFloatingIpsRequest struct{}"
	}

	return strings.Join([]string{"NeutronListFloatingIpsRequest", string(data)}, " ")
}
