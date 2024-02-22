package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CancelAsyncInvocationRequestBody struct {

	// 被停止的请求id
	RequestId string `json:"request_id"`

	// 停止的类型 支持recursive, force。 recursive: 停止正在调用的子函数。 force: 直接终止runtime。
	Type *CancelAsyncInvocationRequestBodyType `json:"type,omitempty"`
}

func (o CancelAsyncInvocationRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CancelAsyncInvocationRequestBody struct{}"
	}

	return strings.Join([]string{"CancelAsyncInvocationRequestBody", string(data)}, " ")
}

type CancelAsyncInvocationRequestBodyType struct {
	value string
}

type CancelAsyncInvocationRequestBodyTypeEnum struct {
	FORCE     CancelAsyncInvocationRequestBodyType
	RECURSIVE CancelAsyncInvocationRequestBodyType
}

func GetCancelAsyncInvocationRequestBodyTypeEnum() CancelAsyncInvocationRequestBodyTypeEnum {
	return CancelAsyncInvocationRequestBodyTypeEnum{
		FORCE: CancelAsyncInvocationRequestBodyType{
			value: "force",
		},
		RECURSIVE: CancelAsyncInvocationRequestBodyType{
			value: "recursive",
		},
	}
}

func (c CancelAsyncInvocationRequestBodyType) Value() string {
	return c.value
}

func (c CancelAsyncInvocationRequestBodyType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CancelAsyncInvocationRequestBodyType) UnmarshalJSON(b []byte) error {
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
