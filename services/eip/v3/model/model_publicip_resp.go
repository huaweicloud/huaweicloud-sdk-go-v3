package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// PublicipResp 弹性公网IP对象返回模板
type PublicipResp struct {

	// - 功能说明：弹性公网IP的唯一标识
	Id *string `json:"id,omitempty"`

	// - 功能说明：项目ID
	ProjectId *string `json:"project_id,omitempty"`

	// - 功能说明: 弹性公网IP版本号 - 取值范围: 4、6   - 4表示公网IP地址为public_ip_address地址   - 6表示公网IP地址为public_ipv6_address地址
	IpVersion *PublicipRespIpVersion `json:"ip_version,omitempty"`

	// - 功能说明: 弹性公网IPv4地址
	PublicIpAddress *string `json:"public_ip_address,omitempty"`

	// - 功能说明: 弹性公网IPv6地址
	PublicIpv6Address *string `json:"public_ipv6_address,omitempty"`

	// - 功能说明：弹性公网IP的状态 - 取值范围：FREEZED，DOWN，ACTIVE，ERROR。   - FREEZED表示弹性公网IP处于冻结状态   - DOWN表示弹性公网IP未绑定实例   - ACTIVE表示弹性公网IP绑定实例，正在使用中   - ERROR表示弹性公网IP状态异常
	Status *PublicipRespStatus `json:"status,omitempty"`

	// - 功能说明：弹性公网IP的描述信息 - 约束：用户以自定义方式标识资源，系统不感知
	Description *string `json:"description,omitempty"`

	// - 功能说明：弹性公网IP的创建时间 - 约束：UTC时间格式（2018-12-25T10:07:24Z）
	CreatedAt *string `json:"created_at,omitempty"`

	// - 功能说明：弹性公网IP最近的更新时间 - 约束：UTC时间格式（2018-12-25T10:09:20Z）
	UpdatedAt *string `json:"updated_at,omitempty"`

	// - 功能说明：弹性公网IP的类型 - 取值范围：EIP，DUALSTACK
	Type *PublicipRespType `json:"type,omitempty"`

	Vnic *VnicResp `json:"vnic,omitempty"`

	Bandwidth *BandwidthResp `json:"bandwidth,omitempty"`

	// - 功能说明：弹性公网IP的企业项目ID
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// - 功能说明：弹性公网IP的订单信息 - 约束：包周期才会有订单信息，按需资源此字段为空
	BillingInfo *string `json:"billing_info,omitempty"`

	// - 功能说明：记录弹性公网IP当前的冻结状态 - 约束：metadata类型，标识欠费冻结、公安冻结
	LockStatus *string `json:"lock_status,omitempty"`

	// - 功能说明：弹性公网IP绑定的实例类型 - 取值范围：PORT、NATGW、ELB、ELBV1、VPN
	AssociateInstanceType *PublicipRespAssociateInstanceType `json:"associate_instance_type,omitempty"`

	// - 功能说明：弹性公网IP绑定的实例ID
	AssociateInstanceId *string `json:"associate_instance_id,omitempty"`

	// - 功能说明：弹性公网IP所属网络的ID。publicip_pool_name对应的网络ID
	PublicipPoolId *string `json:"publicip_pool_id,omitempty"`

	// - 功能说明：弹性公网IP的网络类型， 包括公共池类型，如5_bgp/5_sbgp...，和用户购买的专属池。专属池见publcip_pool相关接口
	PublicipPoolName *string `json:"publicip_pool_name,omitempty"`

	// - 功能说明：弹性公网IP别名
	Alias *string `json:"alias,omitempty"`

	// - 功能说明：中心还是边缘。中心CENTER，边缘为各边缘az名称
	PublicBorderGroup *string `json:"public_border_group,omitempty"`
}

func (o PublicipResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PublicipResp struct{}"
	}

	return strings.Join([]string{"PublicipResp", string(data)}, " ")
}

type PublicipRespIpVersion struct {
	value int32
}

type PublicipRespIpVersionEnum struct {
	E_4 PublicipRespIpVersion
	E_6 PublicipRespIpVersion
}

func GetPublicipRespIpVersionEnum() PublicipRespIpVersionEnum {
	return PublicipRespIpVersionEnum{
		E_4: PublicipRespIpVersion{
			value: 4,
		}, E_6: PublicipRespIpVersion{
			value: 6,
		},
	}
}

func (c PublicipRespIpVersion) Value() int32 {
	return c.value
}

func (c PublicipRespIpVersion) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PublicipRespIpVersion) UnmarshalJSON(b []byte) error {
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

type PublicipRespStatus struct {
	value string
}

type PublicipRespStatusEnum struct {
	FREEZED PublicipRespStatus
	DOWN    PublicipRespStatus
	ACTIVE  PublicipRespStatus
	ERROR   PublicipRespStatus
}

func GetPublicipRespStatusEnum() PublicipRespStatusEnum {
	return PublicipRespStatusEnum{
		FREEZED: PublicipRespStatus{
			value: "FREEZED",
		},
		DOWN: PublicipRespStatus{
			value: "DOWN",
		},
		ACTIVE: PublicipRespStatus{
			value: "ACTIVE",
		},
		ERROR: PublicipRespStatus{
			value: "ERROR",
		},
	}
}

func (c PublicipRespStatus) Value() string {
	return c.value
}

func (c PublicipRespStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PublicipRespStatus) UnmarshalJSON(b []byte) error {
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

type PublicipRespType struct {
	value string
}

type PublicipRespTypeEnum struct {
	EIP       PublicipRespType
	DUALSTACK PublicipRespType
}

func GetPublicipRespTypeEnum() PublicipRespTypeEnum {
	return PublicipRespTypeEnum{
		EIP: PublicipRespType{
			value: "EIP",
		},
		DUALSTACK: PublicipRespType{
			value: "DUALSTACK",
		},
	}
}

func (c PublicipRespType) Value() string {
	return c.value
}

func (c PublicipRespType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PublicipRespType) UnmarshalJSON(b []byte) error {
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

type PublicipRespAssociateInstanceType struct {
	value string
}

type PublicipRespAssociateInstanceTypeEnum struct {
	PORT  PublicipRespAssociateInstanceType
	NATGW PublicipRespAssociateInstanceType
	ELB   PublicipRespAssociateInstanceType
	ELBV1 PublicipRespAssociateInstanceType
	VPN   PublicipRespAssociateInstanceType
}

func GetPublicipRespAssociateInstanceTypeEnum() PublicipRespAssociateInstanceTypeEnum {
	return PublicipRespAssociateInstanceTypeEnum{
		PORT: PublicipRespAssociateInstanceType{
			value: "PORT",
		},
		NATGW: PublicipRespAssociateInstanceType{
			value: "NATGW",
		},
		ELB: PublicipRespAssociateInstanceType{
			value: "ELB",
		},
		ELBV1: PublicipRespAssociateInstanceType{
			value: "ELBV1",
		},
		VPN: PublicipRespAssociateInstanceType{
			value: "VPN",
		},
	}
}

func (c PublicipRespAssociateInstanceType) Value() string {
	return c.value
}

func (c PublicipRespAssociateInstanceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PublicipRespAssociateInstanceType) UnmarshalJSON(b []byte) error {
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
