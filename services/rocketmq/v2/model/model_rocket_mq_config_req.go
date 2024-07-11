package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type RocketMqConfigReq struct {

	// RocketMQ配置名称。
	Name *RocketMqConfigReqName `json:"name,omitempty"`

	// RocketMQ配置目标值。
	Value *string `json:"value,omitempty"`
}

func (o RocketMqConfigReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RocketMqConfigReq struct{}"
	}

	return strings.Join([]string{"RocketMqConfigReq", string(data)}, " ")
}

type RocketMqConfigReqName struct {
	value string
}

type RocketMqConfigReqNameEnum struct {
	FILE_RESERVED_TIME RocketMqConfigReqName
}

func GetRocketMqConfigReqNameEnum() RocketMqConfigReqNameEnum {
	return RocketMqConfigReqNameEnum{
		FILE_RESERVED_TIME: RocketMqConfigReqName{
			value: "fileReservedTime",
		},
	}
}

func (c RocketMqConfigReqName) Value() string {
	return c.value
}

func (c RocketMqConfigReqName) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RocketMqConfigReqName) UnmarshalJSON(b []byte) error {
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
