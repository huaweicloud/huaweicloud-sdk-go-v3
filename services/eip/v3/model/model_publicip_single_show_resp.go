package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/sdktime"
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"
	"errors"
	"strings"
)

// 公网IP字段信息
type PublicipSingleShowResp struct {
	// 功能说明：弹性公网IP唯一标识

	Id *string `json:"id,omitempty"`
	// 功能说明：项目ID

	ProjectId *string `json:"project_id,omitempty"`
	// 功能说明：IP版本信息 取值范围：4表示公网IP地址为public_ip_address地址;6表示公网IP地址为public_ipv6_address地址\"

	IpVersion *PublicipSingleShowRespIpVersion `json:"ip_version,omitempty"`
	// 功能说明：弹性公网IP或者IPv6端口的地址

	PublicIpAddress *string `json:"public_ip_address,omitempty"`
	// 功能说明：IPv4时无此字段，IPv6时为申请到的弹性公网IP地址

	PublicIpv6Address *string `json:"public_ipv6_address,omitempty"`
	// 废弃，功能由publicip_pool_name继承，默认不显示。功能说明：弹性公网IP的网络类型

	NetworkType *string `json:"network_type,omitempty"`
	// 功能说明：弹性公网IP的状态  取值范围：冻结FREEZED，绑定失败BIND_ERROR，绑定中BINDING，释放中PENDING_DELETE， 创建中PENDING_CREATE，创建中NOTIFYING，释放中NOTIFY_DELETE，更新中PENDING_UPDATE， 未绑定DOWN ，绑定ACTIVE，绑定ELB，绑定VPN，失败ERROR。

	Status *PublicipSingleShowRespStatus `json:"status,omitempty"`
	// 功能说明：弹性公网IP描述信息 约束：用户以自定义方式标识资源，系统不感知

	Description *string `json:"description,omitempty"`
	// 功能说明：表示中心站点资源或者边缘站点资源 取值范围： center、边缘站点名称 约束：publicip只能绑定该字段相同的资源

	PublicBorderGroup *string `json:"public_border_group,omitempty"`
	// 功能说明：资源创建UTC时间 格式:yyyy-MM-ddTHH:mm:ssZ

	CreatedAt *sdktime.SdkTime `json:"created_at,omitempty"`
	// 功能说明：资源更新UTC时间 格式:yyyy-MM-ddTHH:mm:ssZ

	UpdatedAt *sdktime.SdkTime `json:"updated_at,omitempty"`
	// 功能说明：弹性公网IP类型

	Type *PublicipSingleShowRespType `json:"type,omitempty"`

	Vnic *VnicInfo `json:"vnic,omitempty"`

	Bandwidth *PublicipBandwidthInfo `json:"bandwidth,omitempty"`
	// 功能说明：企业项目ID。最大长度36字节,带“-”连字符的UUID格式,或者是字符串“0”。创建弹性公网IP时,给弹性公网IP绑定企业项目ID。

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
	// 功能说明：公网IP的订单信息 约束：包周期才会有订单信息，按需资源此字段为空

	BillingInfo *string `json:"billing_info,omitempty"`
	// 功能说明：记录公网IP当前的冻结状态 约束：metadata类型，标识欠费冻结、公安冻结 取值范围：police，locked

	LockStatus *string `json:"lock_status,omitempty"`
	// 功能说明：公网IP绑定的实例类型 取值范围：PORT、NATGW、ELB、ELBV1、VPN、null

	AssociateInstanceType *PublicipSingleShowRespAssociateInstanceType `json:"associate_instance_type,omitempty"`
	// 功能说明：公网IP绑定的实例ID

	AssociateInstanceId *string `json:"associate_instance_id,omitempty"`
	// 功能说明：公网IP所属网络的ID。publicip_pool_name对应的网络ID

	PublicipPoolId *string `json:"publicip_pool_id,omitempty"`
	// 功能说明：弹性公网IP的网络类型, 包括公共池类型，如5_bgp/5_sbgp...，和用户购买的专属池。 专属池见publcip_pool相关接口

	PublicipPoolName *string `json:"publicip_pool_name,omitempty"`
	// 功能说明：弹性公网IP名称

	Alias *string `json:"alias,omitempty"`

	Profile *ProfileInfo `json:"profile,omitempty"`
	// 默认不显示。该字段仅仅用于表示eip的bgp类型是否是真实的静态sbgp * 1. 如果为true，则该eip可以切换bgp类型 * 2. 如果为false，则该eip不可以切换bgp类型

	FakeNetworkType *bool `json:"fake_network_type,omitempty"`
	// 默认不显示。用户标签

	Tags *[]TagsInfo `json:"tags,omitempty"`
	// 默认不显示。记录实例的更上一层归属。例如associate_instance_type为PORT，此字段记录PORT的device_id和device_owner信息。仅有限场景记录。

	AssociateInstanceMetadata *string `json:"associate_instance_metadata,omitempty"`
	// 默认不显示。开启支持直通模式后展示，表示直通模式的标识。

	AssociateMode *string `json:"associate_mode,omitempty"`
	// 功能说明：表示此publicip可以加入的共享带宽类型列表，如果为空列表，则表示该           publicip不能加入任何共享带宽 约束：publicip只能加入到有该带宽类型的共享带宽中

	AllowShareBandwidthTypes *[]string `json:"allow_share_bandwidth_types,omitempty"`
	// 默认不显示。表示该eip是否支持与实例同步删除。

	CascadeDeleteByInstance *bool `json:"cascade_delete_by_instance,omitempty"`
}

func (o PublicipSingleShowResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PublicipSingleShowResp struct{}"
	}

	return strings.Join([]string{"PublicipSingleShowResp", string(data)}, " ")
}

