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

	// 虚拟接口更新时间
	UpdateTime *sdktime.SdkTime `json:"update_time,omitempty"`

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
