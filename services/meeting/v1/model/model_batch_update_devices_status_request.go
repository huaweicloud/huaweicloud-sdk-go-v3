package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// BatchUpdateDevicesStatusRequest Request Object
type BatchUpdateDevicesStatusRequest struct {

	// 请求requestId，用来标识一路请求，用于问题跟踪定位，建议使用UUID，若不携带，则后台自动生成。
	XRequestId *string `json:"X-Request-Id,omitempty"`

	// 语言参数，默认为中文zh-CN，英文为en-US。
	AcceptLanguage *string `json:"Accept-Language,omitempty"`

	// 状态。 * 0：启用 * 1：停用
	Value BatchUpdateDevicesStatusRequestValue `json:"value"`

	Body *[]string `json:"body,omitempty"`
}

func (o BatchUpdateDevicesStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateDevicesStatusRequest struct{}"
	}

	return strings.Join([]string{"BatchUpdateDevicesStatusRequest", string(data)}, " ")
}

type BatchUpdateDevicesStatusRequestValue struct {
	value int32
}

type BatchUpdateDevicesStatusRequestValueEnum struct {
	E_0 BatchUpdateDevicesStatusRequestValue
	E_1 BatchUpdateDevicesStatusRequestValue
}

func GetBatchUpdateDevicesStatusRequestValueEnum() BatchUpdateDevicesStatusRequestValueEnum {
	return BatchUpdateDevicesStatusRequestValueEnum{
		E_0: BatchUpdateDevicesStatusRequestValue{
			value: 0,
		}, E_1: BatchUpdateDevicesStatusRequestValue{
			value: 1,
		},
	}
}

func (c BatchUpdateDevicesStatusRequestValue) Value() int32 {
	return c.value
}

func (c BatchUpdateDevicesStatusRequestValue) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchUpdateDevicesStatusRequestValue) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
