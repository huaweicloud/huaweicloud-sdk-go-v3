/*
    * EIP
    *
    * 云服务接口
    *
*/

package model
import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
)

// floatingip对象
type FloatingIpResp struct {
	// 关联端口的私有IP地址。
	FixedIpAddress string `json:"fixed_ip_address,omitempty"`
	// 浮动IP地址。
	FloatingIpAddress string `json:"floating_ip_address,omitempty"`
	// 外部网络的id。只能使用固定的外网，外部网络的信息请通过GET /v2.0/networks?router:external=True或GET /v2.0/networks?name={floating_network}或neutron net-external-list方式查询。
	FloatingNetworkId string `json:"floating_network_id,omitempty"`
	// 浮动IP地址的id。
	Id string `json:"id,omitempty"`
	// 端口id。
	PortId string `json:"port_id,omitempty"`
	// 所属路由器id。
	RouterId string `json:"router_id,omitempty"`
	// 网络状态，可以为ACTIVE， DOWN或ERROR。  DOWN：未绑定  ACTIVE：绑定  ERROR：异常
	Status string `json:"status,omitempty"`
	// 项目id。
	TenantId string `json:"tenant_id,omitempty"`
	// 项目id。
	ProjectId string `json:"project_id,omitempty"`
	// DNS名称(目前仅广州局点支持)
	DnsName string `json:"dns_name,omitempty"`
	// DNS域地址(目前仅广州局点支持)
	DnsDomain string `json:"dns_domain,omitempty"`
	// 资源创建时间  采用UTC时间  格式：YYYY-MM-DDTHH:MM:SS
	CreatedAt *sdktime.SdkTime `json:"created_at,omitempty"`
	// 资源更新时间  采用UTC时间  格式：YYYY-MM-DDTHH:MM:SS
	UpdatedAt *sdktime.SdkTime `json:"updated_at,omitempty"`
}
