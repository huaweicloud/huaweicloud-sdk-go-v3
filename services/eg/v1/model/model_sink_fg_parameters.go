package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type SinkFgParameters struct {

	// 函数执行方式,同步/异步
	InvokeType *SinkFgParametersInvokeType `json:"invoke_type,omitempty"`

	// 函数链接
	Urn *string `json:"urn,omitempty"`

	// 租户委托
	Agency *string `json:"agency,omitempty"`
}

func (o SinkFgParameters) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SinkFgParameters struct{}"
	}

	return strings.Join([]string{"SinkFgParameters", string(data)}, " ")
}

type SinkFgParametersInvokeType struct {
	value string
}

type SinkFgParametersInvokeTypeEnum struct {
	SYNCASYNC SinkFgParametersInvokeType
}

func GetSinkFgParametersInvokeTypeEnum() SinkFgParametersInvokeTypeEnum {
	return SinkFgParametersInvokeTypeEnum{
		SYNCASYNC: SinkFgParametersInvokeType{
			value: "SYNC，ASYNC",
		},
	}
}

func (c SinkFgParametersInvokeType) Value() string {
	return c.value
}

func (c SinkFgParametersInvokeType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SinkFgParametersInvokeType) UnmarshalJSON(b []byte) error {
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
