package model

import (
	"encoding/json"
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"strings"
)

// Request Object
type ListNatGatewaySnatRulesRequest struct {
	// 解冻/冻结状态。 取值范围： \"true\"：解冻 \"false\"：冻结

	AdminStateUp *bool `json:"admin_state_up,omitempty"`
	// 可以是网段或者主机格式，与network_id参数二选一。 Source_type=0时，cidr必须是vpc子网网段的子集(不能相等）; Source_type=1时，cidr必须指定专线侧网段。

	Cidr *string `json:"cidr,omitempty"`
	// 功能说明：每页返回的个数。 取值范围：0~2000。 默认值：2000。

	Limit *int32 `json:"limit,omitempty"`
	// 功能说明：弹性公网IP，多个弹性公网IP使用逗号分隔。 取值范围：最大长度1024字节。

	FloatingIpAddress *string `json:"floating_ip_address,omitempty"`
	// 功能说明：弹性公网IP的id，多个弹性公网IP使用逗号分隔。 取值范围：最大长度4096字节。

	FloatingIpId *string `json:"floating_ip_id,omitempty"`
	// SNAT规则的ID。

	Id *string `json:"id,omitempty"`
	// SNAT规则的描述，长度限制为255。

	Description *string `json:"description,omitempty"`
	// SNAT规则的创建时间，遵循UTC时间，格式是yyyy-mm-ddThh:mm:ssZ。

	CreatedAt *sdktime.SdkTime `json:"created_at,omitempty"`
	// 公网NAT网关实例的ID。

	NatGatewayId *[]string `json:"nat_gateway_id,omitempty"`
	// 规则使用的网络id。与cidr参数二选一。

	NetworkId *string `json:"network_id,omitempty"`
	// 0：VPC侧，可以指定network_id 或者cidr 1：专线侧，只能指定cidr 不输入默认为0（VPC）

	SourceType *int32 `json:"source_type,omitempty"`
	// 功能说明：SNAT规则的状态。

	Status *ListNatGatewaySnatRulesRequestStatus `json:"status,omitempty"`
}

func (o ListNatGatewaySnatRulesRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListNatGatewaySnatRulesRequest struct{}"
	}

	return strings.Join([]string{"ListNatGatewaySnatRulesRequest", string(data)}, " ")
}

type ListNatGatewaySnatRulesRequestStatus struct {
	value string
}

type ListNatGatewaySnatRulesRequestStatusEnum struct {
	ACTIVE         ListNatGatewaySnatRulesRequestStatus
	PENDING_CREATE ListNatGatewaySnatRulesRequestStatus
	PENDING_UPDATE ListNatGatewaySnatRulesRequestStatus
	PENDING_DELETE ListNatGatewaySnatRulesRequestStatus
	EIP_FREEZED    ListNatGatewaySnatRulesRequestStatus
	INACTIVE       ListNatGatewaySnatRulesRequestStatus
}

func GetListNatGatewaySnatRulesRequestStatusEnum() ListNatGatewaySnatRulesRequestStatusEnum {
	return ListNatGatewaySnatRulesRequestStatusEnum{
		ACTIVE: ListNatGatewaySnatRulesRequestStatus{
			value: "ACTIVE",
		},
		PENDING_CREATE: ListNatGatewaySnatRulesRequestStatus{
			value: "PENDING_CREATE",
		},
		PENDING_UPDATE: ListNatGatewaySnatRulesRequestStatus{
			value: "PENDING_UPDATE",
		},
		PENDING_DELETE: ListNatGatewaySnatRulesRequestStatus{
			value: "PENDING_DELETE",
		},
		EIP_FREEZED: ListNatGatewaySnatRulesRequestStatus{
			value: "EIP_FREEZED",
		},
		INACTIVE: ListNatGatewaySnatRulesRequestStatus{
			value: "INACTIVE",
		},
	}
}

func (c ListNatGatewaySnatRulesRequestStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *ListNatGatewaySnatRulesRequestStatus) UnmarshalJSON(b []byte) error {
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
