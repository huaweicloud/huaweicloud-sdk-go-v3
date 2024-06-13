package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// PublicipUpdateResp 公网IP字段信息
type PublicipUpdateResp struct {

	// 功能说明：弹性公网IP唯一标识
	Id *string `json:"id,omitempty"`

	// 功能说明：项目ID
	ProjectId *string `json:"project_id,omitempty"`

	// 功能说明：IP版本信息 取值范围：4表示公网IP地址为public_ip_address地址;6表示公网IP地址为public_ipv6_address地址\"
	IpVersion *PublicipUpdateRespIpVersion `json:"ip_version,omitempty"`

	// 功能说明：弹性公网IP或者IPv6端口的地址
	PublicIpAddress *string `json:"public_ip_address,omitempty"`

	// 功能说明：IPv4时无此字段，IPv6时为申请到的弹性公网IP地址
	PublicIpv6Address *string `json:"public_ipv6_address,omitempty"`

	// 功能说明：弹性公网IP的状态  取值范围：冻结FREEZED，绑定失败BIND_ERROR，绑定中BINDING，释放中PENDING_DELETE， 创建中PENDING_CREATE，创建中NOTIFYING，释放中NOTIFY_DELETE，更新中PENDING_UPDATE， 未绑定DOWN ，绑定ACTIVE，绑定ELB，绑定VPN，失败ERROR。
	Status *PublicipUpdateRespStatus `json:"status,omitempty"`

	// 功能说明：弹性公网IP描述信息 约束：用户以自定义方式标识资源，系统不感知
	Description *string `json:"description,omitempty"`

	// 功能说明：表示中心站点资源或者边缘站点资源 取值范围： center、边缘站点名称 约束：publicip只能绑定该字段相同的资源
	PublicBorderGroup *string `json:"public_border_group,omitempty"`

	// 功能说明：资源创建UTC时间 格式:yyyy-MM-ddTHH:mm:ssZ
	CreatedAt *sdktime.SdkTime `json:"created_at,omitempty"`

	// 功能说明：资源更新UTC时间 格式:yyyy-MM-ddTHH:mm:ssZ
	UpdatedAt *sdktime.SdkTime `json:"updated_at,omitempty"`

	// 功能说明：弹性公网IP类型
	Type *PublicipUpdateRespType `json:"type,omitempty"`

	Vnic *VnicInfo `json:"vnic,omitempty"`

	Bandwidth *PublicipBandwidthInfo `json:"bandwidth,omitempty"`

	// 功能说明：企业项目ID。最大长度36字节,带“-”连字符的UUID格式,或者是字符串“0”。创建弹性公网IP时,给弹性公网IP绑定企业项目ID。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 功能说明：公网IP的订单信息 约束：包周期才会有订单信息，按需资源此字段为空
	BillingInfo *string `json:"billing_info,omitempty"`

	// 功能说明：记录公网IP当前的冻结状态 约束：metadata类型，标识欠费冻结、公安冻结 取值范围：police，locked
	LockStatus *string `json:"lock_status,omitempty"`

	// 功能说明：公网IP绑定的实例类型 取值范围：PORT、NATGW、ELB、ELBV1、VPN、null
	AssociateInstanceType *PublicipUpdateRespAssociateInstanceType `json:"associate_instance_type,omitempty"`

	// 功能说明：公网IP绑定的实例ID
	AssociateInstanceId *string `json:"associate_instance_id,omitempty"`

	// 功能说明：公网IP所属网络的ID。publicip_pool_name对应的网络ID
	PublicipPoolId *string `json:"publicip_pool_id,omitempty"`

	// 功能说明：弹性公网IP的网络类型, 包括公共池类型，如5_bgp/5_sbgp...，和用户购买的专属池。 专属池见publcip_pool相关接口
	PublicipPoolName *string `json:"publicip_pool_name,omitempty"`

	// 功能说明：弹性公网IP名称
	Alias *string `json:"alias,omitempty"`

	// 默认不显示。开启支持直通模式后展示，表示直通模式的标识。
	AssociateMode *string `json:"associate_mode,omitempty"`
}

func (o PublicipUpdateResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PublicipUpdateResp struct{}"
	}

	return strings.Join([]string{"PublicipUpdateResp", string(data)}, " ")
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

func (c PublicipUpdateRespIpVersion) Value() int32 {
	return c.value
}

func (c PublicipUpdateRespIpVersion) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PublicipUpdateRespIpVersion) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
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

func (c PublicipUpdateRespStatus) Value() string {
	return c.value
}

func (c PublicipUpdateRespStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PublicipUpdateRespStatus) UnmarshalJSON(b []byte) error {
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

type PublicipUpdateRespType struct {
	value string
}

type PublicipUpdateRespTypeEnum struct {
	EIP              PublicipUpdateRespType
	DUALSTACK        PublicipUpdateRespType
	DUALSTACK_SUBNET PublicipUpdateRespType
}

func GetPublicipUpdateRespTypeEnum() PublicipUpdateRespTypeEnum {
	return PublicipUpdateRespTypeEnum{
		EIP: PublicipUpdateRespType{
			value: "EIP",
		},
		DUALSTACK: PublicipUpdateRespType{
			value: "DUALSTACK",
		},
		DUALSTACK_SUBNET: PublicipUpdateRespType{
			value: "DUALSTACK_SUBNET",
		},
	}
}

func (c PublicipUpdateRespType) Value() string {
	return c.value
}

func (c PublicipUpdateRespType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PublicipUpdateRespType) UnmarshalJSON(b []byte) error {
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

type PublicipUpdateRespAssociateInstanceType struct {
	value string
}

type PublicipUpdateRespAssociateInstanceTypeEnum struct {
	PORT  PublicipUpdateRespAssociateInstanceType
	NATGW PublicipUpdateRespAssociateInstanceType
	ELB   PublicipUpdateRespAssociateInstanceType
	ELBV1 PublicipUpdateRespAssociateInstanceType
	VPN   PublicipUpdateRespAssociateInstanceType
}

func GetPublicipUpdateRespAssociateInstanceTypeEnum() PublicipUpdateRespAssociateInstanceTypeEnum {
	return PublicipUpdateRespAssociateInstanceTypeEnum{
		PORT: PublicipUpdateRespAssociateInstanceType{
			value: "PORT",
		},
		NATGW: PublicipUpdateRespAssociateInstanceType{
			value: "NATGW",
		},
		ELB: PublicipUpdateRespAssociateInstanceType{
			value: "ELB",
		},
		ELBV1: PublicipUpdateRespAssociateInstanceType{
			value: "ELBV1",
		},
		VPN: PublicipUpdateRespAssociateInstanceType{
			value: "VPN",
		},
	}
}

func (c PublicipUpdateRespAssociateInstanceType) Value() string {
	return c.value
}

func (c PublicipUpdateRespAssociateInstanceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PublicipUpdateRespAssociateInstanceType) UnmarshalJSON(b []byte) error {
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
