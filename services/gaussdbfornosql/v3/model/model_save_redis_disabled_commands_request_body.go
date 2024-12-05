package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type SaveRedisDisabledCommandsRequestBody struct {

	// 禁用类型。
	DisabledType SaveRedisDisabledCommandsRequestBodyDisabledType `json:"disabled_type"`

	// disabled_type为command时传入该参数。
	Commands *[]string `json:"commands,omitempty"`

	// disabled_type为key时传入该参数，最多20个。
	Keys *[]RedisDisabledCommandsDetail `json:"keys,omitempty"`
}

func (o SaveRedisDisabledCommandsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SaveRedisDisabledCommandsRequestBody struct{}"
	}

	return strings.Join([]string{"SaveRedisDisabledCommandsRequestBody", string(data)}, " ")
}

type SaveRedisDisabledCommandsRequestBodyDisabledType struct {
	value string
}

type SaveRedisDisabledCommandsRequestBodyDisabledTypeEnum struct {
	COMMAND SaveRedisDisabledCommandsRequestBodyDisabledType
	KEY     SaveRedisDisabledCommandsRequestBodyDisabledType
}

func GetSaveRedisDisabledCommandsRequestBodyDisabledTypeEnum() SaveRedisDisabledCommandsRequestBodyDisabledTypeEnum {
	return SaveRedisDisabledCommandsRequestBodyDisabledTypeEnum{
		COMMAND: SaveRedisDisabledCommandsRequestBodyDisabledType{
			value: "command",
		},
		KEY: SaveRedisDisabledCommandsRequestBodyDisabledType{
			value: "key",
		},
	}
}

func (c SaveRedisDisabledCommandsRequestBodyDisabledType) Value() string {
	return c.value
}

func (c SaveRedisDisabledCommandsRequestBodyDisabledType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SaveRedisDisabledCommandsRequestBodyDisabledType) UnmarshalJSON(b []byte) error {
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
