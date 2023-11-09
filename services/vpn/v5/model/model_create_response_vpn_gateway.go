package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CreateResponseVpnGateway struct {

	// VPN网关ID
	Id *string `json:"id,omitempty"`

	// VPN网关名称
	Name *string `json:"name,omitempty"`

	// 关联模式
	AttachmentType *CreateResponseVpnGatewayAttachmentType `json:"attachment_type,omitempty"`

	// VPN网关所连接的ER实例的ID
	ErId *string `json:"er_id,omitempty"`

	// VPN网关所连接的VPC的ID
	VpcId *string `json:"vpc_id,omitempty"`

	// 本端子网
	LocalSubnets *[]string `json:"local_subnets,omitempty"`

	// VPN网关所使用的VPC子网ID
	ConnectSubnet *string `json:"connect_subnet,omitempty"`

	// VPN网关北向类型，默认为公网(public)
	NetworkType *CreateResponseVpnGatewayNetworkType `json:"network_type,omitempty"`

	// VPN网关北向接入VPC ID，不填时默认使用vpc_id字段的值
	AccessVpcId *string `json:"access_vpc_id,omitempty"`

	// VPN网关北向接入VPC中的接入子网ID
	AccessSubnetId *string `json:"access_subnet_id,omitempty"`

	// bgp所使用的asn号
	BgpAsn *int64 `json:"bgp_asn,omitempty"`

	// VPN网关的规格类型
	Flavor *string `json:"flavor,omitempty"`

	// 最大可创建的VPN连接数
	ConnectionNumber *int32 `json:"connection_number,omitempty"`

	// 当前已经使用的VPN连接数
	UsedConnectionNumber *int32 `json:"used_connection_number,omitempty"`

	// 当前已经使用的VPN连接组个数
	UsedConnectionGroup *int32 `json:"used_connection_group,omitempty"`

	// 企业项目ID
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o CreateResponseVpnGateway) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateResponseVpnGateway struct{}"
	}

	return strings.Join([]string{"CreateResponseVpnGateway", string(data)}, " ")
}

type CreateResponseVpnGatewayAttachmentType struct {
	value string
}

type CreateResponseVpnGatewayAttachmentTypeEnum struct {
	VPC CreateResponseVpnGatewayAttachmentType
	ER  CreateResponseVpnGatewayAttachmentType
}

func GetCreateResponseVpnGatewayAttachmentTypeEnum() CreateResponseVpnGatewayAttachmentTypeEnum {
	return CreateResponseVpnGatewayAttachmentTypeEnum{
		VPC: CreateResponseVpnGatewayAttachmentType{
			value: "vpc",
		},
		ER: CreateResponseVpnGatewayAttachmentType{
			value: "er",
		},
	}
}

func (c CreateResponseVpnGatewayAttachmentType) Value() string {
	return c.value
}

func (c CreateResponseVpnGatewayAttachmentType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateResponseVpnGatewayAttachmentType) UnmarshalJSON(b []byte) error {
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

type CreateResponseVpnGatewayNetworkType struct {
	value string
}

type CreateResponseVpnGatewayNetworkTypeEnum struct {
	PUBLIC  CreateResponseVpnGatewayNetworkType
	PRIVATE CreateResponseVpnGatewayNetworkType
}

func GetCreateResponseVpnGatewayNetworkTypeEnum() CreateResponseVpnGatewayNetworkTypeEnum {
	return CreateResponseVpnGatewayNetworkTypeEnum{
		PUBLIC: CreateResponseVpnGatewayNetworkType{
			value: "public",
		},
		PRIVATE: CreateResponseVpnGatewayNetworkType{
			value: "private",
		},
	}
}

func (c CreateResponseVpnGatewayNetworkType) Value() string {
	return c.value
}

func (c CreateResponseVpnGatewayNetworkType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateResponseVpnGatewayNetworkType) UnmarshalJSON(b []byte) error {
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
