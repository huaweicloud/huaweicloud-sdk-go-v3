package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// NatGatewayUpdateSnatRuleResponseBody 更新SNAT规则的响应体。
type NatGatewayUpdateSnatRuleResponseBody struct {

	// SNAT规则的ID。
	Id string `json:"id"`

	// 项目的ID。
	TenantId string `json:"tenant_id"`

	// 公网NAT网关实例的ID。
	NatGatewayId string `json:"nat_gateway_id"`

	// 0：VPC侧，可以指定network_id 或者cidr 1：专线侧，只能指定cidr 不输入默认为0（VPC）
	SourceType int32 `json:"source_type"`

	// cidr，可以是网段或者主机格式，与network_id参数二选一。 Source_type=0时，cidr必须是vpc 子网网段的子集(不能相等）; Source_type=1时，cidr必须指定专线侧网段。
	Cidr string `json:"cidr"`

	// 功能说明：弹性公网IP的id，多个弹性公网IP使用逗号分隔。 取值范围：最大长度4096字节。
	FloatingIpId string `json:"floating_ip_id"`

	// SNAT规则的描述，长度限制为255。
	Description string `json:"description"`

	// 功能说明：SNAT规则的状态。
	Status NatGatewayUpdateSnatRuleResponseBodyStatus `json:"status"`

	// SNAT规则的创建时间，格式是yyyy-mm-dd hh:mm:ss.SSSSSS。
	CreatedAt string `json:"created_at"`

	// 规则使用的网络id。与cidr参数二选一。
	NetworkId string `json:"network_id"`

	// 解冻/冻结状态。 取值范围： - \"true\"：解冻 - \"false\"：冻结
	AdminStateUp bool `json:"admin_state_up"`

	// 功能说明：弹性公网IP，多个弹性公网IP使用逗号分隔。 取值范围：最大长度1024字节。
	FloatingIpAddress string `json:"floating_ip_address"`

	// 功能说明：弹性公网IP，多个弹性公网IP使用逗号分隔。 取值范围：最大长度1024字节。
	PublicIpAddress string `json:"public_ip_address"`
}

func (o NatGatewayUpdateSnatRuleResponseBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NatGatewayUpdateSnatRuleResponseBody struct{}"
	}

	return strings.Join([]string{"NatGatewayUpdateSnatRuleResponseBody", string(data)}, " ")
}

type NatGatewayUpdateSnatRuleResponseBodyStatus struct {
	value string
}

type NatGatewayUpdateSnatRuleResponseBodyStatusEnum struct {
	ACTIVE         NatGatewayUpdateSnatRuleResponseBodyStatus
	PENDING_CREATE NatGatewayUpdateSnatRuleResponseBodyStatus
	PENDING_UPDATE NatGatewayUpdateSnatRuleResponseBodyStatus
	PENDING_DELETE NatGatewayUpdateSnatRuleResponseBodyStatus
	EIP_FREEZED    NatGatewayUpdateSnatRuleResponseBodyStatus
	INACTIVE       NatGatewayUpdateSnatRuleResponseBodyStatus
}

func GetNatGatewayUpdateSnatRuleResponseBodyStatusEnum() NatGatewayUpdateSnatRuleResponseBodyStatusEnum {
	return NatGatewayUpdateSnatRuleResponseBodyStatusEnum{
		ACTIVE: NatGatewayUpdateSnatRuleResponseBodyStatus{
			value: "ACTIVE",
		},
		PENDING_CREATE: NatGatewayUpdateSnatRuleResponseBodyStatus{
			value: "PENDING_CREATE",
		},
		PENDING_UPDATE: NatGatewayUpdateSnatRuleResponseBodyStatus{
			value: "PENDING_UPDATE",
		},
		PENDING_DELETE: NatGatewayUpdateSnatRuleResponseBodyStatus{
			value: "PENDING_DELETE",
		},
		EIP_FREEZED: NatGatewayUpdateSnatRuleResponseBodyStatus{
			value: "EIP_FREEZED",
		},
		INACTIVE: NatGatewayUpdateSnatRuleResponseBodyStatus{
			value: "INACTIVE",
		},
	}
}

func (c NatGatewayUpdateSnatRuleResponseBodyStatus) Value() string {
	return c.value
}

func (c NatGatewayUpdateSnatRuleResponseBodyStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *NatGatewayUpdateSnatRuleResponseBodyStatus) UnmarshalJSON(b []byte) error {
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
