package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowRedisDisabledCommandsResponse Response Object
type ShowRedisDisabledCommandsResponse struct {

	// 总数。
	TotalCount *string `json:"total_count,omitempty"`

	// 禁用类型。
	DisabledType *ShowRedisDisabledCommandsResponseDisabledType `json:"disabled_type,omitempty"`

	// disabled_type为command时展示该参数。
	Commands *[]string `json:"commands,omitempty"`

	// disabled_type为key时展示该参数，最多20个。
	Keys           *[]RedisDisabledCommandsDetail `json:"keys,omitempty"`
	HttpStatusCode int                            `json:"-"`
}

func (o ShowRedisDisabledCommandsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRedisDisabledCommandsResponse struct{}"
	}

	return strings.Join([]string{"ShowRedisDisabledCommandsResponse", string(data)}, " ")
}

type ShowRedisDisabledCommandsResponseDisabledType struct {
	value string
}

type ShowRedisDisabledCommandsResponseDisabledTypeEnum struct {
	COMMAND ShowRedisDisabledCommandsResponseDisabledType
	KEY     ShowRedisDisabledCommandsResponseDisabledType
}

func GetShowRedisDisabledCommandsResponseDisabledTypeEnum() ShowRedisDisabledCommandsResponseDisabledTypeEnum {
	return ShowRedisDisabledCommandsResponseDisabledTypeEnum{
		COMMAND: ShowRedisDisabledCommandsResponseDisabledType{
			value: "command",
		},
		KEY: ShowRedisDisabledCommandsResponseDisabledType{
			value: "key",
		},
	}
}

func (c ShowRedisDisabledCommandsResponseDisabledType) Value() string {
	return c.value
}

func (c ShowRedisDisabledCommandsResponseDisabledType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowRedisDisabledCommandsResponseDisabledType) UnmarshalJSON(b []byte) error {
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
