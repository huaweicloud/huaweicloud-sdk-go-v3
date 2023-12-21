package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// VifPeer 虚拟接口对等体对象
type VifPeer struct {

	// 资源ID
	Id *string `json:"id,omitempty"`

	// 归属租户ID
	TenantId *string `json:"tenant_id,omitempty"`

	// VIF对等体名字
	Name *string `json:"name,omitempty"`

	// VIF对等体名字描述信息
	Description *string `json:"description,omitempty"`

	// 接口的地址簇类型，ipv4，ipv6
	AddressFamily *string `json:"address_family,omitempty"`

	// VIF对等体云侧接口地址
	LocalGatewayIp *string `json:"local_gateway_ip,omitempty"`

	// VIF对等体客户侧接口地址
	RemoteGatewayIp *string `json:"remote_gateway_ip,omitempty"`

	// 路由模式：static/bgp
	RouteMode *VifPeerRouteMode `json:"route_mode,omitempty"`

	// BGP邻居的AS号
	BgpAsn *int32 `json:"bgp_asn,omitempty"`

	// BGP邻居的MD5密码
	BgpMd5 *string `json:"bgp_md5,omitempty"`

	// 远端子网列表，记录租户侧的cidrs
	RemoteEpGroup *[]string `json:"remote_ep_group,omitempty"`

	// 该字段用于公网专线接口,表示租户可以访问云上公网服务地址列表
	ServiceEpGroup *[]string `json:"service_ep_group,omitempty"`

	// 归属的设备ID
	DeviceId *string `json:"device_id,omitempty"`

	// BGP的路由配置规格
	BgpRouteLimit *int32 `json:"bgp_route_limit,omitempty"`

	// 接口BGP协议状态,如果是静态路由接口则状态为 null
	BgpStatus *string `json:"bgp_status,omitempty"`

	// VIF对等体状态
	Status *string `json:"status,omitempty"`

	// vif对等体对应的虚拟接口ID
	VifId *string `json:"vif_id,omitempty"`

	// 路由模式为bgp：receive_route_num值为接收搭配BGP的路由数目； 路由模式为static：该字段无意义，值为-1； 注：若早期接入华为云的部分用户无法获取该字段信息，如需要请联系客服迁移专线端口。
	ReceiveRouteNum *int32 `json:"receive_route_num,omitempty"`

	// 是否使能nqa功能：true或false
	EnableNqa *bool `json:"enable_nqa,omitempty"`

	// 是否使能bfd功能：true或false
	EnableBfd *bool `json:"enable_bfd,omitempty"`
}

func (o VifPeer) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VifPeer struct{}"
	}

	return strings.Join([]string{"VifPeer", string(data)}, " ")
}

type VifPeerRouteMode struct {
	value string
}

type VifPeerRouteModeEnum struct {
	BGP    VifPeerRouteMode
	STATIC VifPeerRouteMode
}

func GetVifPeerRouteModeEnum() VifPeerRouteModeEnum {
	return VifPeerRouteModeEnum{
		BGP: VifPeerRouteMode{
			value: "bgp",
		},
		STATIC: VifPeerRouteMode{
			value: "static",
		},
	}
}

func (c VifPeerRouteMode) Value() string {
	return c.value
}

func (c VifPeerRouteMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *VifPeerRouteMode) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
