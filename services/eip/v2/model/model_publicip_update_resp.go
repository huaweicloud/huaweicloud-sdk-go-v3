package model

import (
	"errors"
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/converter"
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// 弹性公网IP列表返回体
type PublicipUpdateResp struct {
	// 弹性公网IP对应带宽ID

	BandwidthId *string `json:"bandwidth_id,omitempty"`
	// 带宽名称

	BandwidthName *string `json:"bandwidth_name,omitempty"`
	// 表示共享带宽或者独享带宽  取值范围：PER，WHOLE。  WHOLE表示共享带宽  PER表示独享带宽  约束：其中IPv6暂不支持WHOLE类型带宽。

	BandwidthShareType *PublicipUpdateRespBandwidthShareType `json:"bandwidth_share_type,omitempty"`
	// 带宽大小，单位为Mbit/s。

	BandwidthSize *int32 `json:"bandwidth_size,omitempty"`
	// 弹性公网IP申请时间（UTC）

	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`
	// 企业项目ID。最大长度36字节，带“-”连字符的UUID格式，或者是字符串“0”。  创建弹性公网IP时，给弹性公网IP绑定企业项目ID。

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
	// 弹性公网IP唯一标识

	Id *string `json:"id,omitempty"`
	// 功能说明：端口id。  约束：只有绑定了的弹性公网IP查询才会返回该参数

	PortId *string `json:"port_id,omitempty"`
	// 功能说明：绑定弹性公网IP的私有IP地址  约束：只有绑定了的弹性公网IP查询才会返回该参数

	PrivateIpAddress *string `json:"private_ip_address,omitempty"`

	Profile *ProfileResp `json:"profile,omitempty"`
	// IPv4时是申请到的弹性公网IP地址，IPv6时是IPv6地址对应的IPv4地址

	PublicIpAddress *string `json:"public_ip_address,omitempty"`
	// 功能说明：弹性公网IP的状态  取值范围：冻结FREEZED，绑定失败BIND_ERROR，绑定中BINDING，释放中PENDING_DELETE， 创建中PENDING_CREATE，创建中NOTIFYING，释放中NOTIFY_DELETE，更新中PENDING_UPDATE， 未绑定DOWN ，绑定ACTIVE，绑定ELB，绑定VPN，失败ERROR。

	Status *PublicipUpdateRespStatus `json:"status,omitempty"`
	// 项目ID

	TenantId *string `json:"tenant_id,omitempty"`
	// 弹性公网IP的类型

	Type *string `json:"type,omitempty"`
	// IPv4时无此字段，IPv6时为申请到的弹性公网IP地址

	PublicIpv6Address *string `json:"public_ipv6_address,omitempty"`
	// IP版本信息，取值范围是4和6  4：表示IPv4  6：表示IPv6

	IpVersion *PublicipUpdateRespIpVersion `json:"ip_version,omitempty"`
	// 功能说明：弹性公网IP名称 取值范围：1-64个字符，支持数字、字母、中文、_(下划线)、-（中划线）、.（点）

	Alias *string `json:"alias,omitempty"`
}

func (o PublicipUpdateResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PublicipUpdateResp struct{}"
	}

	return strings.Join([]string{"PublicipUpdateResp", string(data)}, " ")
}

type PublicipUpdateRespBandwidthShareType struct {
	value string
}

type PublicipUpdateRespBandwidthShareTypeEnum struct {
	WHOLE PublicipUpdateRespBandwidthShareType
	PER   PublicipUpdateRespBandwidthShareType
}

func GetPublicipUpdateRespBandwidthShareTypeEnum() PublicipUpdateRespBandwidthShareTypeEnum {
	return PublicipUpdateRespBandwidthShareTypeEnum{
		WHOLE: PublicipUpdateRespBandwidthShareType{
			value: "WHOLE",
		},
		PER: PublicipUpdateRespBandwidthShareType{
			value: "PER",
		},
	}
}

func (c PublicipUpdateRespBandwidthShareType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PublicipUpdateRespBandwidthShareType) UnmarshalJSON(b []byte) error {
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

type PublicipUpdateRespStatus struct {
	value string
}

type PublicipUpdateRespStatusEnum struct {
	FREEZED        PublicipUpdateRespStatus
	BIND_ERROR     PublicipUpdateRespStatus
	BINDING        PublicipUpdateRespStatus
	PENDING_DELETE PublicipUpdateRespStatus
	PENDING_CREATE PublicipUpdateRespStatus
	NOTIFYING      PublicipUpdateRespStatus
	NOTIFY_DELETE  PublicipUpdateRespStatus
	PENDING_UPDATE PublicipUpdateRespStatus
	DOWN           PublicipUpdateRespStatus
	ACTIVE         PublicipUpdateRespStatus
	ELB            PublicipUpdateRespStatus
	ERROR          PublicipUpdateRespStatus
	VPN            PublicipUpdateRespStatus
}

func GetPublicipUpdateRespStatusEnum() PublicipUpdateRespStatusEnum {
	return PublicipUpdateRespStatusEnum{
		FREEZED: PublicipUpdateRespStatus{
			value: "FREEZED",
		},
		BIND_ERROR: PublicipUpdateRespStatus{
			value: "BIND_ERROR",
		},
		BINDING: PublicipUpdateRespStatus{
			value: "BINDING",
		},
		PENDING_DELETE: PublicipUpdateRespStatus{
			value: "PENDING_DELETE",
		},
		PENDING_CREATE: PublicipUpdateRespStatus{
			value: "PENDING_CREATE",
		},
		NOTIFYING: PublicipUpdateRespStatus{
			value: "NOTIFYING",
		},
		NOTIFY_DELETE: PublicipUpdateRespStatus{
			value: "NOTIFY_DELETE",
		},
		PENDING_UPDATE: PublicipUpdateRespStatus{
			value: "PENDING_UPDATE",
		},
		DOWN: PublicipUpdateRespStatus{
			value: "DOWN",
		},
		ACTIVE: PublicipUpdateRespStatus{
			value: "ACTIVE",
		},
		ELB: PublicipUpdateRespStatus{
			value: "ELB",
		},
		ERROR: PublicipUpdateRespStatus{
			value: "ERROR",
		},
		VPN: PublicipUpdateRespStatus{
			value: "VPN",
		},
	}
}

func (c PublicipUpdateRespStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PublicipUpdateRespStatus) UnmarshalJSON(b []byte) error {
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

type PublicipUpdateRespIpVersion struct {
	value int32
}

type PublicipUpdateRespIpVersionEnum struct {
	E_4 PublicipUpdateRespIpVersion
	E_6 PublicipUpdateRespIpVersion
}

func GetPublicipUpdateRespIpVersionEnum() PublicipUpdateRespIpVersionEnum {
	return PublicipUpdateRespIpVersionEnum{
		E_4: PublicipUpdateRespIpVersion{
			value: 4,
		}, E_6: PublicipUpdateRespIpVersion{
			value: 6,
		},
	}
}

func (c PublicipUpdateRespIpVersion) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PublicipUpdateRespIpVersion) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(int32)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
