package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateVifPeer 创建虚拟接口对等体参数
type CreateVifPeer struct {

	// VIF对等体名字
	Name string `json:"name"`

	// VIF对等体名字描述信息
	Description *string `json:"description,omitempty"`

	// 接口的地址簇类型，ipv4，ipv6
	AddressFamily string `json:"address_family"`

	// VIF对等体云侧接口地址
	LocalGatewayIp string `json:"local_gateway_ip"`

	// VIF对等体客户侧接口地址
	RemoteGatewayIp string `json:"remote_gateway_ip"`

	// 路由模式：static/bgp
	RouteMode *CreateVifPeerRouteMode `json:"route_mode,omitempty"`

	// BGP邻居的AS号
	BgpAsn *int32 `json:"bgp_asn,omitempty"`

	// BGP邻居的MD5密码
	BgpMd5 *string `json:"bgp_md5,omitempty"`

	// 远端子网列表，记录租户侧的cidrs
	RemoteEpGroup *[]string `json:"remote_ep_group,omitempty"`

	// vif对等体对应的虚拟接口ID
	VifId string `json:"vif_id"`
}

func (o CreateVifPeer) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVifPeer struct{}"
	}

	return strings.Join([]string{"CreateVifPeer", string(data)}, " ")
}

type CreateVifPeerRouteMode struct {
	value string
}

type CreateVifPeerRouteModeEnum struct {
	BGP    CreateVifPeerRouteMode
	STATIC CreateVifPeerRouteMode
}

func GetCreateVifPeerRouteModeEnum() CreateVifPeerRouteModeEnum {
	return CreateVifPeerRouteModeEnum{
		BGP: CreateVifPeerRouteMode{
			value: "bgp",
		},
		STATIC: CreateVifPeerRouteMode{
			value: "static",
		},
	}
}

func (c CreateVifPeerRouteMode) Value() string {
	return c.value
}

func (c CreateVifPeerRouteMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateVifPeerRouteMode) UnmarshalJSON(b []byte) error {
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
