package model

import (
	"encoding/json"
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"strings"
)

// Request Object
type ListNatGatewaysRequest struct {
	// 项目的ID。

	TenantId *string `json:"tenant_id,omitempty"`
	// 公网NAT网关实例的ID。

	Id *string `json:"id,omitempty"`
	// 企业项目ID。创建公网NAT网关实例时，关联的企业项目ID。

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
	// 公网NAT网关实例的描述，长度限制为255。

	Description *string `json:"description,omitempty"`
	// 公网NAT网关实例的创建时间，遵循UTC时间，格式是yyyy-mm-ddThh:mm:ssZ。

	CreatedAt *sdktime.SdkTime `json:"created_at,omitempty"`
	// 公网NAT网关实例的名字，长度限制为64。 公网NAT网关实例的名字仅支持数字、字母、_（下划线）、-（中划线）、中文

	Name *string `json:"name,omitempty"`
	// 公网NAT网关实例的状态。

	Status *[]ListNatGatewaysRequestStatus `json:"status,omitempty"`
	// 公网NAT网关实例的规格。 取值为： \"1\"：小型，SNAT最大连接数10000 \"2\"：中型，SNAT最大连接数50000 \"3\"：大型，SNAT最大连接数200000 \"4\"：超大型，SNAT最大连接数1000000

	Spec *[]ListNatGatewaysRequestSpec `json:"spec,omitempty"`
	// 解冻/冻结状态。 取值范围： \"true\"：解冻 \"false\"：冻结

	AdminStateUp *bool `json:"admin_state_up,omitempty"`
	// 公网NAT网关下行口（DVR的下一跳）所属的network id。

	InternalNetworkId *string `json:"internal_network_id,omitempty"`
	// VPC的id。

	RouterId *string `json:"router_id,omitempty"`
	// 功能说明：每页返回的个数。 取值范围：0~2000。 默认值：2000。

	Limit *int32 `json:"limit,omitempty"`
}

func (o ListNatGatewaysRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListNatGatewaysRequest struct{}"
	}

	return strings.Join([]string{"ListNatGatewaysRequest", string(data)}, " ")
}

type ListNatGatewaysRequestStatus struct {
	value string
}

type ListNatGatewaysRequestStatusEnum struct {
	ACTIVE         ListNatGatewaysRequestStatus
	PENDING_CREATE ListNatGatewaysRequestStatus
	PENDING_UPDATE ListNatGatewaysRequestStatus
	PENDING_DELETE ListNatGatewaysRequestStatus
	INACTIVE       ListNatGatewaysRequestStatus
}

func GetListNatGatewaysRequestStatusEnum() ListNatGatewaysRequestStatusEnum {
	return ListNatGatewaysRequestStatusEnum{
		ACTIVE: ListNatGatewaysRequestStatus{
			value: "ACTIVE",
		},
		PENDING_CREATE: ListNatGatewaysRequestStatus{
			value: "PENDING_CREATE",
		},
		PENDING_UPDATE: ListNatGatewaysRequestStatus{
			value: "PENDING_UPDATE",
		},
		PENDING_DELETE: ListNatGatewaysRequestStatus{
			value: "PENDING_DELETE",
		},
		INACTIVE: ListNatGatewaysRequestStatus{
			value: "INACTIVE",
		},
	}
}

func (c ListNatGatewaysRequestStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *ListNatGatewaysRequestStatus) UnmarshalJSON(b []byte) error {
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

type ListNatGatewaysRequestSpec struct {
	value string
}

type ListNatGatewaysRequestSpecEnum struct {
	E_1 ListNatGatewaysRequestSpec
	E_2 ListNatGatewaysRequestSpec
	E_3 ListNatGatewaysRequestSpec
	E_4 ListNatGatewaysRequestSpec
}

func GetListNatGatewaysRequestSpecEnum() ListNatGatewaysRequestSpecEnum {
	return ListNatGatewaysRequestSpecEnum{
		E_1: ListNatGatewaysRequestSpec{
			value: "1",
		},
		E_2: ListNatGatewaysRequestSpec{
			value: "2",
		},
		E_3: ListNatGatewaysRequestSpec{
			value: "3",
		},
		E_4: ListNatGatewaysRequestSpec{
			value: "4",
		},
	}
}

func (c ListNatGatewaysRequestSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *ListNatGatewaysRequestSpec) UnmarshalJSON(b []byte) error {
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
