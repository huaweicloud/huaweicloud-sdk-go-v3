package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListCloseAccountStatusesRequest Request Object
type ListCloseAccountStatusesRequest struct {

	// 如果正在使用临时安全凭据，则此header是必需的，该值是临时安全凭据的安全令牌（会话令牌）。
	XSecurityToken *string `json:"X-Security-Token,omitempty"`

	// 要包含在响应中的一个或多个状态的列表。如果此参数不存在，则所有请求都包含在响应中。
	States *[]ListCloseAccountStatusesRequestStates `json:"states,omitempty"`
}

func (o ListCloseAccountStatusesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCloseAccountStatusesRequest struct{}"
	}

	return strings.Join([]string{"ListCloseAccountStatusesRequest", string(data)}, " ")
}

type ListCloseAccountStatusesRequestStates struct {
	value string
}

type ListCloseAccountStatusesRequestStatesEnum struct {
	PENDING_CLOSURE ListCloseAccountStatusesRequestStates
	SUSPENDED       ListCloseAccountStatusesRequestStates
}

func GetListCloseAccountStatusesRequestStatesEnum() ListCloseAccountStatusesRequestStatesEnum {
	return ListCloseAccountStatusesRequestStatesEnum{
		PENDING_CLOSURE: ListCloseAccountStatusesRequestStates{
			value: "pending_closure",
		},
		SUSPENDED: ListCloseAccountStatusesRequestStates{
			value: "suspended",
		},
	}
}

func (c ListCloseAccountStatusesRequestStates) Value() string {
	return c.value
}

func (c ListCloseAccountStatusesRequestStates) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListCloseAccountStatusesRequestStates) UnmarshalJSON(b []byte) error {
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
