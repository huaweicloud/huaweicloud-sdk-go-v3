package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// NatGatewayResponseBody 公网NAT网关实例的响应体。
type NatGatewayResponseBody struct {

	// 公网NAT网关实例的ID。
	Id *string `json:"id,omitempty"`

	// 项目的ID。
	TenantId *string `json:"tenant_id,omitempty"`

	// 公网NAT网关实例的名字，长度限制为64。
	Name *string `json:"name,omitempty"`

	// 公网NAT网关实例的描述，长度范围小于等于255个字符，不能包含“<”和“>”。
	Description *string `json:"description,omitempty"`

	// 公网NAT网关的规格。 取值为： “1”：小型，SNAT最大连接数10000 “2”：中型，SNAT最大连接数50000 “3”：大型，SNAT最大连接数200000 “4”：超大型，SNAT最大连接数1000000 “5”：企业型，SNAT最大连接数10000000
	Spec *NatGatewayResponseBodySpec `json:"spec,omitempty"`

	// 公网NAT网关实例的状态。 取值为:  ACTIVE: 可用 PENDING_CREATE: 创建中 PENDING_UPDATE: 更新中 PENDING_DELETE: 删除中 INACTIVE: 不可用
	Status *NatGatewayResponseBodyStatus `json:"status,omitempty"`

	// 解冻/冻结状态。 取值范围： - \"true\"：解冻 - \"false\"：冻结
	AdminStateUp *bool `json:"admin_state_up,omitempty"`

	// 公网NAT网关实例的创建时间，遵循UTC时间，格式是yyyy-mm-ddThh:mm:ssZ。
	CreatedAt *sdktime.SdkTime `json:"created_at,omitempty"`

	// VPC的id。
	RouterId *string `json:"router_id,omitempty"`

	// 公网NAT网关下行口（DVR的下一跳）所属的network id。
	InternalNetworkId *string `json:"internal_network_id,omitempty"`

	// 企业项目ID。 创建公网NAT网关实例时，关联的企业项目ID。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	SessionConf *SessionConfiguration `json:"session_conf,omitempty"`

	// 公网NAT网关私有IP地址，由VPC中子网分配。
	NgportIpAddress *string `json:"ngport_ip_address,omitempty"`

	// 订单信息。此字段只有在订购包周期资源时才会有订单信息，而在订购按需资源时则为空。
	BillingInfo *string `json:"billing_info,omitempty"`

	// 公网NAT网关下DNAT规则数量限制，默认为200。
	DnatRulesLimit *int64 `json:"dnat_rules_limit,omitempty"`

	// 公网NAT网关下SNAT规则EIP池中EIP数量限制，默认为20。
	SnatRulePublicIpLimit *int32 `json:"snat_rule_public_ip_limit,omitempty"`

	// 网关每秒能够接收或发送的数据包数量，代表网络设备的流量处理能力。
	PpsMax *int32 `json:"pps_max,omitempty"`

	// 网关每秒能够接收或发送的带宽大小，代表网络设备的流量处理能力。单位M
	BpsMax *int32 `json:"bps_max,omitempty"`
}

func (o NatGatewayResponseBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NatGatewayResponseBody struct{}"
	}

	return strings.Join([]string{"NatGatewayResponseBody", string(data)}, " ")
}

type NatGatewayResponseBodySpec struct {
	value string
}

type NatGatewayResponseBodySpecEnum struct {
	E_1 NatGatewayResponseBodySpec
	E_2 NatGatewayResponseBodySpec
	E_3 NatGatewayResponseBodySpec
	E_4 NatGatewayResponseBodySpec
	E_5 NatGatewayResponseBodySpec
}

func GetNatGatewayResponseBodySpecEnum() NatGatewayResponseBodySpecEnum {
	return NatGatewayResponseBodySpecEnum{
		E_1: NatGatewayResponseBodySpec{
			value: "1",
		},
		E_2: NatGatewayResponseBodySpec{
			value: "2",
		},
		E_3: NatGatewayResponseBodySpec{
			value: "3",
		},
		E_4: NatGatewayResponseBodySpec{
			value: "4",
		},
		E_5: NatGatewayResponseBodySpec{
			value: "5",
		},
	}
}

func (c NatGatewayResponseBodySpec) Value() string {
	return c.value
}

func (c NatGatewayResponseBodySpec) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *NatGatewayResponseBodySpec) UnmarshalJSON(b []byte) error {
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

type NatGatewayResponseBodyStatus struct {
	value string
}

type NatGatewayResponseBodyStatusEnum struct {
	ACTIVE         NatGatewayResponseBodyStatus
	PENDING_CREATE NatGatewayResponseBodyStatus
	PENDING_UPDATE NatGatewayResponseBodyStatus
	PENDING_DELETE NatGatewayResponseBodyStatus
	INACTIVE       NatGatewayResponseBodyStatus
}

func GetNatGatewayResponseBodyStatusEnum() NatGatewayResponseBodyStatusEnum {
	return NatGatewayResponseBodyStatusEnum{
		ACTIVE: NatGatewayResponseBodyStatus{
			value: "ACTIVE",
		},
		PENDING_CREATE: NatGatewayResponseBodyStatus{
			value: "PENDING_CREATE",
		},
		PENDING_UPDATE: NatGatewayResponseBodyStatus{
			value: "PENDING_UPDATE",
		},
		PENDING_DELETE: NatGatewayResponseBodyStatus{
			value: "PENDING_DELETE",
		},
		INACTIVE: NatGatewayResponseBodyStatus{
			value: "INACTIVE",
		},
	}
}

func (c NatGatewayResponseBodyStatus) Value() string {
	return c.value
}

func (c NatGatewayResponseBodyStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *NatGatewayResponseBodyStatus) UnmarshalJSON(b []byte) error {
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
