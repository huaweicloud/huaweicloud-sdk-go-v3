package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListVpcPeeringInfosRequest Request Object
type ListVpcPeeringInfosRequest struct {

	// **参数解释**： 对等连接的资源ID，填写后按照ID进行过滤，支持传入多个ID过滤。 **取值范围**： 不涉及。
	Id *[]string `json:"id,omitempty"`

	// **参数解释**： 每页返回的资源个数。 **取值范围**： 0~2000。
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**： 分页查询起始的资源ID，为空时查询第一页。 **取值范围**： 对等连接的资源ID。
	Marker *string `json:"marker,omitempty"`

	// **参数解释**： 对等连接的名称，填写后按照名称进行过滤，支持传入多个名称过滤。 **取值范围**： 不涉及。
	Name *[]string `json:"name,omitempty"`

	// **参数解释**： 对等连接的状态，填写后按照状态进行过滤。 **取值范围**： - PENDING_ACCEPTANCE：等待接受 - REJECTED：已拒绝。 - EXPIRED：已过期。 - DELETED：已删除。 - ACTIVE：生效的。
	Status *ListVpcPeeringInfosRequestStatus `json:"status,omitempty"`

	// **参数解释**： 对等连接的项目ID，填写后按照项目ID进行过滤。 **取值范围**： 不涉及。
	TenantId *string `json:"tenant_id,omitempty"`

	// **参数解释**： 对等连接两端的VPC ID，填写后按照VPC ID进行过滤，支持传入多个ID过滤。 **取值范围**： 不涉及。
	VpcId *[]string `json:"vpc_id,omitempty"`
}

func (o ListVpcPeeringInfosRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVpcPeeringInfosRequest struct{}"
	}

	return strings.Join([]string{"ListVpcPeeringInfosRequest", string(data)}, " ")
}

type ListVpcPeeringInfosRequestStatus struct {
	value string
}

type ListVpcPeeringInfosRequestStatusEnum struct {
	PENDING_ACCEPTANCE ListVpcPeeringInfosRequestStatus
	REJECTED           ListVpcPeeringInfosRequestStatus
	EXPIRED            ListVpcPeeringInfosRequestStatus
	DELETED            ListVpcPeeringInfosRequestStatus
	ACTIVE             ListVpcPeeringInfosRequestStatus
}

func GetListVpcPeeringInfosRequestStatusEnum() ListVpcPeeringInfosRequestStatusEnum {
	return ListVpcPeeringInfosRequestStatusEnum{
		PENDING_ACCEPTANCE: ListVpcPeeringInfosRequestStatus{
			value: "PENDING_ACCEPTANCE",
		},
		REJECTED: ListVpcPeeringInfosRequestStatus{
			value: "REJECTED",
		},
		EXPIRED: ListVpcPeeringInfosRequestStatus{
			value: "EXPIRED",
		},
		DELETED: ListVpcPeeringInfosRequestStatus{
			value: "DELETED",
		},
		ACTIVE: ListVpcPeeringInfosRequestStatus{
			value: "ACTIVE",
		},
	}
}

func (c ListVpcPeeringInfosRequestStatus) Value() string {
	return c.value
}

func (c ListVpcPeeringInfosRequestStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListVpcPeeringInfosRequestStatus) UnmarshalJSON(b []byte) error {
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
