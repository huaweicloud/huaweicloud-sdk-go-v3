package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateNatGatewaySnatRuleResponseBody 创建SNAT规则的响应体。
type CreateNatGatewaySnatRuleResponseBody struct {

	// SNAT规则的ID。
	Id string `json:"id"`

	// 项目的ID。
	TenantId string `json:"tenant_id"`

	// 公网NAT网关实例的ID。
	NatGatewayId string `json:"nat_gateway_id"`

	// cidr，可以是网段或者主机格式，与network_id参数二选一。 Source_type=0时，cidr必须是vpc 子网网段的子集(不能相等）; Source_type=1时，cidr必须指定专线侧网段。
	Cidr string `json:"cidr"`

	// 0：VPC侧，可以指定network_id 或者cidr 1：专线侧，只能指定cidr 不输入默认为0（VPC）
	SourceType int32 `json:"source_type"`

	// 功能说明：弹性公网IP的id，多个弹性公网IP使用逗号分隔。 取值范围：最大长度4096字节。
	FloatingIpId string `json:"floating_ip_id"`

	// SNAT规则的描述，长度限制为255。
	Description string `json:"description"`

	// 功能说明：SNAT规则的状态。
	Status CreateNatGatewaySnatRuleResponseBodyStatus `json:"status"`

	// SNAT规则的创建时间，格式是yyyy-mm-dd hh:mm:ss.SSSSSS。
	CreatedAt string `json:"created_at"`

	// 规则使用的网络id。与cidr参数二选一。
	NetworkId string `json:"network_id"`

	// 解冻/冻结状态。 取值范围： - \"true\"：解冻 - \"false\"：冻结
	AdminStateUp bool `json:"admin_state_up"`

	// 功能说明：弹性公网IP，多个弹性公网IP使用逗号分隔。 取值范围：最大长度1024字节。
	FloatingIpAddress string `json:"floating_ip_address"`
}

func (o CreateNatGatewaySnatRuleResponseBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateNatGatewaySnatRuleResponseBody struct{}"
	}

	return strings.Join([]string{"CreateNatGatewaySnatRuleResponseBody", string(data)}, " ")
}

type CreateNatGatewaySnatRuleResponseBodyStatus struct {
	value string
}

type CreateNatGatewaySnatRuleResponseBodyStatusEnum struct {
	ACTIVE         CreateNatGatewaySnatRuleResponseBodyStatus
	PENDING_CREATE CreateNatGatewaySnatRuleResponseBodyStatus
	PENDING_UPDATE CreateNatGatewaySnatRuleResponseBodyStatus
	PENDING_DELETE CreateNatGatewaySnatRuleResponseBodyStatus
	EIP_FREEZED    CreateNatGatewaySnatRuleResponseBodyStatus
	INACTIVE       CreateNatGatewaySnatRuleResponseBodyStatus
}

func GetCreateNatGatewaySnatRuleResponseBodyStatusEnum() CreateNatGatewaySnatRuleResponseBodyStatusEnum {
	return CreateNatGatewaySnatRuleResponseBodyStatusEnum{
		ACTIVE: CreateNatGatewaySnatRuleResponseBodyStatus{
			value: "ACTIVE",
		},
		PENDING_CREATE: CreateNatGatewaySnatRuleResponseBodyStatus{
			value: "PENDING_CREATE",
		},
		PENDING_UPDATE: CreateNatGatewaySnatRuleResponseBodyStatus{
			value: "PENDING_UPDATE",
		},
		PENDING_DELETE: CreateNatGatewaySnatRuleResponseBodyStatus{
			value: "PENDING_DELETE",
		},
		EIP_FREEZED: CreateNatGatewaySnatRuleResponseBodyStatus{
			value: "EIP_FREEZED",
		},
		INACTIVE: CreateNatGatewaySnatRuleResponseBodyStatus{
			value: "INACTIVE",
		},
	}
}

func (c CreateNatGatewaySnatRuleResponseBodyStatus) Value() string {
	return c.value
}

func (c CreateNatGatewaySnatRuleResponseBodyStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateNatGatewaySnatRuleResponseBodyStatus) UnmarshalJSON(b []byte) error {
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
