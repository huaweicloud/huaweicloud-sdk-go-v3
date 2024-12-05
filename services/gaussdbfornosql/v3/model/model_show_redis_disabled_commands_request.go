package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowRedisDisabledCommandsRequest Request Object
type ShowRedisDisabledCommandsRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	// 指令类型
	Type ShowRedisDisabledCommandsRequestType `json:"type"`

	// 索引位置偏移量，表示从指定offset条数据后查询对应的实例信息。 取值大于或等于0。不传该参数时，查询偏移量默认为0。
	Offset *int32 `json:"offset,omitempty"`

	// 查询专属资源个数上限值。   - 取值范围：1~50。不传该参数时，默认查询前50条实例信息。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ShowRedisDisabledCommandsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRedisDisabledCommandsRequest struct{}"
	}

	return strings.Join([]string{"ShowRedisDisabledCommandsRequest", string(data)}, " ")
}

type ShowRedisDisabledCommandsRequestType struct {
	value string
}

type ShowRedisDisabledCommandsRequestTypeEnum struct {
	COMMAND ShowRedisDisabledCommandsRequestType
	KEY     ShowRedisDisabledCommandsRequestType
}

func GetShowRedisDisabledCommandsRequestTypeEnum() ShowRedisDisabledCommandsRequestTypeEnum {
	return ShowRedisDisabledCommandsRequestTypeEnum{
		COMMAND: ShowRedisDisabledCommandsRequestType{
			value: "command",
		},
		KEY: ShowRedisDisabledCommandsRequestType{
			value: "key",
		},
	}
}

func (c ShowRedisDisabledCommandsRequestType) Value() string {
	return c.value
}

func (c ShowRedisDisabledCommandsRequestType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowRedisDisabledCommandsRequestType) UnmarshalJSON(b []byte) error {
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
