package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type SetInstanceDataDumpRequestBody struct {

	// OBS桶名。
	BucketName *string `json:"bucket_name,omitempty"`

	// 开启/关闭实例数据导出。
	Action SetInstanceDataDumpRequestBodyAction `json:"action"`
}

func (o SetInstanceDataDumpRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetInstanceDataDumpRequestBody struct{}"
	}

	return strings.Join([]string{"SetInstanceDataDumpRequestBody", string(data)}, " ")
}

type SetInstanceDataDumpRequestBodyAction struct {
	value string
}

type SetInstanceDataDumpRequestBodyActionEnum struct {
	OPEN  SetInstanceDataDumpRequestBodyAction
	CLOSE SetInstanceDataDumpRequestBodyAction
}

func GetSetInstanceDataDumpRequestBodyActionEnum() SetInstanceDataDumpRequestBodyActionEnum {
	return SetInstanceDataDumpRequestBodyActionEnum{
		OPEN: SetInstanceDataDumpRequestBodyAction{
			value: "open",
		},
		CLOSE: SetInstanceDataDumpRequestBodyAction{
			value: "close",
		},
	}
}

func (c SetInstanceDataDumpRequestBodyAction) Value() string {
	return c.value
}

func (c SetInstanceDataDumpRequestBodyAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SetInstanceDataDumpRequestBodyAction) UnmarshalJSON(b []byte) error {
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
