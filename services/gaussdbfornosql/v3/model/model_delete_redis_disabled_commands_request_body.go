package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type DeleteRedisDisabledCommandsRequestBody struct {

	// 禁用类型。
	DisabledType DeleteRedisDisabledCommandsRequestBodyDisabledType `json:"disabled_type"`

	// disabled_type为command时传入该参数。
	Commands *[]string `json:"commands,omitempty"`

	// disabled_type为key时传入该参数，最多20个。
	Keys *[]RedisDisabledCommandsDetail `json:"keys,omitempty"`
}

func (o DeleteRedisDisabledCommandsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteRedisDisabledCommandsRequestBody struct{}"
	}

	return strings.Join([]string{"DeleteRedisDisabledCommandsRequestBody", string(data)}, " ")
}

type DeleteRedisDisabledCommandsRequestBodyDisabledType struct {
	value string
}

type DeleteRedisDisabledCommandsRequestBodyDisabledTypeEnum struct {
	COMMAND DeleteRedisDisabledCommandsRequestBodyDisabledType
	KEY     DeleteRedisDisabledCommandsRequestBodyDisabledType
}

func GetDeleteRedisDisabledCommandsRequestBodyDisabledTypeEnum() DeleteRedisDisabledCommandsRequestBodyDisabledTypeEnum {
	return DeleteRedisDisabledCommandsRequestBodyDisabledTypeEnum{
		COMMAND: DeleteRedisDisabledCommandsRequestBodyDisabledType{
			value: "command",
		},
		KEY: DeleteRedisDisabledCommandsRequestBodyDisabledType{
			value: "key",
		},
	}
}

func (c DeleteRedisDisabledCommandsRequestBodyDisabledType) Value() string {
	return c.value
}

func (c DeleteRedisDisabledCommandsRequestBodyDisabledType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DeleteRedisDisabledCommandsRequestBodyDisabledType) UnmarshalJSON(b []byte) error {
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
