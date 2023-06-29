package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ServerHaltType 停机类型 * `SOFT` - 普通 * `HARD` - 强制
type ServerHaltType struct {
	value string
}

type ServerHaltTypeEnum struct {
	SOFT ServerHaltType
	HARD ServerHaltType
}

func GetServerHaltTypeEnum() ServerHaltTypeEnum {
	return ServerHaltTypeEnum{
		SOFT: ServerHaltType{
			value: "SOFT",
		},
		HARD: ServerHaltType{
			value: "HARD",
		},
	}
}

func (c ServerHaltType) Value() string {
	return c.value
}

func (c ServerHaltType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ServerHaltType) UnmarshalJSON(b []byte) error {
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
