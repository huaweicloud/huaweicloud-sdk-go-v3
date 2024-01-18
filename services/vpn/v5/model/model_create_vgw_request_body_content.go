package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CreateVgwRequestBodyContent struct {

	// VPN网关名称
	Name *string `json:"name,omitempty"`

	// 关联模式
	AttachmentType *CreateVgwRequestBodyContentAttachmentType `json:"attachment_type,omitempty"`

	// VPN网关所连接的ER实例的ID，当attachment_type配置为\"er\"时必填，否则不填
	ErId *string `json:"er_id,omitempty"`

	// VPN网关所连接的VPC的ID
	VpcId *string `json:"vpc_id,omitempty"`

	// 本端子网，当attachment_type配置为\"vpc\"时必填，否则不填
	LocalSubnets *[]string `json:"local_subnets,omitempty"`

	// VPN网关所使用的VPC子网ID
	ConnectSubnet *string `json:"connect_subnet,omitempty"`

	// bgp所使用的asn号
	BgpAsn *int64 `json:"bgp_asn,omitempty"`

	// VPN网关的规格类型，当attachment_type为er时不能填写Basic
	Flavor *CreateVgwRequestBodyContentFlavor `json:"flavor,omitempty"`

	// 不填写则采用默认可用区。如果需要指定可用区可以调用查询VPN网关可用区接口来选择
	AvailabilityZoneIds *[]string `json:"availability_zone_ids,omitempty"`

	// 企业项目ID
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	Eip1 *CreateRequestEip `json:"eip1,omitempty"`

	Eip2 *CreateRequestEip `json:"eip2,omitempty"`

	// VPN网关北向接入私网IP，当VPN网关的北向类型是私网(private)时有值,主备模式代表主worker的私网IP
	AccessPrivateIp1 *string `json:"access_private_ip_1,omitempty"`

	// VPN网关北向接入私网IP，当VPN网关的北向类型是私网(private)时有值,主备模式代表备worker的私网IP
	AccessPrivateIp2 *string `json:"access_private_ip_2,omitempty"`

	// VPN网关北向类型，默认为公网(public)
	NetworkType *CreateVgwRequestBodyContentNetworkType `json:"network_type,omitempty"`

	// VPN网关北向接入VPC ID，不填时默认使用vpc_id字段的值
	AccessVpcId *string `json:"access_vpc_id,omitempty"`

	// VPN网关北向接入VPC中的接入子网ID，在填写了access_vpc_id时必填
	AccessSubnetId *string `json:"access_subnet_id,omitempty"`

	// ha模式
	HaMode *CreateVgwRequestBodyContentHaMode `json:"ha_mode,omitempty"`

	// 标签
	Tags *[]VpnResourceTag `json:"tags,omitempty"`
}

func (o CreateVgwRequestBodyContent) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVgwRequestBodyContent struct{}"
	}

	return strings.Join([]string{"CreateVgwRequestBodyContent", string(data)}, " ")
}

type CreateVgwRequestBodyContentAttachmentType struct {
	value string
}

type CreateVgwRequestBodyContentAttachmentTypeEnum struct {
	VPC CreateVgwRequestBodyContentAttachmentType
	ER  CreateVgwRequestBodyContentAttachmentType
}

func GetCreateVgwRequestBodyContentAttachmentTypeEnum() CreateVgwRequestBodyContentAttachmentTypeEnum {
	return CreateVgwRequestBodyContentAttachmentTypeEnum{
		VPC: CreateVgwRequestBodyContentAttachmentType{
			value: "vpc",
		},
		ER: CreateVgwRequestBodyContentAttachmentType{
			value: "er",
		},
	}
}

func (c CreateVgwRequestBodyContentAttachmentType) Value() string {
	return c.value
}

func (c CreateVgwRequestBodyContentAttachmentType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateVgwRequestBodyContentAttachmentType) UnmarshalJSON(b []byte) error {
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

type CreateVgwRequestBodyContentFlavor struct {
	value string
}

type CreateVgwRequestBodyContentFlavorEnum struct {
	BASIC                      CreateVgwRequestBodyContentFlavor
	PROFESSIONAL1              CreateVgwRequestBodyContentFlavor
	PROFESSIONAL2              CreateVgwRequestBodyContentFlavor
	PROFESSIONAL1_NON_FIXED_IP CreateVgwRequestBodyContentFlavor
	PROFESSIONAL2_NON_FIXED_IP CreateVgwRequestBodyContentFlavor
}

func GetCreateVgwRequestBodyContentFlavorEnum() CreateVgwRequestBodyContentFlavorEnum {
	return CreateVgwRequestBodyContentFlavorEnum{
		BASIC: CreateVgwRequestBodyContentFlavor{
			value: "Basic",
		},
		PROFESSIONAL1: CreateVgwRequestBodyContentFlavor{
			value: "Professional1",
		},
		PROFESSIONAL2: CreateVgwRequestBodyContentFlavor{
			value: "Professional2",
		},
		PROFESSIONAL1_NON_FIXED_IP: CreateVgwRequestBodyContentFlavor{
			value: "Professional1-NonFixedIP",
		},
		PROFESSIONAL2_NON_FIXED_IP: CreateVgwRequestBodyContentFlavor{
			value: "Professional2-NonFixedIP",
		},
	}
}

func (c CreateVgwRequestBodyContentFlavor) Value() string {
	return c.value
}

func (c CreateVgwRequestBodyContentFlavor) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateVgwRequestBodyContentFlavor) UnmarshalJSON(b []byte) error {
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

type CreateVgwRequestBodyContentNetworkType struct {
	value string
}

type CreateVgwRequestBodyContentNetworkTypeEnum struct {
	PUBLIC  CreateVgwRequestBodyContentNetworkType
	PRIVATE CreateVgwRequestBodyContentNetworkType
}

func GetCreateVgwRequestBodyContentNetworkTypeEnum() CreateVgwRequestBodyContentNetworkTypeEnum {
	return CreateVgwRequestBodyContentNetworkTypeEnum{
		PUBLIC: CreateVgwRequestBodyContentNetworkType{
			value: "public",
		},
		PRIVATE: CreateVgwRequestBodyContentNetworkType{
			value: "private",
		},
	}
}

func (c CreateVgwRequestBodyContentNetworkType) Value() string {
	return c.value
}

func (c CreateVgwRequestBodyContentNetworkType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateVgwRequestBodyContentNetworkType) UnmarshalJSON(b []byte) error {
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

type CreateVgwRequestBodyContentHaMode struct {
	value string
}

type CreateVgwRequestBodyContentHaModeEnum struct {
	ACTIVE_ACTIVE  CreateVgwRequestBodyContentHaMode
	ACTIVE_STANDBY CreateVgwRequestBodyContentHaMode
}

func GetCreateVgwRequestBodyContentHaModeEnum() CreateVgwRequestBodyContentHaModeEnum {
	return CreateVgwRequestBodyContentHaModeEnum{
		ACTIVE_ACTIVE: CreateVgwRequestBodyContentHaMode{
			value: "active-active",
		},
		ACTIVE_STANDBY: CreateVgwRequestBodyContentHaMode{
			value: "active-standby",
		},
	}
}

func (c CreateVgwRequestBodyContentHaMode) Value() string {
	return c.value
}

func (c CreateVgwRequestBodyContentHaMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateVgwRequestBodyContentHaMode) UnmarshalJSON(b []byte) error {
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
