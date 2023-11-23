package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type DeleteLtsConfigsRequestBody struct {

	// 需要解除关联的实例ID列表。
	InstanceIds []string `json:"instance_ids"`

	// 日志类型。slow_log表示慢日志。
	LogType DeleteLtsConfigsRequestBodyLogType `json:"log_type"`
}

func (o DeleteLtsConfigsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteLtsConfigsRequestBody struct{}"
	}

	return strings.Join([]string{"DeleteLtsConfigsRequestBody", string(data)}, " ")
}

type DeleteLtsConfigsRequestBodyLogType struct {
	value string
}

type DeleteLtsConfigsRequestBodyLogTypeEnum struct {
	SLOW_LOG DeleteLtsConfigsRequestBodyLogType
}

func GetDeleteLtsConfigsRequestBodyLogTypeEnum() DeleteLtsConfigsRequestBodyLogTypeEnum {
	return DeleteLtsConfigsRequestBodyLogTypeEnum{
		SLOW_LOG: DeleteLtsConfigsRequestBodyLogType{
			value: "slow_log",
		},
	}
}

func (c DeleteLtsConfigsRequestBodyLogType) Value() string {
	return c.value
}

func (c DeleteLtsConfigsRequestBodyLogType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DeleteLtsConfigsRequestBodyLogType) UnmarshalJSON(b []byte) error {
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
