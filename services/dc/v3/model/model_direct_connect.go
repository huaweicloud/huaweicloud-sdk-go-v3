package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// DirectConnect 物理专线对象
type DirectConnect struct {

	// 物理专线标识符ID
	Id *string `json:"id,omitempty"`

	// 实例所属项目ID。
	TenantId *string `json:"tenant_id,omitempty"`

	// 物理专线名字
	Name *string `json:"name,omitempty"`

	// 物理专线描述信息
	Description *string `json:"description,omitempty"`

	// 物理专线接入接口的类型，支持1G 10G 40G 100G
	PortType *DirectConnectPortType `json:"port_type,omitempty"`

	// 物理专线接入带宽，单位Mbps。
	Bandwidth *int32 `json:"bandwidth,omitempty"`

	// 专线的接入位置信息
	Location *string `json:"location,omitempty"`

	// 物理专线对端所在的物理位置，省/市/街道或IDC名字。
	PeerLocation *string `json:"peer_location,omitempty"`

	// 物理专线连接的设备的标识ID
	DeviceId *string `json:"device_id,omitempty"`

	// 物理专线的类型，类似包括标准(standard),运营专线(hosting),托管专线（hosted）
	Type *DirectConnectType `json:"type,omitempty"`

	// hosted物理专线对应的hosting物理专线的ID
	HostingId *string `json:"hosting_id,omitempty"`

	// 计费模式：prepayment/bandwidth/traffic
	ChargeMode *DirectConnectChargeMode `json:"charge_mode,omitempty"`

	// 物理专线连接的线路运营商如：中国电信 中国联通 中国移动 中国其他 境外其他专线归属的运营商
	Provider *string `json:"provider,omitempty"`

	// 管理状态：true或false
	AdminStateUp *bool `json:"admin_state_up,omitempty"`

	// 为托管hosted物理专线分配的vlan。
	Vlan *int32 `json:"vlan,omitempty"`

	// 资源状态，合法值是：ACTIVE，DOWN，BUILD，ERROR，PENDING_DELETE，DELETED，APPLY，DENY，PENDING_PAY，PAID，ORDERING，ACCEPT，REJECTED
	Status *DirectConnectStatus `json:"status,omitempty"`

	// 物理专线的申请时间
	ApplyTime *sdktime.SdkTime `json:"apply_time,omitempty"`

	// 物理专线的创建时间
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// 物理专线的运营商操作状态，合法值是：ACTIVE， DOWN
	ProviderStatus *DirectConnectProviderStatus `json:"provider_status,omitempty"`

	// 连接对端的端口类型
	PeerPortType *string `json:"peer_port_type,omitempty"`

	// 专线连接对接的运营商
	PeerProvider *string `json:"peer_provider,omitempty"`

	// 物理专线对应订单号，用于支持包周期计费，识别用户订单
	OrderId *string `json:"order_id,omitempty"`

	// 物理专线订单对应产品标识，用于订制包周期套餐等计费策略
	ProductId *string `json:"product_id,omitempty"`

	// 物理专线订单对应产品规格，用于订制包周期套餐等计费策略
	SpecCode *string `json:"spec_code,omitempty"`

	// 物理专线对应订单号对应包周期的类型
	PeriodType *int32 `json:"period_type,omitempty"`

	// 物理专线对应的包周期时间
	PeriodNum *int32 `json:"period_num,omitempty"`

	// 专线要求的网关类型
	VgwType *DirectConnectVgwType `json:"vgw_type,omitempty"`

	// 物理专线归属的链路聚合组(lag）的ID
	LagId *string `json:"lag_id,omitempty"`

	// 专线协议的签暑状态
	SignedAgreementStatus *DirectConnectSignedAgreementStatus `json:"signed_agreement_status,omitempty"`

	// 实例所属企业项目ID
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 标签信息
	Tags *[]Tag `json:"tags,omitempty"`
}

func (o DirectConnect) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DirectConnect struct{}"
	}

	return strings.Join([]string{"DirectConnect", string(data)}, " ")
}

type DirectConnectPortType struct {
	value string
}

type DirectConnectPortTypeEnum struct {
	E_1_G   DirectConnectPortType
	E_10_G  DirectConnectPortType
	E_40_G  DirectConnectPortType
	E_100_G DirectConnectPortType
}

func GetDirectConnectPortTypeEnum() DirectConnectPortTypeEnum {
	return DirectConnectPortTypeEnum{
		E_1_G: DirectConnectPortType{
			value: "1G",
		},
		E_10_G: DirectConnectPortType{
			value: "10G",
		},
		E_40_G: DirectConnectPortType{
			value: "40G",
		},
		E_100_G: DirectConnectPortType{
			value: "100G",
		},
	}
}

func (c DirectConnectPortType) Value() string {
	return c.value
}

