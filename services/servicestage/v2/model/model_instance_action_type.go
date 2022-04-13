package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// 操作，支持start, stop, restart, scale, rollback。
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

func (c InstanceActionType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *InstanceActionType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
