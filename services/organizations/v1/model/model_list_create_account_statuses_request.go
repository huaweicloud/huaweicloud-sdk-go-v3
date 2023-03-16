package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ListCreateAccountStatusesRequest struct {

	// 要包含在响应中的一个或多个状态的列表。如果此参数不存在，则所有请求都包含在响应中。
	States *[]ListCreateAccountStatusesRequestStates `json:"states,omitempty"`

	// 页面中最大结果数量。
	Limit *int32 `json:"limit,omitempty"`

	// 分页标记。
	Marker *string `json:"marker,omitempty"`
}

func (o ListCreateAccountStatusesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCreateAccountStatusesRequest struct{}"
	}

	return strings.Join([]string{"ListCreateAccountStatusesRequest", string(data)}, " ")
}

type ListCreateAccountStatusesRequestStates struct {
	value string
}

type ListCreateAccountStatusesRequestStatesEnum struct {
	IN_PROGRESS ListCreateAccountStatusesRequestStates
	SUCCEEDED   ListCreateAccountStatusesRequestStates
	FAILED      ListCreateAccountStatusesRequestStates
}

func GetListCreateAccountStatusesRequestStatesEnum() ListCreateAccountStatusesRequestStatesEnum {
	return ListCreateAccountStatusesRequestStatesEnum{
		IN_PROGRESS: ListCreateAccountStatusesRequestStates{
			value: "in_progress",
		},
		SUCCEEDED: ListCreateAccountStatusesRequestStates{
			value: "succeeded",
		},
		FAILED: ListCreateAccountStatusesRequestStates{
			value: "failed",
		},
	}
}

func (c ListCreateAccountStatusesRequestStates) Value() string {
	return c.value
}

func (c ListCreateAccountStatusesRequestStates) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListCreateAccountStatusesRequestStates) UnmarshalJSON(b []byte) error {
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
