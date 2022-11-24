package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 虚拟网关对象
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
	Type *VirtualGatewayType `json:"type,omitempty"`

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

type VirtualGatewayType struct {
	value string
}

type VirtualGatewayTypeEnum struct {
	DEFAULT VirtualGatewayType
}

func GetVirtualGatewayTypeEnum() VirtualGatewayTypeEnum {
	return VirtualGatewayTypeEnum{
		DEFAULT: VirtualGatewayType{
			value: "default",
		},
	}
}

func (c VirtualGatewayType) Value() string {
	return c.value
}

func (c VirtualGatewayType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *VirtualGatewayType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
