package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateVirtualInterface 创建虚拟接口对象参数
type CreateVirtualInterface struct {

	// 虚拟接口名字
	Name *string `json:"name,omitempty"`

	// 虚拟接口描述信息
	Description *string `json:"description,omitempty"`

	// 虚拟接口关联的物理专线ID。在创建虚拟接口时，需要关联direct_connect_id或lag_id，当LAG功能在当前局点不支持时，该参数为必选。
	DirectConnectId *string `json:"direct_connect_id,omitempty"`

	// 虚拟接口的类型,private
	Type CreateVirtualInterfaceType `json:"type"`

	// 接入网关类型：VGW/GDGW/LGW
	ServiceType *CreateVirtualInterfaceServiceType `json:"service_type,omitempty"`

	// 对接客户侧vlan。当关联物理连接选择托管专线时，vlan需要和托管专线保持一致。
	Vlan int32 `json:"vlan"`

	// 虚拟接口接入带宽
	Bandwidth int32 `json:"bandwidth"`

	// 虚拟接口的优先级，支持两种优先级状态normal和low。 接口优先级相同时表示负载关系，接口优先级不同时表示主备关系，出云流量优先转到优先级更高的normal接口。 目前仅BGP模式接口支持。
	Priority *CreateVirtualInterfacePriority `json:"priority,omitempty"`

	// 云侧网关IPv4接口地址,如果address_family是IPv4，是必选参数
	LocalGatewayV4Ip *string `json:"local_gateway_v4_ip,omitempty"`

	// 客户侧网关IPv4接口地址,如果address_family是IPv4，是必选参数
	RemoteGatewayV4Ip *string `json:"remote_gateway_v4_ip,omitempty"`

	// 接口的地址簇类型，ipv4，ipv6
	AddressFamily *string `json:"address_family,omitempty"`

	// 云侧网关IPv6接口地址,如果address_family是IPv6，是必选参数
	LocalGatewayV6Ip *string `json:"local_gateway_v6_ip,omitempty"`

	// 客户侧网关IPv6接口地址,如果address_family是IPv6，是必选参数
	RemoteGatewayV6Ip *string `json:"remote_gateway_v6_ip,omitempty"`

	// 虚拟接口连接的虚拟网关的ID
	VgwId string `json:"vgw_id"`

	// 虚拟接口关联的网关的ID。 当虚拟接口接入网关类型为全域接入网关时该参数必填，gateway_id值为全域接入网关的id
	GatewayId *string `json:"gateway_id,omitempty"`

	// 路由模式：static/bgp
	RouteMode CreateVirtualInterfaceRouteMode `json:"route_mode"`

	// 客户侧BGP邻居的AS号
	BgpAsn *int32 `json:"bgp_asn,omitempty"`

	// BGP邻居的MD5密码
	BgpMd5 *string `json:"bgp_md5,omitempty"`

	// 远端子网列表，记录租户侧的cidrs
	RemoteEpGroup []string `json:"remote_ep_group"`

	// 访问公网服务的子网列表
	ServiceEpGroup *[]string `json:"service_ep_group,omitempty"`

	// 是否使能bfd功能：true或false。[（预留字段暂不支持）](tag:dt)
	EnableBfd *bool `json:"enable_bfd,omitempty"`

	// 是否使能nqa功能：true或false。[（预留字段暂不支持）](tag:dt)
	EnableNqa *bool `json:"enable_nqa,omitempty"`

	// 虚拟接口关联的链路聚合组ID
	LagId *string `json:"lag_id,omitempty"`

	// 目标的租户的ID,用于跨租户创建虚拟接口场景
	ResourceTenantId *string `json:"resource_tenant_id,omitempty"`

	// 实例所属企业项目ID
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 标签信息
	Tags *[]Tag `json:"tags,omitempty"`
}

func (o CreateVirtualInterface) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVirtualInterface struct{}"
	}

	return strings.Join([]string{"CreateVirtualInterface", string(data)}, " ")
}

type CreateVirtualInterfaceType struct {
	value string
}

type CreateVirtualInterfaceTypeEnum struct {
	PRIVATE CreateVirtualInterfaceType
	PUBLIC  CreateVirtualInterfaceType
}

func GetCreateVirtualInterfaceTypeEnum() CreateVirtualInterfaceTypeEnum {
	return CreateVirtualInterfaceTypeEnum{
		PRIVATE: CreateVirtualInterfaceType{
			value: "private",
		},
		PUBLIC: CreateVirtualInterfaceType{
			value: "public",
		},
	}
}

func (c CreateVirtualInterfaceType) Value() string {
	return c.value
}

func (c CreateVirtualInterfaceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateVirtualInterfaceType) UnmarshalJSON(b []byte) error {
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

type CreateVirtualInterfaceServiceType struct {
	value string
}

type CreateVirtualInterfaceServiceTypeEnum struct {
	VGW  CreateVirtualInterfaceServiceType
	GDGW CreateVirtualInterfaceServiceType
	LGW  CreateVirtualInterfaceServiceType
}

func GetCreateVirtualInterfaceServiceTypeEnum() CreateVirtualInterfaceServiceTypeEnum {
	return CreateVirtualInterfaceServiceTypeEnum{
		VGW: CreateVirtualInterfaceServiceType{
			value: "VGW",
		},
		GDGW: CreateVirtualInterfaceServiceType{
			value: "GDGW",
		},
		LGW: CreateVirtualInterfaceServiceType{
			value: "LGW",
		},
	}
}

func (c CreateVirtualInterfaceServiceType) Value() string {
	return c.value
}

func (c CreateVirtualInterfaceServiceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateVirtualInterfaceServiceType) UnmarshalJSON(b []byte) error {
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

type CreateVirtualInterfacePriority struct {
	value string
}

type CreateVirtualInterfacePriorityEnum struct {
	NORMAL CreateVirtualInterfacePriority
	LOW    CreateVirtualInterfacePriority
}

func GetCreateVirtualInterfacePriorityEnum() CreateVirtualInterfacePriorityEnum {
	return CreateVirtualInterfacePriorityEnum{
		NORMAL: CreateVirtualInterfacePriority{
			value: "normal",
		},
		LOW: CreateVirtualInterfacePriority{
			value: "low",
		},
	}
}

func (c CreateVirtualInterfacePriority) Value() string {
	return c.value
}

func (c CreateVirtualInterfacePriority) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateVirtualInterfacePriority) UnmarshalJSON(b []byte) error {
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

type CreateVirtualInterfaceRouteMode struct {
	value string
}

type CreateVirtualInterfaceRouteModeEnum struct {
	STATIC CreateVirtualInterfaceRouteMode
	BGP    CreateVirtualInterfaceRouteMode
}

func GetCreateVirtualInterfaceRouteModeEnum() CreateVirtualInterfaceRouteModeEnum {
	return CreateVirtualInterfaceRouteModeEnum{
		STATIC: CreateVirtualInterfaceRouteMode{
			value: "static",
		},
		BGP: CreateVirtualInterfaceRouteMode{
			value: "bgp",
		},
	}
}

func (c CreateVirtualInterfaceRouteMode) Value() string {
	return c.value
}

func (c CreateVirtualInterfaceRouteMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateVirtualInterfaceRouteMode) UnmarshalJSON(b []byte) error {
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
