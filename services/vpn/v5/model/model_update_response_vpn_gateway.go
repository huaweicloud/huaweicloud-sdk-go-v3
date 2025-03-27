package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

type UpdateResponseVpnGateway struct {

	// VPN网关ID
	Id *string `json:"id,omitempty"`

	// VPN网关名称
	Name *string `json:"name,omitempty"`

	// 关联模式
	AttachmentType *UpdateResponseVpnGatewayAttachmentType `json:"attachment_type,omitempty"`

	// 网关的IP协议版本
	IpVersion *UpdateResponseVpnGatewayIpVersion `json:"ip_version,omitempty"`

	CertificateId *string `json:"certificate_id,omitempty"`

	// VPN网关所连接的ER实例的ID
	ErId *string `json:"er_id,omitempty"`

	// VPN网关所连接的VPC的ID
	VpcId *string `json:"vpc_id,omitempty"`

	// 本端子网
	LocalSubnets *[]string `json:"local_subnets,omitempty"`

	// 使能ipv6的本端子网
	LocalSubnetsV6 *[]string `json:"local_subnets_v6,omitempty"`

	// VPN网关所使用的VPC子网ID
	ConnectSubnet *string `json:"connect_subnet,omitempty"`

	// VPN网关的网络类型，默认为公网(public)
	NetworkType *UpdateResponseVpnGatewayNetworkType `json:"network_type,omitempty"`

	// VPN网关北向接入VPC ID，不填时默认使用vpc_id字段的值
	AccessVpcId *string `json:"access_vpc_id,omitempty"`

	// VPN网关北向接入VPC中的接入子网ID
	AccessSubnetId *string `json:"access_subnet_id,omitempty"`

	// 私网类型VPN网关的接入私网IP，VPN网关使用该私网IP与对端网关建连。双活网关表示使用的第一个私网地址，主备表示主私网地址。
	AccessPrivateIp1 *string `json:"access_private_ip_1,omitempty"`

	// 私网类型VPN网关的接入私网IP，VPN网关使用该私网IP与对端网关建连。双活网关表示使用的第二个私网地址，主备表示备私网地址。
	AccessPrivateIp2 *string `json:"access_private_ip_2,omitempty"`

	// bgp所使用的asn号
	BgpAsn *int64 `json:"bgp_asn,omitempty"`

	// VPN网关的规格类型
	Flavor *string `json:"flavor,omitempty"`

	// 可用区列表
	AvailabilityZoneIds *[]string `json:"availability_zone_ids,omitempty"`

	// 最大可创建的VPN连接数
	ConnectionNumber *int32 `json:"connection_number,omitempty"`

	// 当前已经使用的VPN连接数
	UsedConnectionNumber *int32 `json:"used_connection_number,omitempty"`

	// 当前已经使用的VPN连接组个数
	UsedConnectionGroup *int32 `json:"used_connection_group,omitempty"`

	// 企业项目ID
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// ha模式
	HaMode *string `json:"ha_mode,omitempty"`

	Eip1 *ResponseEip `json:"eip1,omitempty"`

	Eip2 *ResponseEip `json:"eip2,omitempty"`

	// 创建时间
	CreatedAt *sdktime.SdkTime `json:"created_at,omitempty"`

	// 更新时间
	UpdatedAt *sdktime.SdkTime `json:"updated_at,omitempty"`

	PolicyTemplate *PolicyTemplate `json:"policy_template,omitempty"`

	// 标签
	Tags *[]VpnResourceTag `json:"tags,omitempty"`
}

func (o UpdateResponseVpnGateway) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateResponseVpnGateway struct{}"
	}

	return strings.Join([]string{"UpdateResponseVpnGateway", string(data)}, " ")
}

type UpdateResponseVpnGatewayAttachmentType struct {
	value string
}

type UpdateResponseVpnGatewayAttachmentTypeEnum struct {
	VPC UpdateResponseVpnGatewayAttachmentType
	ER  UpdateResponseVpnGatewayAttachmentType
}

func GetUpdateResponseVpnGatewayAttachmentTypeEnum() UpdateResponseVpnGatewayAttachmentTypeEnum {
	return UpdateResponseVpnGatewayAttachmentTypeEnum{
		VPC: UpdateResponseVpnGatewayAttachmentType{
			value: "vpc",
		},
		ER: UpdateResponseVpnGatewayAttachmentType{
			value: "er",
		},
	}
}

func (c UpdateResponseVpnGatewayAttachmentType) Value() string {
	return c.value
}

func (c UpdateResponseVpnGatewayAttachmentType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateResponseVpnGatewayAttachmentType) UnmarshalJSON(b []byte) error {
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

type UpdateResponseVpnGatewayIpVersion struct {
	value string
}

type UpdateResponseVpnGatewayIpVersionEnum struct {
	IPV4 UpdateResponseVpnGatewayIpVersion
	IPV6 UpdateResponseVpnGatewayIpVersion
}

func GetUpdateResponseVpnGatewayIpVersionEnum() UpdateResponseVpnGatewayIpVersionEnum {
	return UpdateResponseVpnGatewayIpVersionEnum{
		IPV4: UpdateResponseVpnGatewayIpVersion{
			value: "ipv4",
		},
		IPV6: UpdateResponseVpnGatewayIpVersion{
			value: "ipv6",
		},
	}
}

func (c UpdateResponseVpnGatewayIpVersion) Value() string {
	return c.value
}

func (c UpdateResponseVpnGatewayIpVersion) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateResponseVpnGatewayIpVersion) UnmarshalJSON(b []byte) error {
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

type UpdateResponseVpnGatewayNetworkType struct {
	value string
}

type UpdateResponseVpnGatewayNetworkTypeEnum struct {
	PUBLIC  UpdateResponseVpnGatewayNetworkType
	PRIVATE UpdateResponseVpnGatewayNetworkType
}

func GetUpdateResponseVpnGatewayNetworkTypeEnum() UpdateResponseVpnGatewayNetworkTypeEnum {
	return UpdateResponseVpnGatewayNetworkTypeEnum{
		PUBLIC: UpdateResponseVpnGatewayNetworkType{
			value: "public",
		},
		PRIVATE: UpdateResponseVpnGatewayNetworkType{
			value: "private",
		},
	}
}

func (c UpdateResponseVpnGatewayNetworkType) Value() string {
	return c.value
}

func (c UpdateResponseVpnGatewayNetworkType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateResponseVpnGatewayNetworkType) UnmarshalJSON(b []byte) error {
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
