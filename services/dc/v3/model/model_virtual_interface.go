package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// VirtualInterface 虚拟接口对象
type VirtualInterface struct {

	// 虚拟接口的ID
	Id *string `json:"id,omitempty"`

	// 虚拟接口的名字
	Name *string `json:"name,omitempty"`

	// 管理状态：true或false
	AdminStateUp *bool `json:"admin_state_up,omitempty"`

	// 虚拟接口接入带宽
	Bandwidth *int32 `json:"bandwidth,omitempty"`

	// 虚拟接口创建时间
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// 虚拟接口的描述
	Description *string `json:"description,omitempty"`

	// 物理专线的ID
	DirectConnectId *string `json:"direct_connect_id,omitempty"`

	// 接入网关的类型：包括VGW,GDGW,LGW等
	ServiceType *VirtualInterfaceServiceType `json:"service_type,omitempty"`

	// 操作状态，合法值是：ACTIVE，DOWN，BUILD，ERROR，PENDING_CREATE，PENDING_UPDATE，PENDING_DELETE，DELETED，AUTHORIZATION，REJECTED
	Status *string `json:"status,omitempty"`

	// 租户ID
	TenantId *string `json:"tenant_id,omitempty"`

	// 表示接口类型：private
	Type *VirtualInterfaceType `json:"type,omitempty"`

	// 虚拟网关的ID
	VgwId *string `json:"vgw_id,omitempty"`

	// 同用户网关对接的vlan, 配置范围0-3999
	Vlan *int32 `json:"vlan,omitempty"`

	// VIF远端子网路由配置规格
	RouteLimit *int32 `json:"route_limit,omitempty"`

	// 是否使能nqa功能：true或false
	EnableNqa *bool `json:"enable_nqa,omitempty"`

	// 是否使能nqa功能：true或false
	EnableBfd *bool `json:"enable_bfd,omitempty"`

	// VIF关联的链路聚合组ID
	LagId *string `json:"lag_id,omitempty"`

	// 归属的设备ID
	DeviceId *string `json:"device_id,omitempty"`

	// 实例所属企业项目ID
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 标签信息
	Tags *[]Tag `json:"tags,omitempty"`

	// 云侧网关IPv4接口地址，该字段现已经移到vifpeer参数列表中，未来将会废弃。
	LocalGatewayV4Ip *string `json:"local_gateway_v4_ip,omitempty"`

	// 客户侧网关IPv4接口地址，该字段现已经移到vifpeer参数列表中，未来将会废弃。
	RemoteGatewayV4Ip *string `json:"remote_gateway_v4_ip,omitempty"`

	// 归属的IES站点的ID[（功能暂不支持）](tag:dt)
	IesId *string `json:"ies_id,omitempty"`

	// 如果资源的状态是Error的情况下，该参数会显示相关错误信息。
	Reason *string `json:"reason,omitempty"`

	// 标识虚拟接口是否开启限速
	RateLimit *bool `json:"rate_limit,omitempty"`

	// 接口的地址簇类型，ipv4，ipv6。该字段现已迁移到vifpeer参数列表中，未来将会废弃。
	AddressFamily *string `json:"address_family,omitempty"`

	// 云侧网关IPv6接口地址，该字段现已迁移到vifpeer参数列表中，未来将会废弃。
	LocalGatewayV6Ip *string `json:"local_gateway_v6_ip,omitempty"`

	// 客户侧网关IPv6接口地址，该字段现已迁移到vifpeer参数列表中，未来将会废弃。
	RemoteGatewayV6Ip *string `json:"remote_gateway_v6_ip,omitempty"`

	// 本地网关的ID，用于IES场景。[（功能暂不支持）](tag:dt)
	LgwId *string `json:"lgw_id,omitempty"`

	// 虚拟接口关联的网关的ID
	GatewayId *string `json:"gateway_id,omitempty"`

	// 远端子网列表，记录租户侧的cidrs。该字段现已迁移到vifpeer参数列表中，未来将会废弃。
	RemoteEpGroup *[]string `json:"remote_ep_group,omitempty"`

	// 该字段用于公网专线接口，表示租户可以访问云上公网服务地址列表。该字段现已迁移到vifpeer参数列表中，未来将会废弃。
	ServiceEpGroup *[]string `json:"service_ep_group,omitempty"`

	// BGP的路由配置规格
	BgpRouteLimit *int32 `json:"bgp_route_limit,omitempty"`

	// 虚拟接口的优先级，支持两种优先级状态normal和low。 接口优先级相同时表示负载关系，接口优先级不同时表示主备关系，出云流量优先转到优先级更高的normal接口。 目前仅BGP模式接口支持。
	Priority *VirtualInterfacePriority `json:"priority,omitempty"`

	// vif的Peer的相关信息
	VifPeers *[]VifPeer `json:"vif_peers,omitempty"`

	ExtendAttribute *VifExtendAttribute `json:"extend_attribute,omitempty"`
}

func (o VirtualInterface) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VirtualInterface struct{}"
	}

	return strings.Join([]string{"VirtualInterface", string(data)}, " ")
}

type VirtualInterfaceServiceType struct {
	value string
}

type VirtualInterfaceServiceTypeEnum struct {
	VGW  VirtualInterfaceServiceType
	GDGW VirtualInterfaceServiceType
	LGW  VirtualInterfaceServiceType
}

func GetVirtualInterfaceServiceTypeEnum() VirtualInterfaceServiceTypeEnum {
	return VirtualInterfaceServiceTypeEnum{
		VGW: VirtualInterfaceServiceType{
			value: "VGW",
		},
		GDGW: VirtualInterfaceServiceType{
			value: "GDGW",
		},
		LGW: VirtualInterfaceServiceType{
			value: "LGW",
		},
	}
}

func (c VirtualInterfaceServiceType) Value() string {
	return c.value
}

func (c VirtualInterfaceServiceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *VirtualInterfaceServiceType) UnmarshalJSON(b []byte) error {
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

type VirtualInterfaceType struct {
	value string
}

type VirtualInterfaceTypeEnum struct {
	PRIVATE VirtualInterfaceType
	PUBLIC  VirtualInterfaceType
}

func GetVirtualInterfaceTypeEnum() VirtualInterfaceTypeEnum {
	return VirtualInterfaceTypeEnum{
		PRIVATE: VirtualInterfaceType{
			value: "private",
		},
		PUBLIC: VirtualInterfaceType{
			value: "public",
		},
	}
}

func (c VirtualInterfaceType) Value() string {
	return c.value
}

func (c VirtualInterfaceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *VirtualInterfaceType) UnmarshalJSON(b []byte) error {
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

type VirtualInterfacePriority struct {
	value string
}

type VirtualInterfacePriorityEnum struct {
	NORMAL VirtualInterfacePriority
	LOW    VirtualInterfacePriority
}

func GetVirtualInterfacePriorityEnum() VirtualInterfacePriorityEnum {
	return VirtualInterfacePriorityEnum{
		NORMAL: VirtualInterfacePriority{
			value: "normal",
		},
		LOW: VirtualInterfacePriority{
			value: "low",
		},
	}
}

func (c VirtualInterfacePriority) Value() string {
	return c.value
}

func (c VirtualInterfacePriority) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *VirtualInterfacePriority) UnmarshalJSON(b []byte) error {
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
