package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// InstanceActionType 操作，支持start, stop, restart, scale, rollback。
type InstanceActionType struct {
	value string
}

type InstanceActionTypeEnum struct {
	START    InstanceActionType
	STOP     InstanceActionType
	RESTART  InstanceActionType
	SCALE    InstanceActionType
	ROLLBACK InstanceActionType
}

func GetInstanceActionTypeEnum() InstanceActionTypeEnum {
	return InstanceActionTypeEnum{
		START: InstanceActionType{
			value: "start",
		},
		STOP: InstanceActionType{
			value: "stop",
		},
		RESTART: InstanceActionType{
			value: "restart",
		},
		SCALE: InstanceActionType{
			value: "scale",
		},
		ROLLBACK: InstanceActionType{
			value: "rollback",
		},
	}
}

func (c InstanceActionType) Value() string {
	return c.value
}

func (c InstanceActionType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *InstanceActionType) UnmarshalJSON(b []byte) error {
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