type PublicipSingleShowRespIpVersion struct {
	value int32
}

type PublicipSingleShowRespIpVersionEnum struct {
	E_4 PublicipSingleShowRespIpVersion
	E_6 PublicipSingleShowRespIpVersion
}

func GetPublicipSingleShowRespIpVersionEnum() PublicipSingleShowRespIpVersionEnum {
	return PublicipSingleShowRespIpVersionEnum{
		E_4: PublicipSingleShowRespIpVersion{
			value: 4,
		}, E_6: PublicipSingleShowRespIpVersion{
			value: 6,
		},
	}
}

func (c PublicipSingleShowRespIpVersion) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PublicipSingleShowRespIpVersion) UnmarshalJSON(b []byte) error {
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

type PublicipSingleShowRespStatus struct {
	value string
}

type PublicipSingleShowRespStatusEnum struct {
	FREEZED        PublicipSingleShowRespStatus
	BIND_ERROR     PublicipSingleShowRespStatus
	BINDING        PublicipSingleShowRespStatus
	PENDING_DELETE PublicipSingleShowRespStatus
	PENDING_CREATE PublicipSingleShowRespStatus
	NOTIFYING      PublicipSingleShowRespStatus
	NOTIFY_DELETE  PublicipSingleShowRespStatus
	PENDING_UPDATE PublicipSingleShowRespStatus
	DOWN           PublicipSingleShowRespStatus
	ACTIVE         PublicipSingleShowRespStatus
	ELB            PublicipSingleShowRespStatus
	ERROR          PublicipSingleShowRespStatus
	VPN            PublicipSingleShowRespStatus
}

func GetPublicipSingleShowRespStatusEnum() PublicipSingleShowRespStatusEnum {
	return PublicipSingleShowRespStatusEnum{
		FREEZED: PublicipSingleShowRespStatus{
			value: "FREEZED",
		},
		BIND_ERROR: PublicipSingleShowRespStatus{
			value: "BIND_ERROR",
		},
		BINDING: PublicipSingleShowRespStatus{
			value: "BINDING",
		},
		PENDING_DELETE: PublicipSingleShowRespStatus{
			value: "PENDING_DELETE",
		},
		PENDING_CREATE: PublicipSingleShowRespStatus{
			value: "PENDING_CREATE",
		},
		NOTIFYING: PublicipSingleShowRespStatus{
			value: "NOTIFYING",
		},
		NOTIFY_DELETE: PublicipSingleShowRespStatus{
			value: "NOTIFY_DELETE",
		},
		PENDING_UPDATE: PublicipSingleShowRespStatus{
			value: "PENDING_UPDATE",
		},
		DOWN: PublicipSingleShowRespStatus{
			value: "DOWN",
		},
		ACTIVE: PublicipSingleShowRespStatus{
			value: "ACTIVE",
		},
		ELB: PublicipSingleShowRespStatus{
			value: "ELB",
		},
		ERROR: PublicipSingleShowRespStatus{
			value: "ERROR",
		},
		VPN: PublicipSingleShowRespStatus{
			value: "VPN",
		},
	}
}

func (c PublicipSingleShowRespStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PublicipSingleShowRespStatus) UnmarshalJSON(b []byte) error {
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

type PublicipSingleShowRespType struct {
	value string
}

type PublicipSingleShowRespTypeEnum struct {
	EIP              PublicipSingleShowRespType
	DUALSTACK        PublicipSingleShowRespType
	DUALSTACK_SUBNET PublicipSingleShowRespType
}

func GetPublicipSingleShowRespTypeEnum() PublicipSingleShowRespTypeEnum {
	return PublicipSingleShowRespTypeEnum{
		EIP: PublicipSingleShowRespType{
			value: "EIP",
		},
		DUALSTACK: PublicipSingleShowRespType{
			value: "DUALSTACK",
		},
		DUALSTACK_SUBNET: PublicipSingleShowRespType{
			value: "DUALSTACK_SUBNET",
		},
	}
}

func (c PublicipSingleShowRespType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PublicipSingleShowRespType) UnmarshalJSON(b []byte) error {
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

type PublicipSingleShowRespAssociateInstanceType struct {
	value string
}

type PublicipSingleShowRespAssociateInstanceTypeEnum struct {
	PORT  PublicipSingleShowRespAssociateInstanceType
	NATGW PublicipSingleShowRespAssociateInstanceType
	ELB   PublicipSingleShowRespAssociateInstanceType
	ELBV1 PublicipSingleShowRespAssociateInstanceType
	VPN   PublicipSingleShowRespAssociateInstanceType
	NULL  PublicipSingleShowRespAssociateInstanceType
}

func GetPublicipSingleShowRespAssociateInstanceTypeEnum() PublicipSingleShowRespAssociateInstanceTypeEnum {
	return PublicipSingleShowRespAssociateInstanceTypeEnum{
		PORT: PublicipSingleShowRespAssociateInstanceType{
			value: "PORT",
		},
		NATGW: PublicipSingleShowRespAssociateInstanceType{
			value: "NATGW",
		},
		ELB: PublicipSingleShowRespAssociateInstanceType{
			value: "ELB",
		},
		ELBV1: PublicipSingleShowRespAssociateInstanceType{
			value: "ELBV1",
		},
		VPN: PublicipSingleShowRespAssociateInstanceType{
			value: "VPN",
		},
		NULL: PublicipSingleShowRespAssociateInstanceType{
			value: "null",
		},
	}
}

func (c PublicipSingleShowRespAssociateInstanceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PublicipSingleShowRespAssociateInstanceType) UnmarshalJSON(b []byte) error {
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