func (c DirectConnectPortType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DirectConnectPortType) UnmarshalJSON(b []byte) error {
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

type DirectConnectType struct {
	value string
}

type DirectConnectTypeEnum struct {
	STANDARD DirectConnectType
	HOSTING  DirectConnectType
	HOSTED   DirectConnectType
}

func GetDirectConnectTypeEnum() DirectConnectTypeEnum {
	return DirectConnectTypeEnum{
		STANDARD: DirectConnectType{
			value: "standard",
		},
		HOSTING: DirectConnectType{
			value: "hosting",
		},
		HOSTED: DirectConnectType{
			value: "hosted",
		},
	}
}

func (c DirectConnectType) Value() string {
	return c.value
}

func (c DirectConnectType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DirectConnectType) UnmarshalJSON(b []byte) error {
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

type DirectConnectChargeMode struct {
	value string
}

type DirectConnectChargeModeEnum struct {
	PREPAYMENT DirectConnectChargeMode
	BANDWIDTH  DirectConnectChargeMode
	TRAFFIC    DirectConnectChargeMode
}

func GetDirectConnectChargeModeEnum() DirectConnectChargeModeEnum {
	return DirectConnectChargeModeEnum{
		PREPAYMENT: DirectConnectChargeMode{
			value: "prepayment",
		},
		BANDWIDTH: DirectConnectChargeMode{
			value: "bandwidth",
		},
		TRAFFIC: DirectConnectChargeMode{
			value: "traffic",
		},
	}
}

func (c DirectConnectChargeMode) Value() string {
	return c.value
}

func (c DirectConnectChargeMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DirectConnectChargeMode) UnmarshalJSON(b []byte) error {
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

type DirectConnectStatus struct {
	value string
}

type DirectConnectStatusEnum struct {
	BUILD          DirectConnectStatus
	PAID           DirectConnectStatus
	APPLY          DirectConnectStatus
	PENDING_SURVEY DirectConnectStatus
	ACTIVE         DirectConnectStatus
	DOWN           DirectConnectStatus
	ERROR          DirectConnectStatus
	PENDING_DELETE DirectConnectStatus
	DELETED        DirectConnectStatus
	DENY           DirectConnectStatus
	PENDING_PAY    DirectConnectStatus
	ORDERING       DirectConnectStatus
	ACCEPT         DirectConnectStatus
	REJECTED       DirectConnectStatus
}

func GetDirectConnectStatusEnum() DirectConnectStatusEnum {
	return DirectConnectStatusEnum{
		BUILD: DirectConnectStatus{
			value: "BUILD",
		},
		PAID: DirectConnectStatus{
			value: "PAID",
		},
		APPLY: DirectConnectStatus{
			value: "APPLY",
		},
		PENDING_SURVEY: DirectConnectStatus{
			value: "PENDING_SURVEY",
		},
		ACTIVE: DirectConnectStatus{
			value: "ACTIVE",
		},
		DOWN: DirectConnectStatus{
			value: "DOWN",
		},
		ERROR: DirectConnectStatus{
			value: "ERROR",
		},
		PENDING_DELETE: DirectConnectStatus{
			value: "PENDING_DELETE",
		},
		DELETED: DirectConnectStatus{
			value: "DELETED",
		},
		DENY: DirectConnectStatus{
			value: "DENY",
		},
		PENDING_PAY: DirectConnectStatus{
			value: "PENDING_PAY",
		},
		ORDERING: DirectConnectStatus{
			value: "ORDERING",
		},
		ACCEPT: DirectConnectStatus{
			value: "ACCEPT",
		},
		REJECTED: DirectConnectStatus{
			value: "REJECTED",
		},
	}
}

func (c DirectConnectStatus) Value() string {
	return c.value
}

func (c DirectConnectStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DirectConnectStatus) UnmarshalJSON(b []byte) error {
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

type DirectConnectProviderStatus struct {
	value string
}

type DirectConnectProviderStatusEnum struct {
	ACTIVE DirectConnectProviderStatus
	DOWN   DirectConnectProviderStatus
}

func GetDirectConnectProviderStatusEnum() DirectConnectProviderStatusEnum {
	return DirectConnectProviderStatusEnum{
		ACTIVE: DirectConnectProviderStatus{
			value: "ACTIVE",
		},
		DOWN: DirectConnectProviderStatus{
			value: "DOWN",
		},
	}
}

func (c DirectConnectProviderStatus) Value() string {
	return c.value
}

func (c DirectConnectProviderStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DirectConnectProviderStatus) UnmarshalJSON(b []byte) error {
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

type DirectConnectVgwType struct {
	value string
}

type DirectConnectVgwTypeEnum struct {
	DEFAULT DirectConnectVgwType
}

func GetDirectConnectVgwTypeEnum() DirectConnectVgwTypeEnum {
	return DirectConnectVgwTypeEnum{
		DEFAULT: DirectConnectVgwType{
			value: "default",
		},
	}
}

func (c DirectConnectVgwType) Value() string {
	return c.value
}

func (c DirectConnectVgwType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DirectConnectVgwType) UnmarshalJSON(b []byte) error {
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

type DirectConnectSignedAgreementStatus struct {
	value string
}

type DirectConnectSignedAgreementStatusEnum struct {
	SIGNED DirectConnectSignedAgreementStatus
}

func GetDirectConnectSignedAgreementStatusEnum() DirectConnectSignedAgreementStatusEnum {
	return DirectConnectSignedAgreementStatusEnum{
		SIGNED: DirectConnectSignedAgreementStatus{
			value: "signed",
		},
	}
}

func (c DirectConnectSignedAgreementStatus) Value() string {
	return c.value
}

func (c DirectConnectSignedAgreementStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DirectConnectSignedAgreementStatus) UnmarshalJSON(b []byte) error {
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
