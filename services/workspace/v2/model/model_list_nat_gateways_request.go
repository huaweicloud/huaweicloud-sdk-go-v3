package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListNatGatewaysRequest Request Object
type ListNatGatewaysRequest struct {

	// 公网NAT网关实例的ID。
	Id *string `json:"id,omitempty"`

	// 企业项目ID。创建公网NAT网关实例时，关联的企业项目ID。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 公网NAT网关实例的描述，长度限制为255。
	Description *string `json:"description,omitempty"`

	// 公网NAT网关实例的创建时间，遵循UTC时间，格式是yyyy-mm-ddThh:mm:ssZ。
	CreatedAt *string `json:"created_at,omitempty"`

	// 公网NAT网关实例的名字，长度限制为64。公网NAT网关实例的名字仅支持数字、字母、_（下划线）、-（中划线）、中文。
	Name *string `json:"name,omitempty"`

	// 公网NAT网关实例的状态。 枚举值：  ACTIVE PENDING_CREATE PENDING_UPDATE PENDING_DELETE INACTIVE
	Status *[]ListNatGatewaysRequestStatus `json:"status,omitempty"`

	// 公网NAT网关实例的规格。取值为： \"1\"：小型，SNAT最大连接数10000；\"2\"：中型，SNAT最大连接数50000；\"3\"：大型，SNAT最大连接数200000；\"4\"：超大型，SNAT最大连接数1000000。
	Spec *[]string `json:"spec,omitempty"`

	// VPC的id。
	RouterId *string `json:"router_id,omitempty"`

	// 功能说明：每页返回的个数。取值范围：1~2000。默认值：2000。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListNatGatewaysRequest) String() string {
	data, err := utils.Marshal(o)
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
	INACTIVE       ListNatGatewaysRequestStatus
	PENDING_CREATE ListNatGatewaysRequestStatus
	PENDING_UPDATE ListNatGatewaysRequestStatus
	PENDING_DELETE ListNatGatewaysRequestStatus
}

func GetListNatGatewaysRequestStatusEnum() ListNatGatewaysRequestStatusEnum {
	return ListNatGatewaysRequestStatusEnum{
		ACTIVE: ListNatGatewaysRequestStatus{
			value: "ACTIVE",
		},
		INACTIVE: ListNatGatewaysRequestStatus{
			value: "INACTIVE",
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
	}
}

func (c ListNatGatewaysRequestStatus) Value() string {
	return c.value
}

func (c ListNatGatewaysRequestStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListNatGatewaysRequestStatus) UnmarshalJSON(b []byte) error {
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
