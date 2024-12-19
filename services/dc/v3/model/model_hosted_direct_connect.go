package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// HostedDirectConnect 托管专线对象信息
type HostedDirectConnect struct {

	// 托管专线ID
	Id *string `json:"id,omitempty"`

	// 实例所属项目ID。
	TenantId *string `json:"tenant_id,omitempty"`

	// 托管专线名字
	Name *string `json:"name,omitempty"`

	// 托管专线的描述信息
	Description *string `json:"description,omitempty"`

	// 托管专线接入带宽，单位Mbps。
	Bandwidth *int32 `json:"bandwidth,omitempty"`

	// 专线的接入位置信息
	Location *string `json:"location,omitempty"`

	// 托管专线对端所在的物理位置，省/市/街道或IDC名字。
	PeerLocation *string `json:"peer_location,omitempty"`

	// hosted物理专线对应的hosting物理专线的ID
	HostingId *string `json:"hosting_id,omitempty"`

	// 专线线路的提供商
	Provider *string `json:"provider,omitempty"`

	// 管理状态：true或false
	AdminStateUp *bool `json:"admin_state_up,omitempty"`

	// hosted物理专线预分配的vlan。
	Vlan *int32 `json:"vlan,omitempty"`

	// 操作状态，合法值是： BUILD：已开通 ACTIVE：托管专线正常 DOWN：专线对应的端口处于down的状态，可能存在线路故障等异常。 ERROR：托管专线异常 PENDING_DELETE：删除中 PENDING_UPDATE：更新中 PENDING_CREATE：创建中
	Status *HostedDirectConnectStatus `json:"status,omitempty"`

	// 托管专线申请时间。采用UTC时间格式，格式为：yyyy-MM-ddTHH:mm:ss.SSSZ
	ApplyTime *sdktime.SdkTime `json:"apply_time,omitempty"`

	// 托管专线创建时间。采用UTC时间格式，格式为：yyyy-MM-ddTHH:mm:ss.SSSZ
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// 托管专线的运营商操作状态，合法值是：ACTIVE， DOWN
	ProviderStatus *HostedDirectConnectProviderStatus `json:"provider_status,omitempty"`

	// 托管专线接入接口的类型，支持1G 10G 40G 100G
	PortType *HostedDirectConnectPortType `json:"port_type,omitempty"`

	// 托管专线的类型，类型包括标准(standard)，运营专线(hosting)，托管专线（hosted）[，一站式标准（onestop_standard），一站式托管（onestop_hosted）](tag:hws)。
	Type *HostedDirectConnectType `json:"type,omitempty"`
}

func (o HostedDirectConnect) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HostedDirectConnect struct{}"
	}

	return strings.Join([]string{"HostedDirectConnect", string(data)}, " ")
}

type HostedDirectConnectStatus struct {
	value string
}

type HostedDirectConnectStatusEnum struct {
	BUILD          HostedDirectConnectStatus
	ACTIVE         HostedDirectConnectStatus
	DOWN           HostedDirectConnectStatus
	ERROR          HostedDirectConnectStatus
	PENDING_DELETE HostedDirectConnectStatus
	PENDING_UPDATE HostedDirectConnectStatus
	PENDING_CREATE HostedDirectConnectStatus
}

func GetHostedDirectConnectStatusEnum() HostedDirectConnectStatusEnum {
	return HostedDirectConnectStatusEnum{
		BUILD: HostedDirectConnectStatus{
			value: "BUILD",
		},
		ACTIVE: HostedDirectConnectStatus{
			value: "ACTIVE",
		},
		DOWN: HostedDirectConnectStatus{
			value: "DOWN",
		},
		ERROR: HostedDirectConnectStatus{
			value: "ERROR",
		},
		PENDING_DELETE: HostedDirectConnectStatus{
			value: "PENDING_DELETE",
		},
		PENDING_UPDATE: HostedDirectConnectStatus{
			value: "PENDING_UPDATE",
		},
		PENDING_CREATE: HostedDirectConnectStatus{
			value: "PENDING_CREATE",
		},
	}
}

func (c HostedDirectConnectStatus) Value() string {
	return c.value
}

func (c HostedDirectConnectStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *HostedDirectConnectStatus) UnmarshalJSON(b []byte) error {
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

type HostedDirectConnectProviderStatus struct {
	value string
}

type HostedDirectConnectProviderStatusEnum struct {
	ACTIVE HostedDirectConnectProviderStatus
	DOWN   HostedDirectConnectProviderStatus
}

func GetHostedDirectConnectProviderStatusEnum() HostedDirectConnectProviderStatusEnum {
	return HostedDirectConnectProviderStatusEnum{
		ACTIVE: HostedDirectConnectProviderStatus{
			value: "ACTIVE",
		},
		DOWN: HostedDirectConnectProviderStatus{
			value: "DOWN",
		},
	}
}

func (c HostedDirectConnectProviderStatus) Value() string {
	return c.value
}

func (c HostedDirectConnectProviderStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *HostedDirectConnectProviderStatus) UnmarshalJSON(b []byte) error {
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

type HostedDirectConnectPortType struct {
	value string
}

type HostedDirectConnectPortTypeEnum struct {
	E_1_G   HostedDirectConnectPortType
	E_10_G  HostedDirectConnectPortType
	E_40_G  HostedDirectConnectPortType
	E_100_G HostedDirectConnectPortType
}

func GetHostedDirectConnectPortTypeEnum() HostedDirectConnectPortTypeEnum {
	return HostedDirectConnectPortTypeEnum{
		E_1_G: HostedDirectConnectPortType{
			value: "1G",
		},
		E_10_G: HostedDirectConnectPortType{
			value: "10G",
		},
		E_40_G: HostedDirectConnectPortType{
			value: "40G",
		},
		E_100_G: HostedDirectConnectPortType{
			value: "100G",
		},
	}
}

func (c HostedDirectConnectPortType) Value() string {
	return c.value
}

func (c HostedDirectConnectPortType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *HostedDirectConnectPortType) UnmarshalJSON(b []byte) error {
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

type HostedDirectConnectType struct {
	value string
}

type HostedDirectConnectTypeEnum struct {
	STANDARD         HostedDirectConnectType
	HOSTING          HostedDirectConnectType
	HOSTED           HostedDirectConnectType
	ONESTOP_STANDARD HostedDirectConnectType
	ONESTOP_HOSTED   HostedDirectConnectType
}

func GetHostedDirectConnectTypeEnum() HostedDirectConnectTypeEnum {
	return HostedDirectConnectTypeEnum{
		STANDARD: HostedDirectConnectType{
			value: "standard",
		},
		HOSTING: HostedDirectConnectType{
			value: "hosting",
		},
		HOSTED: HostedDirectConnectType{
			value: "hosted",
		},
		ONESTOP_STANDARD: HostedDirectConnectType{
			value: "onestop_standard",
		},
		ONESTOP_HOSTED: HostedDirectConnectType{
			value: "onestop_hosted",
		},
	}
}

func (c HostedDirectConnectType) Value() string {
	return c.value
}

func (c HostedDirectConnectType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *HostedDirectConnectType) UnmarshalJSON(b []byte) error {
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
